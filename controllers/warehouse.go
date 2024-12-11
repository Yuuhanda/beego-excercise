package controllers

import (
    "myproject/services"
)


type WarehouseController struct {
    // ... other fields
    warehouseService *services.WarehouseService
}

func NewWarehouseController() *WarehouseController {
    return &WarehouseController{
        warehouseService: services.NewWarehouseService(),
    }
}
