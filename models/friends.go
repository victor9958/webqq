package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Friends struct {
	Id int
	//FriendId int
	User *User `orm:"rel(one)"`
	//Friends    *Friends   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
	GroupId int
	//Group *Group `orm:"rel(one)"`
	RemarkName string `orm:"size(50)"`
}

