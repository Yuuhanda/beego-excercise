package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type EmployeeService struct {
    ormer orm.Ormer
}

func NewEmployeeService() *EmployeeService {
    return &EmployeeService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new employee
func (s *EmployeeService) Create(employee *models.Employee) error {
    _, err := s.ormer.Insert(employee)
    return err
}

// GetByID retrieves employee by ID
func (s *EmployeeService) GetByID(id uint) (*models.Employee, error) {
    employee := &models.Employee{IdEmployee: id}
    err := s.ormer.Read(employee)
    if err == orm.ErrNoRows {
        return nil, errors.New("employee not found")
    }
    return employee, err
}

// GetByEmail retrieves employee by email
func (s *EmployeeService) GetByEmail(email string) (*models.Employee, error) {
    employee := &models.Employee{Email: email}
    err := s.ormer.Read(employee, "Email")
    if err == orm.ErrNoRows {
        return nil, errors.New("employee not found")
    }
    return employee, err
}

// Update updates employee information
func (s *EmployeeService) Update(employee *models.Employee) error {
    if employee.IdEmployee == 0 {
        return errors.New("employee ID is required")
    }
    _, err := s.ormer.Update(employee)
    return err
}

// Delete deletes an employee
func (s *EmployeeService) Delete(id uint) error {
    employee := &models.Employee{IdEmployee: id}
    _, err := s.ormer.Delete(employee)
    return err
}

// List retrieves employees with pagination
func (s *EmployeeService) List(page, pageSize int) ([]*models.Employee, int64, error) {
    var employees []*models.Employee
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.Employee))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&employees)
    return employees, total, err
}
