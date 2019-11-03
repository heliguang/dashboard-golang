package storage

import (
	"fmt"
	"testing"
)

func TestUserInsert(t *testing.T) {
	user := &User{
		Account:  "admin",
		Password: "123456",
		Role:     "admin",
	}

	fmt.Println(UserInsert(user))
	fmt.Println(UserInsert(user))
}

func TestUserGet(t *testing.T) {
	fmt.Println(UserGet("admin"))
}

func TestUserExist(t *testing.T) {
	fmt.Println(UserExist("admin"))
}
