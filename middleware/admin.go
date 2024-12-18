package middleware

import (
    "github.com/beego/beego/v2/server/web/context"
    "github.com/beego/beego/v2/server/web"
    "net/http"
	"myproject/models"
)

func AdminMiddleware() web.FilterFunc {
    return func(ctx *context.Context) {
        user := ctx.Input.GetData("user")
        if user == nil {
            ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Authentication required",
            }, true, false)
            return
        }

        if user.(*models.User).Superadmin != 1 {
            ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Admin access required",
            }, true, false)
            return
        }
    }
}
