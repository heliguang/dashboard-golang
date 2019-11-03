package storage

import (
	"errors"
)

type User struct {
	Id          int64  `xorm:"pk autoincr" json:"-"`  //指定主键并自增
	Account     string `xorm:"unique" json:"account"` //唯一的
	Password    string `json:"-"`
	Role        string `json:"role"`
	UpdatedTime int64  `xorm:"updated" json:"-"` //修改后自动更新时间
	CreateTime  int64  `xorm:"created" json:"-"` //创建时间
}

//增
func UserInsert(user *User) (int64, error) {
	return engine.Insert(user)
}

//删
func UserDelete(account string) (int64, error) {
	return engine.Delete(&User{Account: account})
}

func UserUpdate(account string, user *User) (int64, error) {
	affected, err := engine.Update(user, &User{Account: account})
	if err != nil {
		return affected, err
	}
	return 0, nil
}

func UserUpdateRole(account string, role string) (int64, error) {
	affected, err := engine.Update(&User{Role: role}, &User{Account: account})
	if err != nil {
		return affected, err
	}
	return affected, nil
}

//查
func UserGet(account string) (*User, error) {
	user := &User{Account: account}
	is, err := engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !is {
		return nil, errors.New("account doesn't exists")
	}
	return user, nil
}

func UserExist(account string) (bool, error) {
	user := &User{Account: account}
	is, err := engine.Get(user)
	if err != nil {
		return false, err
	}
	return is, nil
}
