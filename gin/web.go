package gin

import (
	"dashboard/config"
	"dashboard/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func webRouter() http.Handler {
	router := gin.New()
	router.Use(ginLogger())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}

var webServer = &http.Server{
	Addr:         config.Conf.WebAddress,
	Handler:      webRouter(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}

func RunWebServer() error {
	logger.Info("run web server on:[" + webServer.Addr + "]")
	return webServer.ListenAndServe()
}
