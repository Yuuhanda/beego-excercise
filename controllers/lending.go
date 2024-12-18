package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "time"
    "github.com/beego/beego/v2/core/logs"
    "fmt"
    "path/filepath"
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

func (c *LendingController) Create() {
    // Handle file upload
    file, header, err := c.GetFile("pic_loan")
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get uploaded file",
        }
        c.ServeJSON()
        return
    }
    defer file.Close()

    filename, err := c.lendingService.UploadLoanPicture(file, header)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to upload loan picture",
        }
        c.ServeJSON()
        return
    }

    // Get form values
    input := struct {
        IdUnit     uint
        IdUser     int
        IdEmployee uint
        Type       uint
    }{}

    unitId, _ := c.GetInt("id_unit")
    userId, _ := c.GetInt("user_id")
    employeeId, _ := c.GetInt("id_employee")
    typeId, _ := c.GetInt("type")

    input.IdUnit = uint(unitId)
    input.IdUser = userId
    input.IdEmployee = uint(employeeId)
    input.Type = uint(typeId)



    lending := &models.Lending{
        IdUnit:     &models.ItemUnit{IdUnit: input.IdUnit},
        IdUser:     &models.User{Id: input.IdUser},
        IdEmployee: &models.Employee{IdEmployee: input.IdEmployee},
        Type:       &models.LendingTypeLookup{IdType: input.Type},
        Date:       time.Now(),
        PicLoan:    filename,
    }

    if err := c.lendingService.Create(lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create lending record",
        }
        c.ServeJSON()
        return
    }

    // Get username
    user, err := c.userService.GetByID(input.IdUser)
    if err != nil {
        logs.Error("Failed to get user: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get user",
        }
        c.ServeJSON()
        return
    }

    // Get item unit
    itemUnit, err := c.itemUnitService.Get(int(input.IdUnit))
    if err != nil {
        logs.Error("Failed to get item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get item unit",
        }
        c.ServeJSON()
        return
    }

    // Get employee
    employee, err := c.employeeService.GetByID(input.IdEmployee)
    if err != nil {
        logs.Error("Failed to get employee: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to get employee",
        }
        c.ServeJSON()
        return
    }

    // Create UnitLog
    unitLogService := services.NewUnitLogService()
    unitLog := &models.UnitLog{
        IdUnit:       &models.ItemUnit{IdUnit: itemUnit.IdUnit},
        Content:      "Unit lent out",
        ActorsAction: "Unit " + itemUnit.SerialNumber + " lent out " + employee.EmpName + " by " + user.Username,
        UpdateAt:     time.Now(),
    }

    if err = unitLogService.Create(unitLog); err != nil {
        logs.Error("Failed to create unit log: %v", err)
    }

    // Update item unit status
    itemUnitService := services.NewItemUnitService()
    updatedItemUnit := &models.ItemUnit{
        IdUnit:       input.IdUnit,
        Comment:      itemUnit.Comment,
        StatusLookup: &models.StatusLookup{IdStatus: uint(2)},
        Warehouse:    &models.Warehouse{IdWh: itemUnit.Warehouse.IdWh},
        CondLookup:   &models.ConditionLookup{IdCondition: itemUnit.CondLookup.IdCondition},
        User:         &models.User{Id: itemUnit.User.Id},
    }

    if err = itemUnitService.Update(updatedItemUnit); err != nil {
        logs.Error("Failed to update item unit: %v", err)
        c.Data["json"] = map[string]interface{}{
            "error": "Failed to update item unit",
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Lending created successfully",
        "data":    lending,
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
    // Handle file upload first
    file, header, err := c.GetFile("pic_return")
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get uploaded file",
        }
        c.ServeJSON()
        return
    }
    defer file.Close()

    filename, err := c.lendingService.UploadReturnPicture(file, header)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to upload return picture",
        }
        c.ServeJSON()
        return
    }

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

    // Get existing lending data
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

    var IdEmployee uint
    // Get form values
    input := struct {
        UserId     int
        Condition  int
        Comment    string
        IdWh       int
    }{}

    userId, _ := c.GetInt("user_id")
    condition, _ := c.GetInt("condition")
    warehouseId, _ := c.GetInt("IdWh")

    input.UserId = userId
    
    input.Condition = condition
    input.Comment = c.GetString("comment")
    input.IdWh = warehouseId
    lending := &models.Lending{
        IdLending:  uint(lendingId),
        IdUnit:     existingLending.IdUnit,
        IdUser:     &models.User{Id: userId},
        IdEmployee: existingLending.IdEmployee,
        Type:       &models.LendingTypeLookup{IdType: 2},
        Date:       time.Now(),
        PicLoan:    existingLending.PicLoan,
        PicReturn:  filename,
    }
    IdEmployee = existingLending.IdEmployee.IdEmployee

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
    employee, err := c.employeeService.GetByID(uint(IdEmployee))
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
        ActorsAction: "Unit "+ itemUnit.SerialNumber +" lent out to "+ employee.EmpName+" returned to "+ warehouse.WhName +" by " + user.Username,
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


// GetLoanImage serves the lending's loan picture
// @router /lending/:id/loan-image [get]
// how to use it
// const loanImageUrl = `http://your-api-domain/api/lending/${lendingId}/loan-image`;
// const returnImageUrl = `http://your-api-domain/api/lending/${lendingId}/return-image`;

func (c *LendingController) GetLoanImage() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Ctx.Output.Status = 400
        c.Ctx.Output.Body([]byte("Invalid lending ID"))
        return
    }

    lending, err := c.lendingService.GetByID(id)
    if err != nil {
        c.Ctx.Output.Status = 404
        c.Ctx.Output.Body([]byte("Lending not found"))
        return
    }

    imagePath := filepath.Join("static/uploads/lending", lending.PicLoan)
    c.Ctx.Output.Download(imagePath, lending.PicLoan)
}

// GetReturnImage serves the lending's return picture
// @router /lending/:id/return-image [get]
func (c *LendingController) GetReturnImage() {
    idStr := c.Ctx.Input.Param(":id")
    var id uint
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Ctx.Output.Status = 400
        c.Ctx.Output.Body([]byte("Invalid lending ID"))
        return
    }

    lending, err := c.lendingService.GetByID(id)
    if err != nil {
        c.Ctx.Output.Status = 404
        c.Ctx.Output.Body([]byte("Lending not found"))
        return
    }

    imagePath := filepath.Join("static/uploads/lending/return", lending.PicReturn)
    c.Ctx.Output.Download(imagePath, lending.PicReturn)
}
