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

    beego.Router("/item", &controllers.ItemController{}, "post:CreateItem")
    beego.Router("/item/:id", &controllers.ItemController{}, "get:GetItem;put:UpdateItem;delete:DeleteItem")
    beego.Router("/items", &controllers.ItemController{}, "get:ListItems")

    categoryCtrl := controllers.NewItemCategoryController()
    
    // Item Category Routes
    beego.Router("/api/categories", categoryCtrl, "post:Create")
    beego.Router("/api/categories/:id", categoryCtrl, "get:Get")
    beego.Router("/api/categories", categoryCtrl, "get:List")
    beego.Router("/api/categories/:id", categoryCtrl, "put:Update")
    beego.Router("/api/categories/:id", categoryCtrl, "delete:Delete")
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

    // Item routes using the initialized controller instance
    beego.Router("/item", &controllers.ItemController{}, "post:CreateItem")
    beego.Router("/item/:id", &controllers.ItemController{}, "get:GetItem;put:UpdateItem;delete:DeleteItem")
    beego.Router("/items", &controllers.ItemController{}, "get:ListItems")

    // Item Category Routes
    categoryCtrl := controllers.NewItemCategoryController()
    beego.Router("/api/categories", categoryCtrl, "get:List")
    beego.Router("/api/categories/:id", categoryCtrl, "get:Get")
    beego.Router("/api/categories", categoryCtrl, "post:Create")
    beego.Router("/api/categories/:id", categoryCtrl, "put:Update")
    beego.Router("/api/categories/:id", categoryCtrl, "delete:Delete")
}