package controllers

import (
    "encoding/json"
    "myproject/models"
    "myproject/services"
    "strconv"
    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "net/http"
)

type RepairLogController struct {
    beego.Controller
    repairLogService *services.RepairLogService
}

func (c *RepairLogController) Prepare() {
    c.repairLogService = services.NewRepairLogService()
}

// Create creates a new repair log
// @router /repair-logs [post]
func (c *RepairLogController) Create() {
    var repairLog models.RepairLog
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Invalid request body"}
        c.ServeJSON()
        return
    }

    if err := c.repairLogService.Create(&repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = repairLog
    c.ServeJSON()
}

// Get retrieves a repair log by ID
// @router /repair-logs/:id [get]
func (c *RepairLogController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    o := orm.NewOrm()
    repairLog, err := c.repairLogService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Repair log not found",
            "error":   err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
        c.ServeJSON()
        return
    }

    if repairLog.IdUnit != nil {
        o.LoadRelated(repairLog.IdUnit, "Item")
        o.LoadRelated(repairLog.IdUnit, "StatusLookup")
        o.LoadRelated(repairLog.IdUnit, "Warehouse")
        o.LoadRelated(repairLog.IdUnit, "CondLookup")
        o.LoadRelated(repairLog.IdUnit, "User")

        if repairLog.IdUnit.Item != nil {
            o.LoadRelated(repairLog.IdUnit.Item, "Category")
        }
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    repairLog,
    }
    c.ServeJSON()
}



// GetByUnit retrieves repair logs by unit ID
// @router /repair-logs/unit/:unitId [get]
func (c *RepairLogController) GetByUnit() {
    unitIDStr := c.Ctx.Input.Param(":unitId")
    unitID, _ := strconv.Atoi(unitIDStr)

    logs, err := c.repairLogService.GetByUnit(unitID)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = logs
    c.ServeJSON()
}

// List retrieves repair logs with pagination
// @router /repair-logs [get]
func (c *RepairLogController) List() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "10"))

    filters := make(map[string]string)
    filters["item_name"] = c.GetString("item_name")
    filters["serial_number"] = c.GetString("serial_number")
    filters["sku"] = c.GetString("sku")
    filters["warehouse"] = c.GetString("warehouse")
    filters["rep_type"] = c.GetString("rep_type")
    filters["start_date"] = c.GetString("start_date")
    filters["end_date"] = c.GetString("end_date")

    logs, total, err := c.repairLogService.List(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "data":  logs,
            "total": total,
            "page":  page,
        }
    }
    c.ServeJSON()
}



// Update updates a repair log
// @router /repair-logs/:id [put]
func (c *RepairLogController) Update() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    var repairLog models.RepairLog
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Invalid request body"}
        c.ServeJSON()
        return
    }

    repairLog.IdRepair = id
    if err := c.repairLogService.Update(&repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = repairLog
    c.ServeJSON()
}

// Delete deletes a repair log
// @router /repair-logs/:id [delete]
func (c *RepairLogController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    if err := c.repairLogService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{"message": "Successfully deleted"}
    c.ServeJSON()
}
