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
    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        EmpName string `json:"emp_name"`
        Phone   string `json:"phone"`
        Email   string `json:"email"`
        Address string `json:"address"`
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

    employee := &models.Employee{
        EmpName: input.EmpName,
        Phone:   input.Phone,
        Email:   input.Email,
        Address: input.Address,
    }

    service := services.NewEmployeeService()
    if err := service.Create(employee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create employee",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Employee created successfully",
            "data":    employee,
        }
    }
    c.ServeJSON()
}

func (c *EmployeeController) Update() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    body := c.Ctx.Input.CopyBody(1048576)
    
    var input struct {
        EmpName string `json:"emp_name,omitempty"`
        Phone   string `json:"phone,omitempty"`
        Email   string `json:"email,omitempty"`
        Address string `json:"address,omitempty"`
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

    service := services.NewEmployeeService()
    existingEmployee, err := service.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Employee not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if input.EmpName != "" {
        existingEmployee.EmpName = input.EmpName
    }
    if input.Phone != "" {
        existingEmployee.Phone = input.Phone
    }
    if input.Email != "" {
        existingEmployee.Email = input.Email
    }
    if input.Address != "" {
        existingEmployee.Address = input.Address
    }

    if err := service.Update(existingEmployee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update employee",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Employee updated successfully",
            "data":    existingEmployee,
        }
    }
    c.ServeJSON()
}

func (c *EmployeeController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    service := services.NewEmployeeService()
    existingEmployee, err := service.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Employee not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    if err := service.Delete(existingEmployee); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete employee",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Employee deleted successfully",
        }
    }
    c.ServeJSON()
}

func (c *EmployeeController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)

    service := services.NewEmployeeService()
    employees, total, err := service.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to fetch employees",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "items": employees,
            "total": total,
            "page":  page,
        },
    }
    c.ServeJSON()
}

func (c *EmployeeController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    service := services.NewEmployeeService()
    employee, err := service.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Employee not found",
            "error":   err.Error(),
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
