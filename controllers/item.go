package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "fmt"
    "strconv"
    "encoding/json"
)

type ItemController struct {
    web.Controller
    itemService *services.ItemService
}

func (c *ItemController) Prepare() {
    c.itemService = services.NewItemService()
}

// CreateItem handles item creation
// @router /item [post]

func (c *ItemController) CreateItem() {
    // Read the request body
    body := c.Ctx.Input.CopyBody(1048576) // Read up to 1MB
    
    var input struct {
        ItemName    string `json:"ItemName"`
        SKU         string `json:"SKU"`
        Imagefile   string `json:"Imagefile"`
        CategoryId  int    `json:"CategoryId"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
            "received": string(body),  // This will show what was actually received
        }
        c.ServeJSON()
        return
    }

    item := &models.Item{
        ItemName:  input.ItemName,
        SKU:       input.SKU,
        Imagefile: input.Imagefile,
        Category:  &models.ItemCategory{IdCategory: input.CategoryId},
    }

    if err := c.itemService.Create(item); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create item",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Item created successfully",
            "data":    item,
        }
    }
    c.ServeJSON()
}


// GetItem retrieves item by ID
// @router /item/:id [get]
func (c *ItemController) GetItem() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid item ID",
        }
        c.ServeJSON()
        return
    }

    item, err := c.itemService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Item not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    item,
        }
    }
    c.ServeJSON()
}

// ListItems retrieves all items
// @router /items [get]
// ListItems retrieves all items with filters
func (c *ItemController) ListItems() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)
    
    // Get filter parameters
    filters := make(map[string]interface{})
    
    if name := c.GetString("name"); name != "" {
        filters["name"] = name
    }
    if category := c.GetString("category"); category != "" {
        filters["category"] = category
    }
    if sku := c.GetString("sku"); sku != "" {
        filters["sku"] = sku
    }
    if statusStr := c.GetString("status"); statusStr != "" {
        if status, err := strconv.ParseUint(statusStr, 10, 32); err == nil {
            filters["status"] = uint(status)
        }
    }

    items, total, err := c.itemService.List(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "items": items,
        "total": total,
        "page":  page,
        "size":  pageSize,
    }
    c.ServeJSON()
}

// UpdateItem updates an item
// @router /item/:id [put]
func (c *ItemController) UpdateItem() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid item ID",
        }
        c.ServeJSON()
        return
    }

    body := c.Ctx.Input.CopyBody(1048576)
    var input struct {
        ItemName    *string `json:"ItemName,omitempty"`
        SKU         *string `json:"SKU,omitempty"`
        Imagefile   *string `json:"Imagefile,omitempty"`
        CategoryId  *int    `json:"CategoryId,omitempty"`
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

    // Get existing item
    item, err := c.itemService.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Item not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Update only provided fields
    if input.ItemName != nil {
        item.ItemName = *input.ItemName
    }
    if input.SKU != nil {
        item.SKU = *input.SKU
    }
    if input.Imagefile != nil {
        item.Imagefile = *input.Imagefile
    }
    if input.CategoryId != nil {
        item.Category = &models.ItemCategory{IdCategory: *input.CategoryId}
    }

    if err := c.itemService.Update(item); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update item",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Item updated successfully",
            "data":    item,
        }
    }
    c.ServeJSON()
}


// DeleteItem deletes an item
// @router /item/:id [delete]
func (c *ItemController) DeleteItem() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid item ID",
        }
        c.ServeJSON()
        return
    }

    if err := c.itemService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete item",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Item deleted successfully",
        }
    }
    c.ServeJSON()
}