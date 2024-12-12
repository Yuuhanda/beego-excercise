package controllers

import (
    "encoding/json"
    "myproject/models"
    "myproject/services"
    "strconv"

    "github.com/beego/beego/v2/server/web"
)

type ItemCategoryController struct {
    web.Controller
    categoryService *services.ItemCategoryService
}

func NewItemCategoryController() *ItemCategoryController {
    return &ItemCategoryController{
        categoryService: services.NewItemCategoryService(),
    }
}

// Create handles the creation of a new item category
func (c *ItemCategoryController) Create() {
    var category models.ItemCategory
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.categoryService.Create(&category); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = category
    c.ServeJSON()
}

// Get retrieves a category by ID
func (c *ItemCategoryController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    category, err := c.categoryService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = category
    c.ServeJSON()
}

// List retrieves all categories with pagination
func (c *ItemCategoryController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)

    categories, total, err := c.categoryService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "data":  categories,
        "total": total,
        "page":  page,
    }
    c.ServeJSON()
}

// Update handles category updates
func (c *ItemCategoryController) Update() {
    var category models.ItemCategory
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.categoryService.Update(&category); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = category
    c.ServeJSON()
}

// Delete removes a category
func (c *ItemCategoryController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    if err := c.categoryService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "message": "Category deleted successfully",
    }
    c.ServeJSON()
}