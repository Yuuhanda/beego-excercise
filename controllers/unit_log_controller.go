package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
)

type UnitLogController struct {
    web.Controller
    unitLogService *services.UnitLogService
}

func (c *UnitLogController) Prepare() {
    c.unitLogService = services.NewUnitLogService()
}

// Create creates a new unit log
// @router /unit-logs [post]
func (c *UnitLogController) Create() {
    var unitLog models.UnitLog
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.unitLogService.Create(&unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
    } else {
        c.Data["json"] = unitLog
    }
    c.ServeJSON()
}

// Get retrieves a unit log by ID
// @router /unit-logs/:id [get]
func (c *UnitLogController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    unitLog, err := c.unitLogService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
    } else {
        c.Data["json"] = unitLog
    }
    c.ServeJSON()
}

// GetByUnit retrieves all logs for a specific unit
// @router /unit-logs/unit/:unitId [get]
func (c *UnitLogController) GetByUnit() {
    unitIDStr := c.Ctx.Input.Param(":unitId")
    unitID, _ := strconv.Atoi(unitIDStr)

    logs, err := c.unitLogService.GetByUnitID(unitID)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
    } else {
        c.Data["json"] = logs
    }
    c.ServeJSON()
}

// List retrieves unit logs with pagination
// @router /unit-logs [get]
func (c *UnitLogController) List() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "10"))

    logs, total, err := c.unitLogService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
    } else {
        c.Data["json"] = map[string]interface{}{
            "data":  logs,
            "total": total,
            "page":  page,
            "size":  pageSize,
        }
    }
    c.ServeJSON()
}

// Update updates a unit log
// @router /unit-logs/:id [put]
func (c *UnitLogController) Update() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    var unitLog models.UnitLog
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    unitLog.IdLog = id
    if err := c.unitLogService.Update(&unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
    } else {
        c.Data["json"] = unitLog
    }
    c.ServeJSON()
}

// Delete deletes a unit log
// @router /unit-logs/:id [delete]
func (c *UnitLogController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    if err := c.unitLogService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
    } else {
        c.Data["json"] = map[string]interface{}{
            "message": "Unit log deleted successfully",
        }
    }
    c.ServeJSON()
}
