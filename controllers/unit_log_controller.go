package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "time"
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
    var request struct {
        IdUnit       uint      `json:"IdUnit"`
        Content      string    `json:"Content"`
        ActorsAction string    `json:"ActorsAction"`
        UpdateAt     time.Time `json:"UpdateAt,omitempty"`
    }

    err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: request.IdUnit},
        Content:      request.Content,
        ActorsAction: request.ActorsAction,
        UpdateAt:     time.Now(),
    }

    if request.UpdateAt.IsZero() {
        unitLog.UpdateAt = time.Now()
    } else {
        unitLog.UpdateAt = request.UpdateAt
    }

    if err := c.unitLogService.Create(unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create unit log",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Unit log created successfully",
        "data":    unitLog,
    }
    c.ServeJSON()
}


// Get retrieves a unit log by ID
// @router /unit-logs/:id [get]
func (c *UnitLogController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    o := orm.NewOrm()
    unitLog := &models.UnitLog{IdLog: id}

    err := o.QueryTable(unitLog).Filter("IdLog", id).RelatedSel().One(unitLog)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
        c.ServeJSON()
        return
    }

    // Load related data
    if unitLog.IdUnit != nil {
        o.LoadRelated(unitLog.IdUnit, "Item")
        o.LoadRelated(unitLog.IdUnit, "StatusLookup")
        o.LoadRelated(unitLog.IdUnit, "Warehouse")
        o.LoadRelated(unitLog.IdUnit, "CondLookup")
        o.LoadRelated(unitLog.IdUnit, "User")

        if unitLog.IdUnit.Item != nil {
            o.LoadRelated(unitLog.IdUnit.Item, "Category")
        }
    }

    c.Data["json"] = map[string]interface{}{
        "data": unitLog,
    }
    c.ServeJSON()
}


// GetByUnit retrieves all logs for a specific unit
// @router /unit-logs/unit/:unitId [get]
func (c *UnitLogController) GetByUnit() {
    unitIDStr := c.Ctx.Input.Param(":unitId")
    unitID, _ := strconv.Atoi(unitIDStr)

    o := orm.NewOrm()
    var unitLogs []*models.UnitLog

    _, err := o.QueryTable(new(models.UnitLog)).Filter("IdUnit__IdUnit", unitID).RelatedSel().All(&unitLogs)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
        c.ServeJSON()
        return
    }

    // Load related data
    for _, unitLog := range unitLogs {
        if unitLog.IdUnit != nil {
            o.LoadRelated(unitLog.IdUnit, "Item")
            o.LoadRelated(unitLog.IdUnit, "StatusLookup")
            o.LoadRelated(unitLog.IdUnit, "Warehouse")
            o.LoadRelated(unitLog.IdUnit, "CondLookup")
            o.LoadRelated(unitLog.IdUnit, "User")

            if unitLog.IdUnit.Item != nil {
                o.LoadRelated(unitLog.IdUnit.Item, "Category")
            }
        }
    }

    c.Data["json"] = map[string]interface{}{
        "data": unitLogs,
    }
    c.ServeJSON()
}


// List retrieves unit logs with pagination
// @router /unit-logs [get]
func (c *UnitLogController) List() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "10"))

    filters := make(map[string]string)
    filters["serial_number"] = c.GetString("serial_number")
    filters["content"] = c.GetString("content")
    filters["start_date"] = c.GetString("start_date")
    filters["end_date"] = c.GetString("end_date")

    logs, total, err := c.unitLogService.List(page, pageSize, filters)
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

    var request struct {
        IdUnit       uint      `json:"IdUnit"`
        Content      string    `json:"Content"`
        ActorsAction string    `json:"ActorsAction"`
        UpdateAt     time.Time `json:"UpdateAt,omitempty"`
    }

    err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    unitLog := &models.UnitLog{
        IdLog:        id,
        IdUnit:       &models.ItemUnit{IdUnit: request.IdUnit},
        Content:      request.Content,
        ActorsAction: request.ActorsAction,
        UpdateAt:     time.Now(),
    }

    if err := c.unitLogService.Update(unitLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update unit log",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Unit log updated successfully",
        "data":    unitLog,
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
