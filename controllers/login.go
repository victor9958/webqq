package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
	//beego.Controller
	//BaseController
}

func (c *LoginController) Get()  {
	//c.Ctx.WriteString("loginfunc")
	name :=c.GetString("name")
	beego.Info(name)
	if c.Data["isLogin"] != nil {
		beego.Info(c.Data["isLogin"])
	}
	c.TplName = "login/login.html"
}


func (c *LoginController) Post()  {
	name :=c.GetString("name")
	pwd :=c.GetString("pwd")
	beego.Info(name)
	beego.Info(pwd)
	//c.TplName = "login/login.html"
	c.Ctx.Output.JSON(map[string]interface{}{"name":name,"pwd":pwd,"age":111}, true, false)
}

type Address struct {
	City int
	CityName string
	parentId int
}
//type User struct {
//	Id int
//	Name string
//	AddData []*Address
//}


func (c *LoginController) Put()  {

	//hz := &Address{15,"杭州市",1}
	//zj := &Address{1,"杭州市",0}
	//my := &User{648,"victor",[]*Address{hz,zj}}
	//
	//js,_:=json.Marshal(my)
	//
	//fmt.Printf("JSON format: %s", js)
	//
	//var f interface{}
	//
	//err := json.Unmarshal(js,&f)
	//var temp interface{}
	//temp = 2
	//switch vvv:=temp.(type) {
	// 	case int:println("int:",vvv)
	//default:println("这个不是int")
	//
	//}
	var f interface{}
	var str1 = `{"Id":648,"Name":"victor","AddData":[{"City":15,"CityName":"杭州市"},{"City":1,"CityName":"杭州市"}]}`
	//m := f.(map[string]interface{})
	//err := json.Unmarshal(str1,&f)
	err := json.Unmarshal([]byte(str1),&f)
	if err!=nil {
		return
	}
	m := f.(map[string]interface{})
	for k,v:=range m{
		//beego.Info(reflect.TypeOf(v))
		if _,ok:=v.(string);ok {
			beego.Info(k,"==stirng==>",v)
		}
		if _,ok2:=v.(float64);ok2 {
			beego.Info(k,"==float64==>",v)
		}
		beego.Info(k,"==>",v)
	}
	//m := f.(map[string]interface{})
		//for k,v:=range m{
		//	println(v)
		//	switch vv:=v.(type) {
		//		case string:
		//			println(k,"is string",vv)
		//			break
		//		case int:
		//			println(k,"is int",vv)
		//			break
		//		case []interface{}:
		//			println(k,"is array")
		//			for i,u:=range vv{
		//				println(i,u)
		//			}
		//			break
		//	default:
		//		println("default",vv)
		//
		//	}
		//}


	//beego.Info(js)
	c.TplName = "login/login.html"

}

type mango_arealist struct {
	id int
	area_name string `orm:"size(100)"`
	level int8
	parent_id int
	flag int8
	created_at string `orm:"size(100)"`
	deleted_at string `orm:"size(100)"`
	updated_at string `orm:"size(100)"`
	agency_id int
}

func (c *LoginController)Delete(){
	o:=orm.NewOrm()
	var m []orm.Params
	_,err := o.Raw("SELECT * FROM mango_citylist limit 40000,42000").Values(&m)

	if err!=nil{
		return
	}

	//for k,v:=range m{
	//	beego.Info(k,"==>",v)
	//}
	c.Ctx.Output.JSON(&m, true, false)

}



