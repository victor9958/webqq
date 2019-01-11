package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
)

type MainController struct {
	beego.Controller
}

var (
	NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
)

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Ctx.WriteString("hello")

	for _,v:=range NotPV{
		c.Ctx.WriteString(strings.ToUpper(v)+"\r\n")
	}
}

func (c *MainController) Post() {
	str:=strconv.Itoa(555)
	c.Ctx.WriteString(str)
}

