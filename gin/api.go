package gin

import (
	"dashboard/config"
	"dashboard/logger"
	"dashboard/storage"

	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type User struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Role     string `form:"role" json:"role" xml:"role" binding:"required"`
	Status   string `form:"status" json:"status" xml:"status" binding:"required"`
}

func apiRouter() http.Handler {
	router := gin.New()
	router.Use(cors(), ginLogger())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/user/login", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			logger.Error("login bind json err[" + err.Error() + "]")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logger.Info("login info:" + login.User + " " + login.Password)

		//if login.User != "manu" || login.Password != "123" {
		//	logger.Error("login bind user or password error")
		//	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		//	return
		//}

		c.JSON(http.StatusOK, gin.H{"code": 20000, "token": "admin-token"})
	})
	router.POST("/user/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
	})
	router.GET("/user/info", func(c *gin.Context) {
		token := c.Query("token")
		logger.Info("get user info, user token:" + token)

		c.JSON(http.StatusOK,
			gin.H{"code": 20000,
				"roles":        []string{"admin"},
				"introduction": "I am a super administrator",
				"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
				"name":         "Super Admin"},
		)
	})
	router.GET("/user/asyncRoutes", func(c *gin.Context) {
		token := c.Query("token")
		logger.Info("get user info, user token:" + token)

		c.JSON(http.StatusOK,
			gin.H{"code": 20000,
				"routes": config.Conf.Routes,
			},
		)
	})
	router.GET("/table/list", func(c *gin.Context) {
		item := []struct {
			Id          int    `json:"id"`
			Title       string `json:"title"`
			Status      string `json:"status"`
			Author      string `json:"author"`
			DisplayTime string `json:"display_time"`
			Pageviews   int    `json:"pageviews"`
		}{
			{
				1,
				"title",
				"published",
				"author",
				"display_time",
				1,
			},
			{
				1,
				"title",
				"draft",
				"author",
				"display_time",
				1,
			},
			{
				1,
				"title",
				"deleted",
				"author",
				"display_time",
				1,
			},
		}

		data := struct {
			Total int `json:"total"`
			Items []struct {
				Id          int    `json:"id"`
				Title       string `json:"title"`
				Status      string `json:"status"`
				Author      string `json:"author"`
				DisplayTime string `json:"display_time"`
				Pageviews   int    `json:"pageviews"`
			} `json:"items"`
		}{
			len(item),
			item,
		}

		c.JSON(http.StatusOK,
			gin.H{"code": 20000, "data": data},
		)
	})
	router.GET("/admin/getAllUser", func(c *gin.Context) {
		users, err := storage.UserGetAll()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 40000})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 20000, "users": users})
		}
	})
	router.POST("/admin/addUser", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			logger.Error("login bind json err[" + err.Error() + "]")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logger.Info("login info:" + user.User + " " + user.Password + " " + user.Role)
		_, err := storage.UserInsert(&storage.User{
			Account:  user.User,
			Password: user.Password,
			Role:     user.Role,
		})

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 40000})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 20000})
		}
	})
	router.POST("/admin/deleteUser", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			logger.Error("login bind json err[" + err.Error() + "]")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logger.Info("login info:" + user.User + " " + user.Password + " " + user.Role)

		_, err := storage.UserDelete(user.User)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 40000})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 20000})
		}
	})
	router.POST("/admin/updateUser", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			logger.Error("login bind json err[" + err.Error() + "]")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logger.Info("login info:" + user.User + " " + user.Password + " " + user.Role)
		_, err := storage.UserUpdate(user.User, &storage.User{
			Account:  user.User,
			Password: user.Password,
			Role:     user.Role,
		})

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 40000})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 20000})
		}
	})
	return router
}

var apiServer = &http.Server{
	Addr:         config.Conf.ApiAddress,
	Handler:      apiRouter(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}

func RunApiServer() error {
	logger.Info("run api server on:[" + apiServer.Addr + "]")
	return apiServer.ListenAndServe()
}
