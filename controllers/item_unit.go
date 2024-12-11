package controllers

import (
    "myproject/services"
)

type ItemUnitController struct {
    // your controller struct fields
    itemUnitService *services.ItemUnitService
}

func NewItemUnitController() *ItemUnitController {
    return &ItemUnitController{
        itemUnitService: services.NewItemUnitService(),
    }
}