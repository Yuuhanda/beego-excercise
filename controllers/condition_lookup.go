package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "fmt"
)

type ConditionLookupController struct {
    web.Controller
    conditionService *services.ConditionLookupService
}

func (c *ConditionLookupController) Prepare() {
    c.conditionService = services.NewConditionLookupService()
}

// GetCondition retrieves condition by ID
// @router /condition/:id [get]
func (c *ConditionLookupController) GetCondition() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid condition ID",
        }
        c.ServeJSON()
        return
    }

    condition, err := c.conditionService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Condition not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    condition,
        }
    }
    c.ServeJSON()
}

// ListConditions retrieves all conditions
// @router /conditions [get]
func (c *ConditionLookupController) ListConditions() {
    conditions, err := c.conditionService.GetAll()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve conditions",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    conditions,
        }
    }
    c.ServeJSON()
}
