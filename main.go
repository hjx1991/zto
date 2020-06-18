package main

import (
	_ "bee_project/models"
	_ "bee_project/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
func init()  {
	beego.AppConfig.String("mysql_user")
	beego.AppConfig.String("mysql_password")
	beego.AppConfig.String("mysql_host")
	beego.AppConfig.String("mysql_dbname")
	beego.BConfig.Listen.AdminPort = 8088
}

func main() {
	beego.SetStaticPath("/down","static")
	beego.Run()
}

