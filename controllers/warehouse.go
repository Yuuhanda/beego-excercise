package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "fmt"
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
)

type WarehouseController struct {
    web.Controller
    warehouseService *services.WarehouseService
}

func (c *WarehouseController) Prepare() {
    c.warehouseService = services.NewWarehouseService()
}

// CreateWarehouse handles warehouse creation
// @router /warehouse [post]
// CreateWarehouse handles warehouse creation
// @router /warehouse [post]
func (c *WarehouseController) CreateWarehouse() {
    body := c.Ctx.Input.CopyBody(1048576) // Read up to 1MB
    logs.Info("Raw request body: %s", string(body))
    
    var input struct{
        WhName    string `json:"WhName"`
        WhAddress string `json:"WhAddress"`
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


    Warehouse := &models.Warehouse{
        WhName:    input.WhName,
        WhAddress: input.WhAddress,
    }

    if err := c.warehouseService.Create(Warehouse); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = Warehouse
    c.ServeJSON()
}


// GetWarehouse retrieves warehouse by ID
// @router /warehouse/:id [get]
func (c *WarehouseController) GetWarehouse() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid warehouse ID",
        }
        c.ServeJSON()
        return
    }

    warehouse, err := c.warehouseService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Warehouse not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    warehouse,
        }
    }
    c.ServeJSON()
}

// UpdateWarehouse updates warehouse information
// @router /warehouse/:id [put]
func (c *WarehouseController) UpdateWarehouse() {
    // Get ID from URL parameter
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid warehouse ID",
        }
        c.ServeJSON()
        return
    }

    var input struct {
        WhName    string `json:"WhName"`
        WhAddress string `json:"WhAddress"`
    }

    body := c.Ctx.Input.CopyBody(1048576)
    logs.Info("Raw request body: %s", string(body))

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

    warehouse := &models.Warehouse{
        IdWh:      id,  // Set the ID from URL parameter
        WhName:    input.WhName,
        WhAddress: input.WhAddress,
    }

    if err := c.warehouseService.Update(warehouse); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = warehouse
    c.ServeJSON()
}


// DeleteWarehouse deletes a warehouse
// @router /warehouse/:id [delete]
func (c *WarehouseController) DeleteWarehouse() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid warehouse ID",
        }
        c.ServeJSON()
        return
    }

    if err := c.warehouseService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete warehouse",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Warehouse deleted successfully",
        }
    }
    c.ServeJSON()
}

// ListWarehouses retrieves paginated and filtered warehouses
// @router /warehouses [get]
func (c *WarehouseController) ListWarehouses() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("page_size", 10)
    whName := c.GetString("wh_name")
    whAddress := c.GetString("wh_address")

    warehouses, total, err := c.warehouseService.List(page, pageSize, whName, whAddress)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve warehouses",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    warehouses,
            "total":   total,
            "page":    page,
            "size":    pageSize,
            "filters": map[string]string{
                "wh_name":    whName,
                "wh_address": whAddress,
            },
        }
    }
    c.ServeJSON()
}

// GetWarehouseUsers retrieves all users associated with a warehouse
// @router /warehouse/:id/users [get]
func (c *WarehouseController) GetWarehouseUsers() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid warehouse ID",
        }
        c.ServeJSON()
        return
    }

    users, err := c.warehouseService.GetWarehouseUsers(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve warehouse users",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    users,
        }
    }
    c.ServeJSON()
}
