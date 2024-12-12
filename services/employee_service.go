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
    o := orm.NewOrm()
    _, err := o.Insert(employee)
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
    o := orm.NewOrm()
    _, err := o.Update(employee)
    return err
}

// Delete deletes an employee
func (s *EmployeeService) Delete(employee *models.Employee) error {
    o := orm.NewOrm()
    _, err := o.Delete(employee)
    return err
}

// List retrieves employees with pagination
func (s *EmployeeService) List(page, pageSize int) ([]*models.Employee, int64, error) {
    var employees []*models.Employee
    offset := (page - 1) * pageSize
    
    o := orm.NewOrm()
    qs := o.QueryTable(new(models.Employee))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&employees)
    return employees, total, err
}
