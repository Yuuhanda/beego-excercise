package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "myproject/models"
    "encoding/json"
)

type AuthItemController struct {
    web.Controller
    authItemService *services.AuthItemService
}

func (c *AuthItemController) Prepare() {
    c.authItemService = services.NewAuthItemService()
}

func (c *AuthItemController) Create() {
    var authItem models.AuthItem
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&authItem); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := c.authItemService.Create(&authItem); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create auth item",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Auth item created successfully",
            "data":    authItem,
        }
    }
    c.ServeJSON()
}

func (c *AuthItemController) Get() {
    id, _ := c.GetInt(":id")
    authItem, err := c.authItemService.GetByID(id)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Auth item not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    authItem,
        }
    }
    c.ServeJSON()
}

func (c *AuthItemController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("page_size", 10)

    authItems, total, err := c.authItemService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve auth items",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data": map[string]interface{}{
                "items": authItems,
                "total": total,
                "page":  page,
            },
        }
    }
    c.ServeJSON()
}

func (c *AuthItemController) Update() {
    id, _ := c.GetInt(":id")
    var authItem models.AuthItem
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&authItem); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    authItem.Id = id
    if err := c.authItemService.Update(&authItem); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update auth item",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Auth item updated successfully",
            "data":    authItem,
        }
    }
    c.ServeJSON()
}

func (c *AuthItemController) Delete() {
    id, _ := c.GetInt(":id")
    
    err := c.authItemService.Delete(id)
    if err != nil {
        message := err.Error()
        if message == "no auth item found with this id" {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": message,
            }
        } else {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": "Failed to delete auth item",
                "error":   err.Error(),
            }
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Auth item deleted successfully",
        }
    }
    c.ServeJSON()
}

func (c *AuthItemController) CreateBulk() {
    var payload struct {
        Role  string   `json:"Role"`
        Paths []string `json:"Paths"`
    }
    
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&payload); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := c.authItemService.CreateBulk(payload.Role, payload.Paths); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create auth items",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Auth items created successfully",
        }
    }
    c.ServeJSON()
}

