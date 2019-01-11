package controllers

import (
	"github.com/astaxie/beego/orm"
	"webqq/models"
	"github.com/astaxie/beego"
	"crypto/md5"
	"fmt"
)

type WeblistController struct {
	BaseController

	//github.com/astaxie/beego/orm
}

func(web *WeblistController)Get(){
	//user  := make()
	var userData models.User
	//参数
	userRes := map[string]interface{}{}
	//获取访问的ip
	req := web.Ctx.Request
	add := req.RemoteAddr
	userRes["username"] = add

	//拿出我的用户信息

	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")

	err := qs.Filter("ip","127.0.0.2").One(&userData)
	//qs.Filter("ip").From("user").Where("ip="+add).
	//userData := models.User{Ip:}
	//if err != nil{
	//	beego.Info(err2)
	//	return
	//}
	//err := o.Read(&userData)
	if err != nil{
		beego.Error(err)

		user := models.User{}
		user.Ip = add

		//生成一个md5的字符串
		pwd := []byte("123456") //二进制
		has := md5.Sum(pwd) //二进制
		has16 := fmt.Sprintf("%x", has) //转化成16进制
		//
		user.Pwd = has16
		user.Sex = 1
		user.HeadImg = "http://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83epiahkUQrtA6XRdd20ZMavWzLM9WrMtuxHxsLxia0O7QgnG75lfADdialvCdr5KxMNxia1kFVWNibK4pvg/132";
		user.Sign = "这是weblist控制器添加获得的简介,ip:"+user.Ip
		user.LineStatus = 1
		id ,err1 := o.Insert(&user)
		if err1 != nil {
			beego.Info("weblist新增错误")
			return
		}


		userRes["id"] = id
		userRes["status"] = "online"
		userRes["sign"] = user.Sign
		userRes["avatar"] = user.HeadImg
		beego.Info("这里是添加")

	}else{
		//本身就存在的用户
		userRes["id"] = userData.Id
		switch userData.Status {
		case 1:userRes["status"] = "online"
			break
		case 2:userRes["status"] = "online"
			break

		}
		userRes["sign"] = userData.Sign
		userRes["avatar"] = userData.HeadImg
		beego.Info("这里是获取")

	}



	//模拟数据
	web.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg": "",
		"data":map[string]interface{}{
			"mine":userRes,
			"friend":[...]interface{}{
				map[string]interface{}{
					"groupname":"后端分组",
					"id":1,//分组id
					"list":[...]interface{}{
						//这里不需要加0
							map[string]interface{}{
							"id":"100001",//好友id
							"avatar":"a.jpg",//头像
							"sign":"这里测试好友签名",//签名
							"status":"online",//状态
						},
					},
				},
			},

			"group":[...]interface{}{
				0:map[string]interface{}{
					"groupname"	:"前端群",//群名称
					"id"	:"101",//id
					"avatar"	:"a.jpg",//头像
				},
			},
	}}
	web.ServeJSON()
}

func(web *WeblistController)Post(){
	beego.Info(11111)
	o := orm.NewOrm()

	//var users []*models.User
	//_,err := o.QueryTable("user").Filter("Ip","127.0.0.1").All(&users)
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}

	////获取分组信息
	//var groups []*models.Group
	//_,err := o.QueryTable("group").
	//	Filter("id",1).
	//	All(&groups)
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}


	//获取分组信息
	//var users []*models.User
	//_,err := o.QueryTable("user").
	//	Filter("id",2).
	//	All(&users)
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}

	//获取我的分组内的用户
	//获取分组信息
	var groups []*models.Group
	_,err := o.QueryTable("group").
		All(&groups)
	if err != nil {
		beego.Error(err)
		return
	}


	web.Data["json"] = groups
	web.ServeJSON()


}


func(web *WeblistController)Put(){
	//获取分组信息
	//var users []*models.User
	//var users models.User
	//err := orm.NewOrm().QueryTable(new(models.User)).
	//	RelatedSel().
	//	Filter("Id",1).
	//	One(&users)


	//_,err := orm.NewOrm().QueryTable(new(models.User)).Filter("Groups__Id",1).RelatedSel().All(&users);


	var groupMaps []orm.Params
	_,err := orm.NewOrm().QueryTable(new(models.Group)).Filter("User",1).
		Values(&groupMaps)
	
	


	if err != nil {
		beego.Error(err)
		return
	}

	for _,m :=range groupMaps{
		var friendsData []*models.Friends
		_,err2 := orm.NewOrm().QueryTable(new(models.Friends)).
			Filter("Id",1).All(&friendsData)
		beego.Info(m["User"],"--------------")
		if err2 == nil {
			//var data [...]interface{}
			//for _,friendM:=range friendsData{
			//	data[0] = map[string]interface{}{
			//		"id":friendM.User.Id,
			//		"avatar":friendM.User.HeadImg,
			//		"sign":friendM.RemarkName,
			//		"status":friendM.User.LineStatus,
			//
			//		//"id":"100001",//好友id
			//		//							"avatar":"a.jpg",//头像
			//		//							"sign":"这里测试好友签名",//签名
			//		//							"status":"online",//状态
			//	}
			//	//data[]{map[]} = friendM.User.Ip
			//	web.Data["json"] = data
			//	web.ServeJSON()
			//}
			//

			web.Data["json"] = friendsData
			web.ServeJSON()
		}
		beego.Error(err2)

	}
	//var f interface{}
	//beego.Info(groupMaps)
	//m := f.(map[string]interface{})
	//for k,v:=range users{
	//	//if _,ok:=v.(string);ok  {
	//	//	beego.Info(k,"==string==>",v)
	//	//}
	//	beego.Info(k,"==string==>",v)
	//}

	//m := f.(map[string]interface{})
	//for k,v:=range m{
	//	//beego.Info(reflect.TypeOf(v))
	//	if _,ok:=v.(string);ok {
	//		beego.Info(k,"==stirng==>",v)
	//	}
	//	if _,ok2:=v.(float64);ok2 {
	//		beego.Info(k,"==float64==>",v)
	//	}
	//	beego.Info(k,"==>",v)
	//}

	//获取我的分组内的用户
	//获取分组信息
	//var groups []*models.Group
	//_,err := o.QueryTable("group").
	//	All(&groups)
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}


	web.Data["json"] = groupMaps
	web.ServeJSON()


}
func(web *WeblistController)Delete(){
	var f models.Friends

	err:=orm.NewOrm().QueryTable(new(models.Friends)).Filter("Id",1).One(&f)

	if err != nil {
		return
	}


	web.Data["json"] = f
	web.ServeJSON()

}