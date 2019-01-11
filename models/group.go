package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Group struct {
	Id int
	//UserId int
	GroupName string `orm:"size(32)"`
	RemarkName string `orm:"size(50)"`
	User *User `orm:"rel(fk)"`
}

