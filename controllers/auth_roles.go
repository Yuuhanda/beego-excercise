package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "myproject/models"
    "fmt"
	"encoding/json"
)

type AuthRolesController struct {
    web.Controller
    roleService *services.AuthRolesService
}

func (c *AuthRolesController) Prepare() {
    c.roleService = services.NewAuthRolesService()
}

func (c *AuthRolesController) Create() {
    var role models.AuthRoles
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&role); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := c.roleService.Create(&role); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create role",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role created successfully",
            "data":    role,
        }
    }
    c.ServeJSON()
}


func (c *AuthRolesController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    fmt.Sscanf(idStr, "%d", &id)

    role, err := c.roleService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Role not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    role,
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("page_size", 10)

    roles, total, err := c.roleService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve roles",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data": map[string]interface{}{
                "roles": roles,
                "total": total,
                "page":  page,
            },
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesController) Update() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    fmt.Sscanf(idStr, "%d", &id)

    var role models.AuthRoles
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&role); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    role.Id = id
    if err := c.roleService.Update(&role); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create role",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role updated successfully",
            "data":    role,
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    fmt.Sscanf(idStr, "%d", &id)

    // Check if role exists first
    role, err := c.roleService.GetByID(id)
    if err != nil || role == nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "No roles associated with id found",
        }
        c.ServeJSON()
        return
    }

    if err := c.roleService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete role",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role deleted successfully",
        }
    }
    c.ServeJSON()
}