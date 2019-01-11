package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id int
	Name string
	age int
	Class *Class `orm:"rel(fk)"`
}

