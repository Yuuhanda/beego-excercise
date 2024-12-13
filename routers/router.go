package routers

import (
    "myproject/controllers"
    "myproject/database"
    beego "github.com/beego/beego/v2/server/web"
)

func init() {
    database.GetInstance()
    InitRoutes()
}

func InitRoutes() {
    // Get database instance first
    database.GetInstance()

    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/user", &controllers.UserController{}, "post:CreateUser")
    beego.Router("/api/user/:id", &controllers.UserController{}, "get:GetUser;put:UpdateUser;delete:DeleteUser")
    beego.Router("/api/users", &controllers.UserController{}, "get:ListUsers")


    // ItemUnit routes using the initialized controller instance
    beego.Router("/api/item-units",  &controllers.ItemUnitController{}, "post:Create")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "get:Get")
    beego.Router("/api/item-units",  &controllers.ItemUnitController{}, "get:List")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "put:Update")
    beego.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "delete:Delete")
    beego.Router("/api/item-units/serial/:serialNumber",  &controllers.ItemUnitController{}, "get:GetBySerialNumber")
    beego.Router("/api/item-units/warehouse/:warehouseId", &controllers.ItemUnitController{}, "get:GetByWarehouse")

    // Item routes using the initialized controller instance
    beego.Router("/api/item", &controllers.ItemController{}, "post:CreateItem")
    beego.Router("/api/item/:id", &controllers.ItemController{}, "get:GetItem;put:UpdateItem;delete:DeleteItem")
    beego.Router("/api/items", &controllers.ItemController{}, "get:ListItems")
    beego.Router("/api/items/dashboard", &controllers.ItemController{}, "get:SearchDashboard")


    // Item Category Routes
    categoryCtrl := controllers.NewItemCategoryController()
    beego.Router("/api/categories", categoryCtrl, "get:List")
    beego.Router("/api/categories/:id", categoryCtrl, "get:Get")
    beego.Router("/api/categories", categoryCtrl, "post:Create")
    beego.Router("/api/categories/:id", categoryCtrl, "put:Update")
    beego.Router("/api/categories/:id", categoryCtrl, "delete:Delete")

    employeeCtrl := controllers.NewEmployeeController()
    // Employee Routes
    beego.Router("/api/employees", employeeCtrl, "post:Create")
    beego.Router("/api/employees/:id", employeeCtrl, "get:Get")
    beego.Router("/api/employees", employeeCtrl, "get:List")
    beego.Router("/api/employees/:id", employeeCtrl, "put:Update")
    beego.Router("/api/employees/:id", employeeCtrl, "delete:Delete")

    // Lending Routes
    lendingCtrl := &controllers.LendingController{}
    beego.Router("/api/lendings", lendingCtrl, "post:Create")
    beego.Router("/api/lendings/:id", lendingCtrl, "get:Get")
    beego.Router("/api/lendings", lendingCtrl, "get:List")
    beego.Router("/api/lendings/:id", lendingCtrl, "put:Update")
    beego.Router("/api/lendings/:id", lendingCtrl, "delete:Delete")
    beego.Router("/api/lendings/report/items", &controllers.LendingController{}, "get:SearchItemReport")
    beego.Router("/api/lendings/report/units", &controllers.LendingController{}, "get:SearchUnitReport")

    
    // Additional lending-specific routes
    beego.Router("/api/lendings/active", lendingCtrl, "get:GetActiveLoans")
    beego.Router("/api/lendings/returned", lendingCtrl, "get:GetReturnedLoans")
    
    // Unit Log Routes
    beego.Router("/api/unit-logs", &controllers.UnitLogController{}, "post:Create;get:List")
    beego.Router("/api/unit-logs/:id", &controllers.UnitLogController{}, "get:Get;put:Update;delete:Delete")
    beego.Router("/api/unit-logs/unit/:unitId", &controllers.UnitLogController{}, "get:GetByUnit")

    // Repair Log Routes
    beego.Router("/api/repair-logs", &controllers.RepairLogController{}, "post:Create;get:List")
    beego.Router("/api/repair-logs/:id", &controllers.RepairLogController{}, "get:Get;put:Update;delete:Delete")
    beego.Router("/api/repair-logs/unit/:unitId", &controllers.RepairLogController{}, "get:GetByUnit")

    // Document Upload Routes
    beego.Router("/api/docs", &controllers.DocUploadedController{}, "post:Create;get:List")
    beego.Router("/api/docs/:id", &controllers.DocUploadedController{}, "get:Get;delete:Delete")

    beego.Router("/api/warehouse", &controllers.WarehouseController{}, "post:CreateWarehouse")
    beego.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "get:GetWarehouse")
    beego.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "put:UpdateWarehouse")
    beego.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "delete:DeleteWarehouse")
    beego.Router("/api/warehouses", &controllers.WarehouseController{}, "get:ListWarehouses")
    beego.Router("/api/warehouse/:id/users", &controllers.WarehouseController{}, "get:GetWarehouseUsers")
}