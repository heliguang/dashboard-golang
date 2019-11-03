package storage

import (
	"errors"
)

type Role struct {
	Id          int64  `xorm:"pk autoincr" json:"-"` //指定主键并自增
	Name        string `xorm:"unique" json:"name"`   //唯一的
	Desc        string `json:"desc"`
	Routes      string `json:"routes"`
	UpdatedTime int64  `xorm:"updated" json:"-"` //修改后自动更新时间
	CreateTime  int64  `xorm:"created" json:"-"` //创建时间
}

//增
func RoleInsert(role *Role) (int64, error) {
	return engine.Insert(role)
}

//删
func RoleDelete(name string) (int64, error) {
	return engine.Delete(&Role{Name: name})
}

func RoleUpdate(role *Role) (int64, error) {
	affected, err := engine.Update(role, &Role{Name: role.Name})
	if err != nil {
		return affected, err
	}
	return 0, nil
}

//查
func RoleGet(name string) (*Role, error) {
	user := &Role{Name: name}
	is, err := engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !is {
		return nil, errors.New("account doesn't exists")
	}
	return user, nil
}

func RoleExist(name string) (bool, error) {
	user := &Role{Name: name}
	is, err := engine.Get(user)
	if err != nil {
		return false, err
	}
	return is, nil
}

func RoleGetAll() ([]Role, error) {
	roles := make([]Role, 0)
	err := engine.Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
