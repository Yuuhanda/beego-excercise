package middleware

import (
    "github.com/beego/beego/v2/server/web/context"
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "net/http"
)

func AuthMiddleware() web.FilterFunc {
    return func(ctx *context.Context) {
        authToken := ctx.Input.Header("Authorization")
        if authToken == "" {
            ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Authentication required",
            }, true, false)
            return
        }

        userService := services.NewUserService()
        user, err := userService.GetByAuthKey(authToken)
        if err != nil {
            ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Invalid authentication token",
            }, true, false)
            return
        }

        ctx.Input.SetData("user", user)
    }
}