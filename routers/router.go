package routers

import (
	"bee_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:id",&controllers.TestController{},"get:Get")
    beego.Router("/register",&controllers.TestController{},"get:Register;post:Commit")
    //beego.Router("/register",&controllers.TestController{},"post:register")
    beego.Router("/login",&controllers.TestController{},"get:Login")
    beego.Router("/insert",&controllers.TestController{},"get:Insert")
    beego.Router("/update",&controllers.TestController{},"get:Update")
    beego.Router("/delete",&controllers.TestController{},"get:Delete")
    beego.Router("/show",&controllers.TestController{},"get:Show")
    beego.Router("/json",&controllers.TestController{},"get:Json;post:Jsons")
    //beego.Router("/test",&controllers.TestController{},"get:Test")
}
