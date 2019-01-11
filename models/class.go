package models

import (
	_ "github.com/go-sql-driver/mysql"
)


type Class struct {
	Id int
	ClassName  string
	Students []*Student `orm:"reverse(many)"`
}
