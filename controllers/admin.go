package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) Get() {
	str:=strconv.Itoa(655)
	c.Ctx.WriteString(str)
}