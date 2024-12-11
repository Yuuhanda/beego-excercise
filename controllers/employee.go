package controllers

import (
    "myproject/services"
)

type EmployeeController struct {
    employeeService *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
    return &EmployeeController{
        employeeService: services.NewEmployeeService(),
    }
}
