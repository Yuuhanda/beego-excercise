package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
)

type UserVisitLogController struct {
    web.Controller
    visitLogService *services.UserVisitLogService
}

func (c *UserVisitLogController) Prepare() {
    c.visitLogService = services.NewUserVisitLogService()
}

// GetUserVisits retrieves visit logs for a user
// @router /user/:id/visits [get]
func (c *UserVisitLogController) GetUserVisits() {
    userID, _ := c.GetInt(":id")
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("page_size", 10)
    
    logs, total, err := c.visitLogService.GetUserVisits(userID, page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve visit logs",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data": map[string]interface{}{
                "logs":  logs,
                "total": total,
                "page":  page,
            },
        }
    }
    c.ServeJSON()
}
