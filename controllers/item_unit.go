package controllers

import (
	"encoding/json"
	"myproject/models"
	"myproject/services"
	"strconv"
	"time"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type ItemUnitController struct {
    web.Controller
    itemUnitService *services.ItemUnitService
    userService      *services.UserService
    warehouseService *services.WarehouseService
}

type UpdateItemUnitRequest struct {
    SerialNumber string `json:"SerialNumber"`
    Comment   string `json:"Comment"`
    Status    uint   `json:"Status"`
    IdWh      uint   `json:"IdWh"`
    Condition uint   `json:"Condition"`
    UpdatedBy int    `json:"UpdatedBy"`
}

func (c *ItemUnitController) Prepare() {
    c.userService = services.NewUserService()
    c.warehouseService = services.NewWarehouseService()
    c.itemUnitService = services.NewItemUnitService()
}

func NewItemUnitController() *ItemUnitController {
    controller := &ItemUnitController{}
    controller.itemUnitService = services.NewItemUnitService()
    return controller
}

// Create handles the creation of a new item unit
func (c *ItemUnitController) Create() {
    body := c.Ctx.Input.CopyBody(1048576) // Read up to 1MB
    logs.Info("Raw request body: %s", string(body))
    
    var input struct {
        SerialNumber string `json:"SerialNumber"`
        Comment      string `json:"Comment"`
        IdItem       uint   `json:"IdItem"`
        Status       uint   `json:"Status"`
        IdWh         uint   `json:"IdWh"`
        Condition    uint   `json:"Condition"`
        UpdatedBy    int    `json:"UpdatedBy"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        logs.Error("JSON unmarshal error: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
            "details": err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    // Generate serial number if empty
    if input.SerialNumber == "" {
        generatedSN, err := c.itemUnitService.GenerateSerialNumber(input.IdItem)
        if err != nil {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "message": "Failed to generate serial number",
                "error":   err.Error(),
            }
            c.ServeJSON()
            return
        }
        input.SerialNumber = generatedSN
    }

    itemUnit := &models.ItemUnit{
        SerialNumber: input.SerialNumber,
        Comment:      input.Comment,
        Item:         &models.Item{IdItem: input.IdItem},
        StatusLookup: &models.StatusLookup{IdStatus: input.Status},
        Warehouse:    &models.Warehouse{IdWh: input.IdWh},
        CondLookup:   &models.ConditionLookup{IdCondition: input.Condition},
        User:         &models.User{Id: input.UpdatedBy},
    }

    if err := c.itemUnitService.Create(itemUnit); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create item unit",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Item unit created successfully",
            "data":    itemUnit,
        }
    }
    
    user, err := c.userService.GetByID(input.UpdatedBy)
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    warehouse, err := c.warehouseService.GetByID(input.IdWh)
    if err != nil {
        logs.Error("Failed to get warehouse: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get warehouse",
        }
        c.ServeJSON()
        return
    }

    unitLogService := services.NewUnitLogService()

    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
        Content:      input.Comment,
        ActorsAction: "New Unit added by " + user.Username + " into " + warehouse.WhName,
        UpdateAt:     time.Now(),
    }

    err = unitLogService.Create(unitLog)
    if err != nil {
        logs.Error("Failed to create unit log: %v", err)
    }

    c.Data["json"] = itemUnit
    c.ServeJSON()
}




// Get retrieves an item unit by ID
func (c *ItemUnitController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
            "code": 400,
        }
        c.ServeJSON()
        return
    }

    o := orm.NewOrm()
    itemUnit := &models.ItemUnit{IdUnit: uint(id)}
    err = o.Read(itemUnit)
    
    if err == orm.ErrNoRows {
        c.Data["json"] = map[string]interface{}{
            "error": "Item unit not found",
            "code": 404,
        }
        c.ServeJSON()
        return
    }
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Database error",
            "code": 500,
        }
        c.ServeJSON()
        return
    }

    // Load all relationships including nested Category
    o.LoadRelated(itemUnit, "Item")
    if itemUnit.Item != nil {
        o.LoadRelated(itemUnit.Item, "Category")
    }
    o.LoadRelated(itemUnit, "StatusLookup")
    o.LoadRelated(itemUnit, "Warehouse")
    o.LoadRelated(itemUnit, "CondLookup")
    o.LoadRelated(itemUnit, "User")

    
    c.Data["json"] = itemUnit
    c.ServeJSON()
}

// List retrieves a paginated list of item units
func (c *ItemUnitController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)
    
    filters := make(map[string]interface{})
    
    if itemName := c.GetString("itemName"); itemName != "" {
        filters["itemName"] = itemName
    }
    if serialNumber := c.GetString("serialNumber"); serialNumber != "" {
        filters["serialNumber"] = serialNumber
    }
    if warehouseIdStr := c.GetString("warehouseId"); warehouseIdStr != "" {
        if warehouseId, err := strconv.ParseUint(warehouseIdStr, 10, 32); err == nil {
            filters["warehouseId"] = uint(warehouseId)
        }
    }
    if status := c.GetString("status"); status != "" {
        filters["status"] = status  // Pass the raw string for comma-separated values
    }
    if condition := c.GetString("condition"); condition != "" {
        filters["condition"] = condition  // Pass the raw string for comma-separated values
    }
    if userIdStr := c.GetString("userId"); userIdStr != "" {
        if userId, err := strconv.ParseUint(userIdStr, 10, 32); err == nil {
            filters["userId"] = uint(userId)
        }
    }

    items, total, err := c.itemUnitService.List(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "items": items,
        "total": total,
        "page":  page,
        "size":  pageSize,
    }
    c.ServeJSON()
}


  func (c *ItemUnitController) Update() {
      logs.Info("Starting update process...")
    
      idStr := c.Ctx.Input.Param(":id")
      id, err := strconv.ParseUint(idStr, 10, 32)
      if err != nil {
          c.Data["json"] = map[string]interface{}{
              "success": false,
              "message": "Invalid ID format",
          }
          c.ServeJSON()
          return
      }
    
      body := c.Ctx.Input.CopyBody(1048576)
      logs.Info("Raw request body: %s", string(body))

      var req UpdateItemUnitRequest
      if err := json.Unmarshal(body, &req); err != nil {
          c.Data["json"] = map[string]interface{}{
              "success": false,
              "message": "Invalid request format",
              "error":   err.Error(),
              "received": string(body),
          }
          c.ServeJSON()
          return
      }
    
      logs.Info("Getting existing unit with ID: %d", id)
      existingUnit, err := c.itemUnitService.Get(int(id))
      if err != nil {
          logs.Error("Failed to get existing unit: %v", err)
          c.Data["json"] = map[string]interface{}{
              "success": false,
              "message": "Failed to find item unit",
              "error":   err.Error(),
          }
          c.ServeJSON()
          return
      }
    
      logs.Info("Found existing unit: %+v", existingUnit)
    
      itemUnit := &models.ItemUnit{
          IdUnit:      uint(id),
          SerialNumber: req.SerialNumber,
          Comment:     req.Comment,
          Item:        existingUnit.Item, // Using the existing item
          StatusLookup: &models.StatusLookup{IdStatus: req.Status},
          Warehouse:   &models.Warehouse{IdWh: req.IdWh},
          CondLookup:  &models.ConditionLookup{IdCondition: req.Condition},
          User:        &models.User{Id: req.UpdatedBy},
      }

      // Generate serial number if empty
      if req.SerialNumber == "" {
          generatedSN, err := c.itemUnitService.GenerateSerialNumber(existingUnit.Item.IdItem)
          if err != nil {
              c.Data["json"] = map[string]interface{}{
                  "success": false,
                  "message": "Failed to generate serial number",
                  "error":   err.Error(),
              }
              c.ServeJSON()
              return
          }
          itemUnit.SerialNumber = generatedSN
      }

      if err := c.itemUnitService.Update(itemUnit); err != nil {
          c.Data["json"] = map[string]interface{}{
              "success": false,
              "message": "Failed to update item unit",
              "error":   err.Error(),
          }
      } else {
          c.Data["json"] = map[string]interface{}{
              "success": true,
              "message": "Item unit updated successfully",
              "data":    itemUnit,
          }
      }

      user, err := c.userService.GetByID(req.UpdatedBy)
      if err != nil {
          logs.Error("Failed to get user: %v", err)
          c.Data["json"] = map[string]interface{}{
              "error": "Failed to get user",
          }
          c.ServeJSON()
          return
      }

      warehouse, err := c.warehouseService.GetByID(req.IdWh)
      if err != nil {
          logs.Error("Failed to get warehouse: %v", err)
          c.Data["json"] = map[string]interface{}{
              "error": "Failed to get warehouse",
          }
          c.ServeJSON()
          return
      }

      unitLogService := services.NewUnitLogService()

      unitLog := &models.UnitLog{
          IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
          Content:      req.Comment,
          ActorsAction: "Unit " + itemUnit.SerialNumber + " updated by " + user.Username + " in " + warehouse.WhName,
          UpdateAt:     time.Now(),
      }

      err = unitLogService.Create(unitLog)
      if err != nil {
          logs.Error("Failed to create unit log: %v", err)
      }

      c.ServeJSON()
  }

// Delete removes an item unit
func (c *ItemUnitController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
            "code": 400,
        }
        c.ServeJSON()
        return
    }

    o := orm.NewOrm()
    itemUnit := &models.ItemUnit{IdUnit: uint(id)}
    
    if err := o.Read(itemUnit); err == orm.ErrNoRows {
        c.Data["json"] = map[string]interface{}{
            "error": "Item unit not found",
            "code": 404,
        }
        c.ServeJSON()
        return
    }

    _, err = o.Delete(itemUnit)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
            "code": 500,
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "message": "Item unit deleted successfully",
        "code": 200,
    }
    c.ServeJSON()
}



// GetBySerialNumber retrieves an item unit by serial number
func (c *ItemUnitController) GetBySerialNumber() {
    serialNumber := c.Ctx.Input.Param(":serialNumber")
    if serialNumber == "" {
        c.Data["json"] = map[string]interface{}{
            "error": "Serial number is required",
            "code": 400,
        }
        c.ServeJSON()
        return
    }

    o := orm.NewOrm()
    itemUnit := &models.ItemUnit{SerialNumber: serialNumber}
    err := o.Read(itemUnit, "SerialNumber")
    
    if err == orm.ErrNoRows {
        c.Data["json"] = map[string]interface{}{
            "error": "Item unit not found",
            "code": 404,
        }
        c.ServeJSON()
        return
    }

    // Load all relationships including nested Category
    o.LoadRelated(itemUnit, "Item")
    if itemUnit.Item != nil {
        o.LoadRelated(itemUnit.Item, "Category")
    }
    o.LoadRelated(itemUnit, "StatusLookup")
    o.LoadRelated(itemUnit, "Warehouse")
    o.LoadRelated(itemUnit, "CondLookup")
    o.LoadRelated(itemUnit, "User")
    
    c.Data["json"] = itemUnit
    c.ServeJSON()
}



// GetByWarehouse retrieves item units by warehouse ID
func (c *ItemUnitController) GetByWarehouse() {
    warehouseIdStr := c.Ctx.Input.Param(":warehouseId")
    warehouseId, err := strconv.ParseUint(warehouseIdStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid warehouse ID format",
            "code": 400,
        }
        c.ServeJSON()
        return
    }

    o := orm.NewOrm()
    var items []*models.ItemUnit
    
    qs := o.QueryTable(new(models.ItemUnit)).Filter("Warehouse__IdWh", warehouseId)
    _, err = qs.All(&items)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Database error",
            "code": 500,
        }
        c.ServeJSON()
        return
    }

    // Load relationships for each item
    for _, item := range items {
        o.LoadRelated(item, "Item")
        if item.Item != nil {
            o.LoadRelated(item.Item, "Category")
        }
        o.LoadRelated(item, "StatusLookup")
        o.LoadRelated(item, "Warehouse")
        o.LoadRelated(item, "CondLookup")
        o.LoadRelated(item, "User")
    }
    
    c.Data["json"] = items
    c.ServeJSON()
}

