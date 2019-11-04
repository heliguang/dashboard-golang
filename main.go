package main

import (
	"encoding/json"
	"fmt"

	"dashboard/config"
	"dashboard/gin"
	"dashboard/logger"
	"dashboard/storage"
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
		Status:   0,
	}
	fmt.Println(storage.UserInsert(adminUser))

	user1 := &storage.User{
		Account:  "user1",
		Password: "123456",
		Role:     "admin",
		Status:   1,
	}
	fmt.Println(storage.UserInsert(user1))

	user2 := &storage.User{
		Account:  "user2",
		Password: "123456",
		Role:     "admin",
		Status:   0,
	}
	fmt.Println(storage.UserInsert(user2))
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
