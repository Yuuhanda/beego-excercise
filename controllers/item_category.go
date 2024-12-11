package controllers

import (
    "myproject/services"
)

type ItemCategoryController struct {
    categoryService *services.ItemCategoryService
}

func NewItemCategoryController() *ItemCategoryController {
    return &ItemCategoryController{
        categoryService: services.NewItemCategoryService(),
    }
}
