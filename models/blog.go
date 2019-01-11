package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id int
	Title string `orm:"size(255)"`
	Summary string `orm:"size(400)"`
	ReleaseDate int
	ClickNum int
	ReplyNum int
	Content string
	KeyWord string
	TypeId int
	Created int
	Updated int

}

