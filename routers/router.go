package routers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/controllers"
    "myproject/middleware"
    "myproject/database"
)

func init() {
    database.GetInstance()
    InitRoutes()
}

func InitRoutes() {
    // Get database instance first
    database.GetInstance()

    // Public routes
    web.Router("/auth/login", &controllers.AuthController{}, "post:Login")

    // Admin-only route with multiple middleware
    web.InsertFilter("/user", web.BeforeRouter, middleware.AuthMiddleware())
    web.InsertFilter("/user", web.BeforeRouter, middleware.AdminMiddleware())
    web.Router("/user", &controllers.UserController{}, "post:CreateUser")

    // Other protected routes
    web.InsertFilter("/user/*", web.BeforeRouter, middleware.AuthMiddleware())
    web.InsertFilter("/users", web.BeforeRouter, middleware.AuthMiddleware())
    web.InsertFilter("/api/*", web.BeforeRouter, middleware.AuthMiddleware())
    
    // User Routes
    web.Router("/auth/logout", &controllers.AuthController{}, "post:Logout")
    web.Router("/user/:id", &controllers.UserController{}, "get:GetUser")
    web.Router("/user/:id", &controllers.UserController{}, "put:UpdateUser")
    web.Router("/user/:id", &controllers.UserController{}, "delete:DeleteUser")
    web.Router("/users", &controllers.UserController{}, "get:ListUsers")
    web.Router("/user/:id/visits", &controllers.UserVisitLogController{}, "get:GetUserVisits")

    // API routes
    // ItemUnit routes using the initialized controller instance
    web.Router("/api/item-units",  &controllers.ItemUnitController{}, "post:Create")
    web.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "get:Get")
    web.Router("/api/item-units",  &controllers.ItemUnitController{}, "get:List")
    web.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "put:Update")
    web.Router("/api/item-units/:id",  &controllers.ItemUnitController{}, "delete:Delete")
    web.Router("/api/item-units/serial/:serialNumber",  &controllers.ItemUnitController{}, "get:GetBySerialNumber")
    web.Router("/api/item-units/warehouse/:warehouseId", &controllers.ItemUnitController{}, "get:GetByWarehouse")

    // Item routes using the initialized controller instance
    web.Router("/api/item", &controllers.ItemController{}, "post:CreateItem")
    web.Router("/api/item/:id", &controllers.ItemController{}, "get:GetItem;put:UpdateItem;delete:DeleteItem")
    web.Router("/api/items", &controllers.ItemController{}, "get:ListItems")
    web.Router("/api/items/dashboard", &controllers.ItemController{}, "get:SearchDashboard")
    //item image route
    web.Router("/api/item/:id/image", &controllers.ItemController{}, "get:GetItemImage")

    // Item Category Routes
    categoryCtrl := controllers.NewItemCategoryController()
    web.Router("/api/categories", categoryCtrl, "get:List")
    web.Router("/api/categories/:id", categoryCtrl, "get:Get")
    web.Router("/api/categories", categoryCtrl, "post:Create")
    web.Router("/api/categories/:id", categoryCtrl, "put:Update")
    web.Router("/api/categories/:id", categoryCtrl, "delete:Delete")

    employeeCtrl := controllers.NewEmployeeController()
    // Employee Routes
    web.Router("/api/employees", employeeCtrl, "post:Create")
    web.Router("/api/employees/:id", employeeCtrl, "get:Get")
    web.Router("/api/employees", employeeCtrl, "get:List")
    web.Router("/api/employees/:id", employeeCtrl, "put:Update")
    web.Router("/api/employees/:id", employeeCtrl, "delete:Delete")

    // Lending Routes
    lendingCtrl := &controllers.LendingController{}
    web.Router("/api/lendings", lendingCtrl, "post:Create")
    web.Router("/api/lendings/:id", lendingCtrl, "get:Get")
    web.Router("/api/lendings", lendingCtrl, "get:List")
    web.Router("/api/lendings/:id", lendingCtrl, "put:Update")
    web.Router("/api/lendings/:id", lendingCtrl, "delete:Delete")
    web.Router("/api/lendings/report/items", &controllers.LendingController{}, "get:SearchItemReport")
    web.Router("/api/lendings/report/units", &controllers.LendingController{}, "get:SearchUnitReport")
    web.Router("/api/lendings/return/:id", &controllers.LendingController{}, "put:Return")
    // Lending images routes
    web.Router("/api/lending/:id/loan-image", &controllers.LendingController{}, "get:GetLoanImage")
    web.Router("/api/lending/:id/return-image", &controllers.LendingController{}, "get:GetReturnImage")

    // Additional lending-specific routes
    web.Router("/api/lendings/active", lendingCtrl, "get:GetActiveLoans")
    web.Router("/api/lendings/returned", lendingCtrl, "get:GetReturnedLoans")
    
    // Unit Log Routes
    web.Router("/api/unit-logs", &controllers.UnitLogController{}, "post:Create;get:List")
    web.Router("/api/unit-logs/:id", &controllers.UnitLogController{}, "get:Get;put:Update;delete:Delete")
    web.Router("/api/unit-logs/unit/:unitId", &controllers.UnitLogController{}, "get:GetByUnit")

    // Repair Log Routes
    web.Router("/api/repair-logs", &controllers.RepairLogController{}, "post:Create;get:List")
    web.Router("/api/repair-logs/:id", &controllers.RepairLogController{}, "get:Get;put:Update;delete:Delete")
    web.Router("/api/repair-logs/unit/:unitId", &controllers.RepairLogController{}, "get:GetByUnit")
    web.Router("/api/repair-logs/finish", &controllers.RepairLogController{}, "post:Finish")

    // Document Upload Routes
    web.Router("/api/docs", &controllers.DocUploadedController{}, "post:Create;get:List")
    web.Router("/api/docs/:id", &controllers.DocUploadedController{}, "get:Get;delete:Delete")
    // Template download
    web.Router("/api/docs/template/download", &controllers.DocUploadedController{}, "get:DownloadTemplate")
    // Upload Document
    web.Router("/api/docs/upload", &controllers.DocUploadedController{}, "post:Upload")


    // Warehouse Routes
    web.Router("/api/warehouse", &controllers.WarehouseController{}, "post:CreateWarehouse")
    web.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "get:GetWarehouse")
    web.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "put:UpdateWarehouse")
    web.Router("/api/warehouse/:id", &controllers.WarehouseController{}, "delete:DeleteWarehouse")
    web.Router("/api/warehouses", &controllers.WarehouseController{}, "get:ListWarehouses")
    web.Router("/api/warehouse/:id/users", &controllers.WarehouseController{}, "get:GetWarehouseUsers")

    // Status Lookup Routes
    web.Router("/api/status/:id", &controllers.StatusLookupController{}, "get:GetStatus")
    web.Router("/api/statuses", &controllers.StatusLookupController{}, "get:ListStatuses")

    //condition lookup routes
    web.Router("/api/condition/:id", &controllers.ConditionLookupController{}, "get:GetCondition")
    web.Router("/api/conditions", &controllers.ConditionLookupController{}, "get:ListConditions")
}