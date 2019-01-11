package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Blogger struct {
	Id int
	Name string `orm:"size(100)"`
	Pwd string
	Age int
	Telephone string
	Created int64
	Updated int64


}

