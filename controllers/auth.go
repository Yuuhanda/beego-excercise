package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
    web.Controller
    userService *services.UserService
    visitLogService *services.UserVisitLogService
}

func (c *AuthController) Prepare() {
    c.userService = services.NewUserService()
    c.visitLogService = services.NewUserVisitLogService()
}

// Login handles user authentication
// @router /auth/login [post]
func (c *AuthController) Login() {
    var loginForm struct {
        Email    string `form:"email"`
        Password string `form:"password"` // Raw password
    }

    if err := c.ParseForm(&loginForm); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
        }
        c.ServeJSON()
        return
    }

    // Get user by email and verify password
    user, err := c.userService.GetByEmail(loginForm.Email)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid credentials",
        }
        c.ServeJSON()
        return
    }

    // Verify password using bcrypt
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password)); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid credentials",
        }
        c.ServeJSON()
        return
    }

    // Generate auth token
    if err := c.userService.GenerateAuthToken(user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to generate auth token",
        }
        c.ServeJSON()
        return
    }

    // Log the visit
    c.visitLogService.LogVisit(
        user.Id,
        c.Ctx.Input.IP(),
        c.Ctx.Input.UserAgent(),
        c.Ctx.Input.Header("Accept-Language"),
    )

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Login successful",
        "data": map[string]interface{}{
            "user": user,
            "token": user.AuthKey,
        },
    }
    c.ServeJSON()
}

// Logout handles user logout
// @router /auth/logout [post]
func (c *AuthController) Logout() {
    // Clear auth token from current user
    authHeader := c.Ctx.Input.Header("Authorization")
    if user, err := c.userService.GetByAuthKey(authHeader); err == nil {
        user.AuthKey = ""
        c.userService.Update(user)
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Logout successful",
    }
    c.ServeJSON()
}
