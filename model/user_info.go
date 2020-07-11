package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//user info
type UserInfo struct {
	Id       string
	Pwd      string
	Token    string
	Name     string
	Email    string
	Phone    string
	Nirthday string
}

// Regist the user
func (user *UserInfo) Regist() {

}
