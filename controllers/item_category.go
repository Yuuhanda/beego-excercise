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
    categoryService services.ItemCategoryService
}

// NewItemCategoryController creates and initializes a new ItemCategoryController
func NewItemCategoryController() *ItemCategoryController {
    controller := &ItemCategoryController{}
    controller.categoryService = *services.NewItemCategoryService()
    return controller
}

// Create handles the creation of a new item category
func (c *ItemCategoryController) Create() {
    // Read the request body
    body := c.Ctx.Input.CopyBody(1048576) // Read up to 1MB
    
    var input struct {
        CategoryName string `json:"category_name"`
        CatCode     string `json:"cat_code"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    category := &models.ItemCategory{
        CategoryName: input.CategoryName,
        CatCode:     input.CatCode,
    }

    service := services.NewItemCategoryService()
    if err := service.Create(category); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create category",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Category created successfully",
            "data":    category,
        }
    }
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

    service := services.NewItemCategoryService()
    category, err := service.GetByID(id)
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
    catName := c.GetString("name")
    catCode := c.GetString("code")

    service := services.NewItemCategoryService()
    categories, total, err := service.List(page, pageSize, catName, catCode)
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
    // Get ID from URL parameter
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Read the request body
    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        CategoryName string `json:"category_name,omitempty"`
        CatCode     string `json:"cat_code,omitempty"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    service := services.NewItemCategoryService()
    
    // Get existing category using ID from URL
    existingCategory, err := service.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Category not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Update only provided fields
    if input.CategoryName != "" {
        existingCategory.CategoryName = input.CategoryName
    }
    if input.CatCode != "" {
        existingCategory.CatCode = input.CatCode
    }

    if err := service.Update(existingCategory); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update category",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Category updated successfully",
            "data":    existingCategory,
        }
    }
    c.ServeJSON()
}



// Delete removes a category
func (c *ItemCategoryController) Delete() {
    // Get ID from URL parameter
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    service := services.NewItemCategoryService()
    
    // Check if category exists before deleting
    existingCategory, err := service.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Category not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := service.Delete(existingCategory); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete category",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Category deleted successfully",
        }
    }
    c.ServeJSON()
}
