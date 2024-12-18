package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "fmt"
)

type StatusLookupController struct {
    web.Controller
    statusService *services.StatusLookupService
}

func (c *StatusLookupController) Prepare() {
    c.statusService = services.NewStatusLookupService()
}

// GetStatus retrieves status by ID
// @router /status/:id [get]
func (c *StatusLookupController) GetStatus() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid status ID",
        }
        c.ServeJSON()
        return
    }

    status, err := c.statusService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Status not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    status,
        }
    }
    c.ServeJSON()
}

// ListStatuses retrieves all status lookups
// @router /statuses [get]
func (c *StatusLookupController) ListStatuses() {
    statuses, err := c.statusService.GetAll()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve statuses",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    statuses,
        }
    }
    c.ServeJSON()
}