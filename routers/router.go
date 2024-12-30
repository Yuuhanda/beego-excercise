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

    //RBAC
    // Routes Scanner
    web.Router("/api/routes/scan", &controllers.APIRouteController{}, "post:ScanRoutes")
    web.Router("/api/routes/list", &controllers.APIRouteController{}, "get:ListRoutes")
    web.Router("/api/routes/:id", &controllers.APIRouteController{}, "get:Get")
    web.Router("/api/routes/:id", &controllers.APIRouteController{}, "delete:DeleteRoute")
    // Auth Roles Routes
    web.Router("/api/roles", &controllers.AuthRolesController{}, "post:Create")
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "get:Get")
    web.Router("/api/roles", &controllers.AuthRolesController{}, "get:List")
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "put:Update")
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "delete:Delete")
    // Auth Roles User Routes
    web.Router("/api/user-roles", &controllers.AuthRolesUserController{}, "post:Create")
    web.Router("/api/user-roles/user/:userId", &controllers.AuthRolesUserController{}, "get:GetUserRoles")
    web.Router("/api/user-roles/role/:roleId", &controllers.AuthRolesUserController{}, "get:GetRoleUsers")
    web.Router("/api/user-roles/:userId/:roleId", &controllers.AuthRolesUserController{}, "delete:Delete")
    // Auth Item Routes
    web.Router("/api/auth-items", &controllers.AuthItemController{}, "post:Create")
    web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "get:Get")
    web.Router("/api/auth-items", &controllers.AuthItemController{}, "get:List")
    web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "put:Update")
    web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "delete:Delete")
    web.Router("/api/auth-items/bulk", &controllers.AuthItemController{}, "post:CreateBulk")

    // Access Control Routes
    web.InsertFilter("/api/routes/*", web.BeforeRouter, middleware.AuthMiddleware())
    web.InsertFilter("/api/roles/*", web.BeforeRouter, middleware.AuthMiddleware())
    web.InsertFilter("/api/*", web.BeforeRouter, middleware.AuthMiddleware())

    // User Routes
    web.Router("/auth/logout", &controllers.AuthController{}, "post:Logout")
    web.Router("/api/user/:id", &controllers.UserController{}, "get:GetUser")
    web.Router("/api/user/:id", &controllers.UserController{}, "put:UpdateUser")
    web.Router("/api/user/:id", &controllers.UserController{}, "delete:DeleteUser")
    web.Router("/api/users", &controllers.UserController{}, "get:ListUsers")
    web.Router("/api/user/visits/:id", &controllers.UserVisitLogController{}, "get:GetUserVisits")
    web.Router("/api/user", &controllers.UserController{}, "post:CreateUser")

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
    web.Router("/api/item/:id", &controllers.ItemController{}, "get:GetItem")
    web.Router("/api/item/:id", &controllers.ItemController{}, "put:UpdateItem")
    web.Router("/api/item/:id", &controllers.ItemController{}, "delete:DeleteItem")
    web.Router("/api/items", &controllers.ItemController{}, "get:ListItems")
    web.Router("/api/items/dashboard", &controllers.ItemController{}, "get:SearchDashboard")
    //item image route
    web.Router("/api/item/:id/image", &controllers.ItemController{}, "get:GetItemImage")

    // Item Category Routes
    web.Router("/api/categories", &controllers.ItemCategoryController{}, "get:List")
    web.Router("/api/categories/:id", &controllers.ItemCategoryController{}, "get:Get")
    web.Router("/api/categories", &controllers.ItemCategoryController{}, "post:Create")
    web.Router("/api/categories/:id", &controllers.ItemCategoryController{}, "put:Update")
    web.Router("/api/categories/:id", &controllers.ItemCategoryController{}, "delete:Delete")

    // Employee Routes
    web.Router("/api/employees", &controllers.EmployeeController{}, "post:Create")
    web.Router("/api/employees/:id", &controllers.EmployeeController{}, "get:Get")
    web.Router("/api/employees", &controllers.EmployeeController{}, "get:List")
    web.Router("/api/employees/:id", &controllers.EmployeeController{}, "put:Update")
    web.Router("/api/employees/:id", &controllers.EmployeeController{}, "delete:Delete")

    // Lending Routes
    web.Router("/api/lendings", &controllers.LendingController{}, "post:Create")
    web.Router("/api/lendings/:id", &controllers.LendingController{}, "get:Get")
    web.Router("/api/lendings", &controllers.LendingController{}, "get:List")
    web.Router("/api/lendings/:id", &controllers.LendingController{}, "put:Update")
    web.Router("/api/lendings/:id", &controllers.LendingController{}, "delete:Delete")
    web.Router("/api/lendings/report/items", &controllers.LendingController{}, "get:SearchItemReport")
    web.Router("/api/lendings/report/units", &controllers.LendingController{}, "get:SearchUnitReport")
    web.Router("/api/lendings/return/:id", &controllers.LendingController{}, "put:Return")
    // Lending images routes
    web.Router("/api/lending/:id/loan-image", &controllers.LendingController{}, "get:GetLoanImage")
    web.Router("/api/lending/:id/return-image", &controllers.LendingController{}, "get:GetReturnImage")
    // Additional lending-specific routes
    web.Router("/api/lendings/active", &controllers.LendingController{}, "get:GetActiveLoans")
    web.Router("/api/lendings/returned", &controllers.LendingController{}, "get:GetReturnedLoans")
    
    // Unit Log Routes
    web.Router("/api/unit-logs", &controllers.UnitLogController{}, "post:Create;get:List")
    web.Router("/api/unit-logs/:id", &controllers.UnitLogController{}, "get:Get")
    web.Router("/api/unit-logs/:id", &controllers.UnitLogController{}, "put:Update")
    web.Router("/api/unit-logs/:id", &controllers.UnitLogController{}, "delete:Delete")
    web.Router("/api/unit-logs/unit/:unitId", &controllers.UnitLogController{}, "get:GetByUnit")

    // Repair Log Routes
    web.Router("/api/repair-logs", &controllers.RepairLogController{}, "post:Create;get:List")
    web.Router("/api/repair-logs/:id", &controllers.RepairLogController{}, "get:Get")
    web.Router("/api/repair-logs/:id", &controllers.RepairLogController{}, "delete:Delete")
    web.Router("/api/repair-logs/:id", &controllers.RepairLogController{}, "put:Update")
    web.Router("/api/repair-logs/unit/:unitId", &controllers.RepairLogController{}, "get:GetByUnit")
    web.Router("/api/repair-logs/finish", &controllers.RepairLogController{}, "post:Finish")

    // Document Upload Routes
    web.Router("/api/docs", &controllers.DocUploadedController{}, "post:Create;get:List")
    web.Router("/api/docs", &controllers.DocUploadedController{}, "get:List")
    web.Router("/api/docs/:id", &controllers.DocUploadedController{}, "get:Get")
    web.Router("/api/docs/:id", &controllers.DocUploadedController{}, "delete:Delete")
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

    web.Router("/test/condition/:id", &controllers.ConditionLookupController{}, "get:GetCondition")
    web.Router("/test/conditions", &controllers.ConditionLookupController{}, "get:ListConditions")
}