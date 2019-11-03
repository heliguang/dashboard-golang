package main

import (
	"dashboard/config"
	"dashboard/gin"
	"dashboard/logger"
	"dashboard/storage"
	"encoding/json"
	"fmt"
)

func createTestDatabase() {
	jsonBytes, err := json.Marshal(config.Conf.Routes)
	if err != nil {
		panic(err)
	}
	adminRole := &storage.Role{
		Name:   "admin",
		Desc:   "超级管理员权限",
		Routes: string(jsonBytes),
	}
	fmt.Println(storage.RoleInsert(adminRole))

	adminUser := &storage.User{
		Account:  "admin",
		Password: "123456",
		Role:     "admin",
	}
	fmt.Println(storage.UserInsert(adminUser))
}

func main() {
	logger.Info("main start")

	// 创建测试数据库
	//createTestDatabase()

	err := gin.RunApiServer()
	if err != nil {
		panic(err)
	}

	logger.Info("main finish")
}
