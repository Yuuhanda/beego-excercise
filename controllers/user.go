package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
	"fmt"
)

type UserController struct {
    web.Controller
    userService *services.UserService
}

func (c *UserController) Prepare() {
    c.userService = services.NewUserService()
}

// CreateUser handles user creation
// @router /user [post]
func (c *UserController) CreateUser() {
    var user models.User
    if err := c.ParseForm(&user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := c.userService.Create(&user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User created successfully",
            "data":    user,
        }
    }
    c.ServeJSON()
}

// GetUser retrieves user by ID
// @router /user/:id [get]
func (c *UserController) GetUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    user, err := c.userService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "User not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    user,
        }
    }
    c.ServeJSON()
}

// ListUsers retrieves paginated users
// @router /users [get]
func (c *UserController) ListUsers() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("page_size", 10)

    users, total, err := c.userService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve users",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data": map[string]interface{}{
                "users": users,
                "total": total,
                "page":  page,
            },
        }
    }
    c.ServeJSON()
}

// UpdateUser updates user information
// @router /user/:id [put]
func (c *UserController) UpdateUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    var user models.User
    if err := c.ParseForm(&user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    user.Id = id
    if err := c.userService.Update(&user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User updated successfully",
            "data":    user,
        }
    }
    c.ServeJSON()
}

// DeleteUser deletes a user
// @router /user/:id [delete]
func (c *UserController) DeleteUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    if err := c.userService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User deleted successfully",
        }
    }
    c.ServeJSON()
}
