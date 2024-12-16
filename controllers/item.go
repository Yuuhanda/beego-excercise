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
    itemCategoryService *services.ItemCategoryService
}

func (c *ItemController) Prepare() {
    c.itemService = services.NewItemService()
    c.itemCategoryService = services.NewItemCategoryService()
}

// CreateItem handles item creation
// @router /item [post]

func (c *ItemController) CreateItem() {
    // Read the request body
    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        ItemName    string `json:"ItemName"`
        SKU         string `json:"SKU"`
        Imagefile   string `json:"Imagefile"`
        CategoryId  uint   `json:"CategoryId"` // Changed to uint to match model
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

    // First verify the category exists using itemService
    category, err := c.itemCategoryService.GetByID(int(input.CategoryId))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid category",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Create item with verified category
    item := &models.Item{
        ItemName:  input.ItemName,
        SKU:       input.SKU,
        Imagefile: input.Imagefile,
        Category:  category,
    }

    // Generate SKU if empty
    if input.SKU == "" {
        if err := c.itemService.GenerateSKU(item); err != nil {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": "Failed to generate SKU",
                "error":   err.Error(),
            }
            c.ServeJSON()
            return
        }
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

    // Verify category if provided
    if input.CategoryId != nil {
        if _, err := c.itemCategoryService.GetByID(*input.CategoryId); err != nil {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": "Invalid category",
                "error":   err.Error(),
            }
            c.ServeJSON()
            return
        }
        item.Category = &models.ItemCategory{IdCategory: *input.CategoryId}
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

    // Generate SKU if provided SKU is empty string
    if input.SKU != nil && *input.SKU == "" {
        if err := c.itemService.GenerateSKU(item); err != nil {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": "Failed to generate SKU",
                "error":   err.Error(),
            }
            c.ServeJSON()
            return
        }
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

func (c *ItemController) SearchDashboard() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "20"))
    warehouseId, _ := strconv.Atoi(c.GetString("warehouse_id", "0"))

    filters := make(map[string]interface{})
    filters["item_name"] = c.GetString("item_name")
    filters["SKU"] = c.GetString("SKU")
    filters["category"] = c.GetString("category")

    items, total, err := c.itemService.SearchDashboard(page, pageSize, filters, warehouseId)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to search items",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "items": items,
            "total": total,
            "page":  page,
            "size":  pageSize,
        },
    }
    c.ServeJSON()
}
