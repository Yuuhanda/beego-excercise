package routers

import (
    "myproject/controllers"
    "myproject/database"
    beego "github.com/beego/beego/v2/server/web"
)

func init() {
    // Get database instance first
    database.GetInstance()

    // Then set up routes
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{}, "post:CreateUser")
    beego.Router("/user/:id", &controllers.UserController{}, "get:GetUser;put:UpdateUser;delete:DeleteUser")
    beego.Router("/users", &controllers.UserController{}, "get:ListUsers")

    itemUnitController := controllers.NewItemUnitController()
    beego.Router("/api/item-units", itemUnitController, "post:Create")
    beego.Router("/api/item-units/:id", itemUnitController, "get:Get")
    beego.Router("/api/item-units", itemUnitController, "get:List")
    beego.Router("/api/item-units/:id", itemUnitController, "put:Update")
    beego.Router("/api/item-units/:id", itemUnitController, "delete:Delete")
    beego.Router("/api/item-units/serial/:serialNumber", itemUnitController, "get:GetBySerialNumber")
    beego.Router("/api/item-units/warehouse/:warehouseId", itemUnitController, "get:GetByWarehouse")
}

func InitRoutes() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{}, "post:CreateUser")
    beego.Router("/user/:id", &controllers.UserController{}, "get:GetUser;put:UpdateUser;delete:DeleteUser")
    beego.Router("/users", &controllers.UserController{}, "get:ListUsers")


    // ItemUnit routes using the initialized controller instance
    beego.Router("/api/item-units",  &controllers.ItemUnitController{}, "post:Create")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "get:Get")
    beego.Router("/api/item-units",  &controllers.ItemUnitController{}, "get:List")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "put:Update")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "delete:Delete")
    beego.Router("/api/item-units/serial/:serialNumber",  &controllers.ItemUnitController{}, "get:GetBySerialNumber")
    beego.Router("/api/item-units/warehouse/:warehouseId", &controllers.ItemUnitController{}, "get:GetByWarehouse")

    
}