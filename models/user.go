package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Ip string `orm:"size(20)"`
	Pwd string `orm:"size(32)"`
	Sex int
	HeadImg string `orm:"size(200)"`
	Sign string `orm:"size(200)"`
	LineStatus int
	Status int
	Groups []*Group `orm:"reverse(many)"`
	//Friends    *Friends   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
	Friends *Friends `orm:"reverse(one)"`
}

