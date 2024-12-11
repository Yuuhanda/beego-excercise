package controllers

import (
    "myproject/services"
)

type ItemController struct {
    // your controller struct fields
    itemService *services.ItemService
}

func NewItemController() *ItemController {
    return &ItemController{
        itemService: services.NewItemService(),
    }
}
