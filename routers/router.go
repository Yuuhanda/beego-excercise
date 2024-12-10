package routers

import (
	"myproject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "post:CreateUser")
    beego.Router("/user/:id", &controllers.UserController{}, "get:GetUser;put:UpdateUser;delete:DeleteUser")
    beego.Router("/users", &controllers.UserController{}, "get:ListUsers")
}

