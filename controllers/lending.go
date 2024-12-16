package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "time"
    "github.com/beego/beego/v2/core/logs"
)

type LendingController struct {
    web.Controller
    lendingService *services.LendingService
    itemUnitService *services.ItemUnitService
    userService      *services.UserService
    employeeService *services.EmployeeService
    warehouseService *services.WarehouseService
}

func (c *LendingController) Prepare() {
    c.lendingService = services.NewLendingService()
    c.userService = services.NewUserService()
    c.itemUnitService = services.NewItemUnitService()
    c.employeeService = services.NewEmployeeService()
    c.warehouseService = services.NewWarehouseService()
}

// Create handles POST request to create new lending
func (c *LendingController) Create() {
    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        IdUnit     int    `json:"id_unit"`
        IdUser     int    `json:"user_id"`
        IdEmployee int    `json:"id_employee"`
        PicLoan    string `json:"pic_loan"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    lending := &models.Lending{
        IdUnit:     &models.ItemUnit{IdUnit: uint(input.IdUnit)},
        IdUser:     &models.User{Id: input.IdUser},
        IdEmployee: &models.Employee{IdEmployee: uint(input.IdEmployee)},
        Type:       &models.LendingTypeLookup{IdType: 2},
        Date:       time.Now(),
        PicLoan:    input.PicLoan,
    }

    if err := c.lendingService.Create(lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create lending",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Lending created successfully",
            "data":    lending,
        }
    }

    //get username that created the item unit
    user, err := c.userService.GetByID(input.IdUser)
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    //get serial number
    itemUnit, err := c.itemUnitService.Get(input.IdUnit)
    if err != nil {
        logs.Error("Failed to get item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get item unit",
        }
        c.ServeJSON()
        return
    }

    //get employee name
    employee, err := c.employeeService.GetByID(uint(input.IdEmployee))
    if err != nil {
        logs.Error("Failed to get employee: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get employee",
        }
        c.ServeJSON()
    }

    // Create UnitLog
    unitLogService := services.NewUnitLogService()

    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
        Content:      "Unit lent out",
        ActorsAction: "Unit "+ itemUnit.SerialNumber +" lent out "+ employee.EmpName+ " by " + user.Username,
        UpdateAt:     time.Now(),
    }


    err = unitLogService.Create(unitLog)
    if err != nil {
        // Handle error case
        logs.Error("Failed to create unit log: %v", err)
    }

    itemUnitService := services.NewItemUnitService()
    // Get existing item unit data first
    existingItemUnit, err := itemUnitService.Get(input.IdUnit)
    if err != nil {
        logs.Error("Failed to get existing item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get existing item unit",
        }
        c.ServeJSON()
        return
    }
    
    itemUnit = &models.ItemUnit{
        IdUnit:       uint(input.IdUnit),
        Comment:      existingItemUnit.Comment,
        StatusLookup: &models.StatusLookup{IdStatus: uint(2)},
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



// Get handles GET request to fetch lending by ID
func (c *LendingController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    lending, err := c.lendingService.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lending
    c.ServeJSON()
}

// List handles GET request to fetch paginated lendings
func (c *LendingController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)
    
    filters := map[string]string{
        "employee_name": c.GetString("employee_name"),
        "item_name":    c.GetString("item_name"),
        "serial_number": c.GetString("serial_number"),
        "username":     c.GetString("username"),
        "type":         c.GetString("type"),
        "start_date":   c.GetString("start_date"),
        "end_date":     c.GetString("end_date"),
    }

    lendings, total, err := c.lendingService.List(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "items": lendings,
            "total": total,
            "page":  page,
        },
    }
    c.ServeJSON()
}

// Update handles PUT request to update lending
func (c *LendingController) Update() {
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

    body := c.Ctx.Input.CopyBody(1048576) // Copy request body with sufficient buffer
    
    var input struct {
        IdUnit     int    `json:"id_unit"`
        UserId     int    `json:"user_id"`
        IdEmployee int    `json:"id_employee"`
        Type       int    `json:"type"`
        Date       string `json:"date"`
        PicLoan    string `json:"pic_loan"`
        PicReturn  string `json:"pic_return"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    date, err := time.Parse(time.RFC3339, input.Date)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid date format",
        }
        c.ServeJSON()
        return
    }

    lending := &models.Lending{
        IdLending:  uint(id),
        IdUnit:     &models.ItemUnit{IdUnit: uint(input.IdUnit)},
        IdUser:     &models.User{Id: input.UserId},
        IdEmployee: &models.Employee{IdEmployee: uint(input.IdEmployee)},
        Type:       &models.LendingTypeLookup{IdType: uint(input.Type)},
        Date:       date,
        PicLoan:    input.PicLoan,
        PicReturn:  input.PicReturn,
    }

    if err := c.lendingService.Update(lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update lending",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Lending updated successfully",
            "data":    lending,
        }
    }
    c.ServeJSON()
}


// Delete handles DELETE request to remove lending
func (c *LendingController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    if err := c.lendingService.Delete(uint(id)); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "message": "Lending deleted successfully",
    }
    c.ServeJSON()
}

func (c *LendingController) GetActiveLoans() {
    filters := map[string]string{
        "employee_name": c.GetString("employee_name"),
        "item_name":    c.GetString("item_name"),
        "serial_number": c.GetString("serial_number"),
        "username":     c.GetString("username"),
        "start_date":   c.GetString("start_date"),
        "end_date":     c.GetString("end_date"),
    }

    lendings, err := c.lendingService.GetActiveLoans(filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    lendings,
    }
    c.ServeJSON()
}

func (c *LendingController) GetReturnedLoans() {
    filters := map[string]string{
        "employee_name": c.GetString("employee_name"),
        "item_name":    c.GetString("item_name"),
        "serial_number": c.GetString("serial_number"),
        "username":     c.GetString("username"),
        "start_date":   c.GetString("start_date"),
        "end_date":     c.GetString("end_date"),
    }

    lendings, err := c.lendingService.GetReturnedLoans(filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    lendings,
    }
    c.ServeJSON()
}

func (c *LendingController) SearchItemReport() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "20"))

    filters := map[string]string{
        "item_name": c.GetString("item_name"),
        "SKU":       c.GetString("SKU"),
        "id_item":   c.GetString("id_item"),
    }

    items, total, err := c.lendingService.SearchItemReport(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get item report",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "items": items,
            "total": total,
            "page":  page,
            "size":  pageSize,
        },
    }
    c.ServeJSON()
}

func (c *LendingController) SearchUnitReport() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "20"))

    filters := map[string]string{
        "item_name":     c.GetString("item_name"),
        "serial_number": c.GetString("serial_number"),
    }

    units, total, err := c.lendingService.SearchUnitReport(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get unit report",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "units": units,
            "total": total,
            "page":  page,
            "size":  pageSize,
        },
    }
    c.ServeJSON()
}


func (c *LendingController) Return() {
    lendingIdStr := c.Ctx.Input.Param(":id")
    lendingId, err := strconv.ParseUint(lendingIdStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid lending ID format",
        }
        c.ServeJSON()
        return
    }

    // First get existing lending data
    existingLending, err := c.lendingService.GetByID(uint(lendingId))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Lending not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        UserId     int    `json:"user_id"`
        IdEmployee int    `json:"id_employee"`
        PicReturn  string `json:"pic_return"`
        Condition int `json:"condition"`
        Comment string `json:"comment"`
        IdWh int    `json:"IdWh"`
    }

    if err := json.Unmarshal(body, &input); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
            "received": string(body),
        }
        c.ServeJSON()
        return
    }

    lending := &models.Lending{
        IdLending:  uint(lendingId),
        IdUnit:     existingLending.IdUnit, // Preserve existing IdUnit
        IdUser:     &models.User{Id: input.UserId},
        IdEmployee: &models.Employee{IdEmployee: uint(input.IdEmployee)},
        Type:       &models.LendingTypeLookup{IdType: 2},
        Date:       time.Now(),
        PicLoan: existingLending.PicLoan,
        PicReturn:  input.PicReturn,
    }

    if err := c.lendingService.Update(lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to process return",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Return processed successfully",
            "data":    lending,
        }
    }

    //get username that created the item unit
    user, err := c.userService.GetByID(input.UserId)
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    //get serial number
    itemUnit, err := c.itemUnitService.Get(int(existingLending.IdUnit.IdUnit))
    if err != nil {
        logs.Error("Failed to get item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get item unit",
        }
        c.ServeJSON()
        return
    }

    //get employee name
    employee, err := c.employeeService.GetByID(uint(input.IdEmployee))
    if err != nil {
        logs.Error("Failed to get employee: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get employee",
        }
        c.ServeJSON()
    }

    //get warehouse name
    warehouse, err := c.warehouseService.GetByID(uint(input.IdWh))
    if err != nil {
        logs.Error("Failed to get warehouse: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get warehouse",
        }
        c.ServeJSON()
        return
    }

    // Create UnitLog
    unitLogService := services.NewUnitLogService()

    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
        Content:      input.Comment,
        ActorsAction: "Unit "+ itemUnit.SerialNumber +" lent out to "+ employee.EmpName+" returned to"+ warehouse.WhName +" by " + user.Username,
        UpdateAt:     time.Now(),
    }


    err = unitLogService.Create(unitLog)
    if err != nil {
        // Handle error case
        logs.Error("Failed to create unit log: %v", err)
    }

    itemUnitService := services.NewItemUnitService()
    itemUnit = &models.ItemUnit{
        IdUnit:      uint(existingLending.IdUnit.IdUnit),
        Comment:     input.Comment,
        StatusLookup: &models.StatusLookup{IdStatus: uint(1)},
        Warehouse:   &models.Warehouse{IdWh: uint(input.IdWh)},
        CondLookup:  &models.ConditionLookup{IdCondition: uint(input.Condition)},
        User:        &models.User{Id: input.UserId},
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


