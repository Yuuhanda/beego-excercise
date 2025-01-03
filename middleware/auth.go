package middleware

import (
    "github.com/beego/beego/v2/server/web/context"
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "log"
)

func AuthMiddleware() web.FilterFunc {
    return func(ctx *context.Context) {
        // Get token from header
        token := ctx.Input.Header("Authorization")
        if token == "" {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Unauthorized access",
            }, true, false)
            return
        }

        // Get user from token
        authService := services.NewAuthService()
        user, err := authService.GetUserFromToken(token)
        if err != nil {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Invalid token",
            }, true, false)
            return
        }

        // Store user in context for later use
        ctx.Input.SetData("user", user)

        // Get user roles
        roleUserService := services.NewAuthRolesUserService()
        roles, err := roleUserService.GetRolesByUserId(user.Id)
        if err != nil {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Failed to get user roles",
            }, true, false)
            return
        }

        // Add after getting roles
        log.Printf("User Roles: %+v", roles)

        // Check if any role has permission for current path
        currentPath := ctx.Request.URL.Path
        currentMethod := ctx.Input.Method()
        hasPermission := false

        // Add before permission check
        log.Printf("Checking path: %s method: %s", currentPath, currentMethod)

        authItemService := services.NewAuthItemService()
        permissionChecks := []map[string]interface{}{}
        
        for _, role := range roles {
            permitted, _ := authItemService.CheckPermission(role.Code, currentPath, currentMethod)
            permissionChecks = append(permissionChecks, map[string]interface{}{
                "role": role.Code,
                "path": currentPath,
                "method": currentMethod,
                "permitted": permitted,
            })
            if permitted {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Access denied",
                "permission_checks": permissionChecks,
            }, true, false)
            return
        }    }
}