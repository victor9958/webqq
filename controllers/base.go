package controllers

import (
	"github.com/astaxie/beego"
)
type BaseController struct {
	beego.Controller
	isLogin bool
}

func (c *BaseController) Prepare()  {

	userLogin := c.GetSession("userLogin")
	if userLogin != nil{
		c.isLogin = false
		beego.Info("false")
	}else{
		c.isLogin = true
		beego.Info("true")
	}

	c.Data["isLogin"] = true

}



