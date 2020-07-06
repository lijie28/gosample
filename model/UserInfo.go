package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//user info
type UserInfo struct {
	Id       int
	token    string
	Name     string
	Email    string
	Phone    string
	Birthday string
}

// Regist the user
func (user *UserInfo) Regist() {

}
