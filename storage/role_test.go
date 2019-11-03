package storage

import (
	"fmt"
	"testing"
)

func TestRoleInsert(t *testing.T) {
	role := &Role{
		Name:    "admin",
		Desc:    "超级管理员",
		Routers: "all",
	}

	fmt.Println(RoleInsert(role))
	fmt.Println(RoleInsert(role))
}

func TestRoleUpdate(t *testing.T) {
	newRole := &Role{
		Name:    "admin",
		Routers: "new-routers",
	}

	fmt.Println(RoleUpdate(newRole))
	fmt.Println(RoleUpdate(newRole))
}

func TestRoleGet(t *testing.T) {
	fmt.Println(RoleGet("admin"))
	fmt.Println(RoleGet("not-exit"))
}
