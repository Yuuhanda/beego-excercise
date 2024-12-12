package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
    "time"
)

type LendingController struct {
    web.Controller
    lendingService *services.LendingService
}

func (c *LendingController) Prepare() {
    c.lendingService = services.NewLendingService()
}

// Create handles POST request to create new lending
func (c *LendingController) Create() {
    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        IdUnit     int    `json:"id_unit"`
        IdUser     int    `json:"user_id"`
        IdEmployee int    `json:"id_employee"`
        Type       int    `json:"type"`
        Date       string `json:"date"`
        PicLoan    string `json:"pic_loan"`
        PicReturn  string `json:"pic_return"`
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

    date, err := time.Parse(time.RFC3339, input.Date)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid date format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    lending := &models.Lending{
        IdUnit:     &models.ItemUnit{IdUnit: uint(input.IdUnit)},
        IdUser:     &models.User{Id: input.IdUser},
        IdEmployee: &models.Employee{IdEmployee: uint(input.IdEmployee)},
        Type:       &models.LendingTypeLookup{IdType: uint(input.Type)},
        Date:       date,
        PicLoan:    input.PicLoan,
        PicReturn:  input.PicReturn,
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

// GetActiveLoans handles GET request to fetch active loans
func (c *LendingController) GetActiveLoans() {
    lendings, err := c.lendingService.GetActiveLoans()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lendings
    c.ServeJSON()
}

// GetReturnedLoans handles GET request to fetch returned loans
func (c *LendingController) GetReturnedLoans() {
    lendings, err := c.lendingService.GetReturnedLoans()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lendings
    c.ServeJSON()
}
