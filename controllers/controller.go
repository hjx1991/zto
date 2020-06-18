package controllers

import (
	"bee_project/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"html/template"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Json(){
	users:=models.Users{Id:1,Username:"hjx",Age:&models.Age{Id:1,Age:28}}
	this.Data["json"]=users
	this.ServeJSON()
}

func (this *TestController) Jsons(){
	u:=models.UserFrom{}
	body:=this.Ctx.Input.RequestBody
	if err:=json.Unmarshal(body,&u);err !=nil{
		fmt.Println("解析错误")
		return
	}
	fmt.Println("json data:",u)
	this.Data["jsondata"]=body
	this.TplName="jsons.tpl"

}

func (this *TestController) Show(){
	user:=[]models.User{}

	orm.NewOrm().QueryTable("user").Filter("username","22K").Filter("password","22K").All(&user,"id","username","password")
	fmt.Println(user)
	//err:=o.Read(&user)
	//if err !=nil{
	//	fmt.Println("查询错误",err)
	//}
	//beego.Info("查询成功:",user)
	//this.Data["user"]=user
	this.TplName="show.html"

}

func (this *TestController) Insert() {
	o := orm.NewOrm()
	user:=[]models.User{{Username:"t1",Password:"123456"},{Username:"t2",Password:"123456"}}
	num,err:=o.InsertMulti(2,&user)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(num)
	this.Data["inserts"]=user
	this.TplName="insert.tpl"
}

func (this *TestController) Update(){
	num,err:=orm.NewOrm().QueryTable("user").Filter("username","hjx").Update(orm.Params{
		"username":"test",
	})
	if err !=nil{
		fmt.Println("err:",err)
		return
	}
	this.Data["num"]=num
	this.TplName="update.tpl"
}

func (this *TestController) Delete() {
	num,err:=orm.NewOrm().QueryTable("user").Filter("username","test").Delete()
	if err !=nil{
		fmt.Println(err)
		return
	}
	this.Data["num"]=&num
	fmt.Println(num)
	this.TplName="delete.tpl"
}

func (this *TestController) Register(){
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	this.TplName="register.tpl"
}

func (this *TestController) Commit()  {
	Username:=this.GetString("username")
	Passwd:=this.GetString("password")
	//this.Data["username"]=Username
	//this.Data["password"]=Passwd
	beego.Info(Username,Passwd)
	user:=models.User{Username:Username,Password:Passwd}
	beego.Info(models.User{}.Password)
	num,err:=orm.NewOrm().Insert(&user)
	if err !=nil{
		fmt.Println(err)
		return
	}
	beego.Info(num)
	this.TplName="commit.tpl"
}

func (this *TestController) Login(){
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	//fmt.Println(this.Data["xsrfdata"])
	this.TplName="post.tpl"
	//this.Ctx.ResponseWriter.Write([]byte("login"))
}


