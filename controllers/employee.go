package controllers

import (
    "encoding/json"
    "myproject/models"
    "myproject/services"
    "strconv"

    beego "github.com/beego/beego/v2/server/web"
)

type EmployeeController struct {
    beego.Controller
    employeeService *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
    return &EmployeeController{
        employeeService: services.NewEmployeeService(),
    }
}

// Create handles POST request to create a new employee
func (c *EmployeeController) Create() {
    var employee models.Employee
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &employee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.employeeService.Create(&employee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    employee,
    }
    c.ServeJSON()
}

// Get handles GET request to retrieve an employee by ID
func (c *EmployeeController) Get() {
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

    employee, err := c.employeeService.GetByID(uint(id))
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
        "data":    employee,
    }
    c.ServeJSON()
}

// List handles GET request to retrieve all employees with pagination
func (c *EmployeeController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)

    employees, total, err := c.employeeService.List(page, pageSize)
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
            "items": employees,
            "total": total,
        },
    }
    c.ServeJSON()
}

// Update handles PUT request to update an employee
func (c *EmployeeController) Update() {
    var employee models.Employee
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &employee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.employeeService.Update(&employee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    employee,
    }
    c.ServeJSON()
}

// Delete handles DELETE request to remove an employee
func (c *EmployeeController) Delete() {
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

    if err := c.employeeService.Delete(uint(id)); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Employee deleted successfully",
    }
    c.ServeJSON()
}
