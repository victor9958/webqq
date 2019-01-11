package controllers

import (
	"webqq/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}


func (c *IndexController) Get() {
	//str:=strconv.Itoa(66666)
	//c.Ctx.WriteString(str)
	//fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id

	///c.Ctx.WriteString(fmt.Sprintf("这是字符串中加[%s]","sfdsfsd"))

	c.TplName = "index.html"
}


func (c *IndexController) Post() {
	var class models.Class
	err :=orm.NewOrm().QueryTable("class").RelatedSel().One(&class)
	if err != nil {
		return
	}
	/*******************成功**********************/
	//class := models.Class{Id:1}
	//err :=orm.NewOrm().Read(&class)
	//if err != nil {
	//	return
	//}
	//_,err2 :=orm.NewOrm().LoadRelated(&class,"Students")
	//
	//if err2 != nil {
	//	return
	//}

	/************成功********************/





	beego.Info(class)
	c.Data["json"] = class
	c.ServeJSON()

}


func (c *IndexController) Put() {
	var students []*models.Student
	_,err :=orm.NewOrm().QueryTable(new(models.Student)).Filter("Class__Id",1).All(&students)
	if err != nil {
		return
	}
	c.Data["json"] = students
	c.ServeJSON()

}