package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Post ()  {

}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}
