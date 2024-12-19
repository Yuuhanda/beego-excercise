package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "myproject/models"
    "encoding/json"
)

type AuthRolesUserController struct {
    web.Controller
    roleUserService *services.AuthRolesUserService
}

func (c *AuthRolesUserController) Prepare() {
    c.roleUserService = services.NewAuthRolesUserService()
}

func (c *AuthRolesUserController) Create() {
    var payload struct {
        UserId int `json:"UserId"`
        RoleId int `json:"RoleId"`
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

    roleUser := &models.AuthRolesUser{
        UserId: &models.User{Id: payload.UserId},
        RoleId: &models.AuthRoles{Id: payload.RoleId},
    }

    if err := c.roleUserService.Create(roleUser); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to assign role to user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role assigned to user successfully",
            "data":    roleUser,
        }
    }
    c.ServeJSON()
}
func (c *AuthRolesUserController) GetUserRoles() {
    userId, _ := c.GetInt(":userId")
    roles, err := c.roleUserService.GetRolesByUserId(userId)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get user roles",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    roles,
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesUserController) GetRoleUsers() {
    roleId, _ := c.GetInt(":roleId")
    users, err := c.roleUserService.GetUsersByRoleId(roleId)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get role users",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    users,
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesUserController) Delete() {
    userId, _ := c.GetInt(":userId")
    roleId, _ := c.GetInt(":roleId")

    if err := c.roleUserService.Delete(userId, roleId); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to remove role from user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role removed from user successfully",
        }
    }
    c.ServeJSON()
}
