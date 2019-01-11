package main

import (
	_ "webqq/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"webqq/models"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/cache"
	"time"
	"crypto/md5"
	"fmt"
)

func insertCeshi(){
	o := orm.NewOrm()
	ceshiData := models.Blogger{}
	ceshiData.Name = "goCeshi"
	Id,err:=o.Insert(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",Id)
}

func insertUser(){
	o := orm.NewOrm()
	userData := models.User{}
	userData.Ip = "127.0.0.1"
	//生成一个md5的字符串
	pwd := []byte("123456") //二进制
	has := md5.Sum(pwd) //二进制
	has16 := fmt.Sprintf("%x", has) //转化成16进制
	//

	userData.Pwd = has16
	userData.Sex = 1
	userData.HeadImg = "http://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83epiahkUQrtA6XRdd20ZMavWzLM9WrMtuxHxsLxia0O7QgnG75lfADdialvCdr5KxMNxia1kFVWNibK4pvg/132";
	userData.Sign = "这是从数据库总获得的简介,ip:"+userData.Ip
	userData.LineStatus = 1

	Id,err:=o.Insert(&userData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",Id)
}

func insertBlog(){
	o := orm.NewOrm()
	ceshiData := models.Blog{}
	ceshiData.Title = "这是博客"
	ceshiData.Summary = "博客"
	ceshiData.Content = "这是博客的内容"
	Id,err:=o.Insert(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",Id)
}

func readCeshi(){
	o := orm.NewOrm()
	ceshiData := models.Blogger{Id:4825}
	err := o.Read(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",ceshiData)
}

func main() {
	//server := ""
	//添加用户
	//insertUser();

	//insertBlog()
	//insertCeshi()
	//readCeshi()
	cacheconn := beego.AppConfig.String("redis.conn")
	cachepwd := beego.AppConfig.String("redis.pwd")
	cachedb := beego.AppConfig.String("redis.cachedbname")
	cachestr := `{"key":"victor","conn":"`+cacheconn+`","dbNum":"`+cachedb+`","password":"`+ cachepwd+`"}`
	//beego.Info(cachestr)
	bm,err :=cache.NewCache("redis",cachestr)
	if err != nil{
		beego.Info("redis作为缓存错误是:",err)
	}
	///beego.Info(bm)
	bm.Put("name",1,10*time.Second)
	bm.Put("age",1000,1000*time.Second)
	//beego.Info(bm.Get("name"))
	beego.Run()
}

