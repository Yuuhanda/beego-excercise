package controllers

import (
    "encoding/json"
    "myproject/models"
    "myproject/services"
    "strconv"
    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    "net/http"
    "time"
    "github.com/beego/beego/v2/core/logs"
)

type RepairLogController struct {
    beego.Controller
    repairLogService *services.RepairLogService
    userService      *services.UserService
    warehouseService *services.WarehouseService
    itemUnitService  *services.ItemUnitService
}

func (c *RepairLogController) Prepare() {
    c.repairLogService = services.NewRepairLogService()
    c.userService = services.NewUserService()
    c.warehouseService = services.NewWarehouseService()
    c.itemUnitService = services.NewItemUnitService()
}

// Create creates a new repair log
// @router /repair-logs [post]
func (c *RepairLogController) Create() {
    var request struct {
        IdUnit  uint   `json:"IdUnit"`
        Comment string `json:"Comment"`
        IdUser  uint   `json:"IdUser"`
    }

    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    repairLog := &models.RepairLog{
        IdUnit:   &models.ItemUnit{IdUnit: request.IdUnit},
        Comment:  request.Comment,
        RepType:  &models.RepTypeLookup{IdRepT: 1},
        Datetime: time.Now(),
    }

    if err := c.repairLogService.Create(repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create repair log",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Repair log created successfully",
        "data":    repairLog,
    }

    //get username that created the item unit
    user, err := c.userService.GetByID(int(request.IdUser))
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    //get serial number
    itemUnit, err := c.itemUnitService.Get(int(request.IdUnit))
    if err != nil {
        logs.Error("Failed to get item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get item unit",
        }
        c.ServeJSON()
        return
    }


    // Create UnitLog
    unitLogService := services.NewUnitLogService()

    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
        Content:      "Unit sent to repair. Info : " + request.Comment,
        ActorsAction: "Unit "+ itemUnit.SerialNumber +" sent to repair by " + user.Username,
        UpdateAt:     time.Now(),
    }


    err = unitLogService.Create(unitLog)
    if err != nil {
        // Handle error case
        logs.Error("Failed to create unit log: %v", err)
    }

    itemUnitService := services.NewItemUnitService()
    // Get existing item unit data first
    existingItemUnit, err := itemUnitService.Get(int(request.IdUnit))
    if err != nil {
        logs.Error("Failed to get existing item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get existing item unit",
        }
        c.ServeJSON()
        return
    }
    
    itemUnit = &models.ItemUnit{
        IdUnit:       uint(request.IdUnit),
        Comment:      existingItemUnit.Comment,
        StatusLookup: &models.StatusLookup{IdStatus: uint(3)},
        Warehouse:    &models.Warehouse{IdWh: uint(existingItemUnit.Warehouse.IdWh)},
        CondLookup:   &models.ConditionLookup{IdCondition: uint(existingItemUnit.CondLookup.IdCondition)},
        User:         &models.User{Id: existingItemUnit.User.Id},
    }

    err = itemUnitService.Update(itemUnit)
    if err != nil {
        logs.Error("Failed to update item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to update item unit",
        }
        c.ServeJSON()
        return
    }

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

    var request struct {
        Comment string `json:"Comment"`
        RepType uint   `json:"RepType"`
    }

    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    repairLog := &models.RepairLog{
        IdRepair: id,
        Comment:  request.Comment,
        RepType:  &models.RepTypeLookup{IdRepT: request.RepType},
    }

    if err := c.repairLogService.Update(repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update repair log",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Repair log updated successfully",
        "data":    repairLog,
    }
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


func (c *RepairLogController) Finish() {
    var request struct {
        IdUnit    uint   `json:"IdUnit"`
        Comment   string `json:"Comment"`
        IdUser    uint   `json:"IdUser"`
        Condition uint   `json:"Condition"` 
        IdWh      uint   `json:"IdWh"`
    }

    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    repairLog := &models.RepairLog{
        IdUnit:   &models.ItemUnit{IdUnit: request.IdUnit},
        Comment:  request.Comment,
        RepType:  &models.RepTypeLookup{IdRepT: 1},
        Datetime: time.Now(),
    }

    if err := c.repairLogService.Create(repairLog); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create repair log",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    //update item unit
    itemUnitService := services.NewItemUnitService()
    existingItemUnit, err := itemUnitService.Get(int(request.IdUnit))
    if err != nil {
        logs.Error("Failed to get existing item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get existing item unit",
        }
        c.ServeJSON()
        return
    }

    updatedItemUnit := &models.ItemUnit{
        IdUnit:       uint(request.IdUnit),
        Comment:      existingItemUnit.Comment,
        StatusLookup: &models.StatusLookup{IdStatus: uint(3)},
        Warehouse:    &models.Warehouse{IdWh: uint(existingItemUnit.Warehouse.IdWh)},
        CondLookup:   &models.ConditionLookup{IdCondition: request.Condition},
        User:         &models.User{Id: existingItemUnit.User.Id},
    }

    err = itemUnitService.Update(updatedItemUnit)
    if err != nil {
        logs.Error("Failed to update item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to update item unit",
        }
        c.ServeJSON()
        return
    }

    user, err := c.userService.GetByID(int(request.IdUser))
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    var warehouseName string
    if request.IdWh != 0 {
        warehouse, err := c.warehouseService.GetByID(request.IdWh)
        if err != nil {
            logs.Error("Failed to get warehouse: %v", err)
            c.Data["json"] = map[string]interface{}{
                "error": "Failed to get warehouse",
            }
            c.ServeJSON()
            return
        }
        warehouseName = warehouse.WhName
    } else {
        warehouseName = existingItemUnit.Warehouse.WhName
    }

    unitLogService := services.NewUnitLogService()
    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: existingItemUnit.IdUnit},
        Content:      "Unit repaired. Info: " + request.Comment,
        ActorsAction: "Unit " + existingItemUnit.SerialNumber + " returned to " + warehouseName + " by " + user.Username,
        UpdateAt:     time.Now(),
    }

    err = unitLogService.Create(unitLog)
    if err != nil {
        logs.Error("Failed to create unit log: %v", err)
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Repair log created successfully",
        "data":    repairLog,
    }
    c.ServeJSON()
}
