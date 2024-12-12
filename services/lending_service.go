package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type LendingService struct {
    ormer orm.Ormer
}

func NewLendingService() *LendingService {
    return &LendingService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new lending record
func (s *LendingService) Create(lending *models.Lending) error {
    o := orm.NewOrm()
    _, err := o.Insert(lending)
    if err != nil {
        return err
    }
    
    // Load and populate all relationships
    if lending.IdUnit != nil {
        unit := &models.ItemUnit{IdUnit: lending.IdUnit.IdUnit}
        o.Read(unit)
        lending.IdUnit = unit
        
        o.LoadRelated(lending.IdUnit, "Item")
        o.LoadRelated(lending.IdUnit, "StatusLookup")
        o.LoadRelated(lending.IdUnit, "Warehouse")
        o.LoadRelated(lending.IdUnit, "CondLookup")
        o.LoadRelated(lending.IdUnit, "User")
        
        if lending.IdUnit.Item != nil {
            o.LoadRelated(lending.IdUnit.Item, "Category")
        }
    }
    
    if lending.IdUser != nil {
        user := &models.User{Id: lending.IdUser.Id}
        o.Read(user)
        lending.IdUser = user
    }
    
    if lending.IdEmployee != nil {
        employee := &models.Employee{IdEmployee: lending.IdEmployee.IdEmployee}
        o.Read(employee)
        lending.IdEmployee = employee
    }
    
    if lending.Type != nil {
        lendType := &models.LendingTypeLookup{IdType: lending.Type.IdType}
        o.Read(lendType)
        lending.Type = lendType
    }
    
    return nil
}



func (s *LendingService) GetByID(id uint) (*models.Lending, error) {
    o := orm.NewOrm()
    lending := &models.Lending{IdLending: id}
    
    if err := o.Read(lending); err != nil {
        return nil, err
    }
    
    o.LoadRelated(lending, "IdUnit")
    o.LoadRelated(lending, "IdUser")
    o.LoadRelated(lending, "IdEmployee")
    o.LoadRelated(lending, "Type")
    
    if lending.IdUnit != nil {
        o.LoadRelated(lending.IdUnit, "Item")
        o.LoadRelated(lending.IdUnit, "StatusLookup")
        o.LoadRelated(lending.IdUnit, "Warehouse")
        o.LoadRelated(lending.IdUnit, "CondLookup")
    }
    
    return lending, nil
}



// GetByUserID retrieves lendings by user ID
func (s *LendingService) GetByUserID(userID int) ([]*models.Lending, error) {
    var lendings []*models.Lending
    _, err := s.ormer.QueryTable(new(models.Lending)).Filter("UserId", userID).All(&lendings)
    return lendings, err
}

// Update updates lending information
func (s *LendingService) Update(lending *models.Lending) error {
    if lending.IdLending == 0 {
        return errors.New("lending ID is required")
    }
    _, err := s.ormer.Update(lending)
    return err
}

// Delete deletes a lending record
func (s *LendingService) Delete(id uint) error {
    lending := &models.Lending{IdLending: id}
    _, err := s.ormer.Delete(lending)
    return err
}

// List retrieves lendings with pagination
func (s *LendingService) List(page, pageSize int, filters map[string]string) ([]*models.Lending, int64, error) {
    var lendings []*models.Lending
    offset := (page - 1) * pageSize
    
    o := orm.NewOrm()
    qs := o.QueryTable(new(models.Lending)).RelatedSel()
    
    // Apply filters
    if empName := filters["employee_name"]; empName != "" {
        qs = qs.Filter("IdEmployee__EmpName__icontains", empName)
    }
    
    if itemName := filters["item_name"]; itemName != "" {
        qs = qs.Filter("IdUnit__Item__ItemName__icontains", itemName)
    }
    
    if serialNumber := filters["serial_number"]; serialNumber != "" {
        qs = qs.Filter("IdUnit__SerialNumber__icontains", serialNumber)
    }
    
    if username := filters["username"]; username != "" {
        qs = qs.Filter("IdUser__Username__icontains", username)
    }
    
    if typeId := filters["type"]; typeId != "" {
        qs = qs.Filter("Type__IdType", typeId)
    }
    
    if startDate := filters["start_date"]; startDate != "" {
        qs = qs.Filter("Date__gte", startDate)
    }
    
    if endDate := filters["end_date"]; endDate != "" {
        qs = qs.Filter("Date__lte", endDate)
    }
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&lendings)
    if err != nil {
        return nil, 0, err
    }

    // Load related data
    for _, lending := range lendings {
        if lending.IdUnit != nil {
            o.LoadRelated(lending.IdUnit, "Item")
            o.LoadRelated(lending.IdUnit, "StatusLookup")
            o.LoadRelated(lending.IdUnit, "Warehouse")
            o.LoadRelated(lending.IdUnit, "CondLookup")
            o.LoadRelated(lending.IdUnit, "User")
            
            if lending.IdUnit.Item != nil {
                o.LoadRelated(lending.IdUnit.Item, "Category")
            }
        }
        
        if lending.IdUser != nil {
            user := &models.User{Id: lending.IdUser.Id}
            o.Read(user)
            lending.IdUser = user
        }
        
        if lending.IdEmployee != nil {
            employee := &models.Employee{IdEmployee: lending.IdEmployee.IdEmployee}
            o.Read(employee)
            lending.IdEmployee = employee
        }
        
        if lending.Type != nil {
            lendType := &models.LendingTypeLookup{IdType: lending.Type.IdType}
            o.Read(lendType)
            lending.Type = lendType
        }
    }
    
    return lendings, total, nil
}




// GetActiveLoans retrieves all active loans (Type = 1)
func (s *LendingService) GetActiveLoans() ([]*models.Lending, error) {
    o := orm.NewOrm()
    var lendings []*models.Lending
    
    _, err := o.QueryTable(new(models.Lending)).Filter("Type__IdType", 1).All(&lendings)
    if err != nil {
        return nil, err
    }

    for _, lending := range lendings {
        // Load ItemUnit and all its relations
        if lending.IdUnit != nil {
            unit := &models.ItemUnit{IdUnit: lending.IdUnit.IdUnit}
            o.Read(unit)
            o.LoadRelated(unit, "Item")
            o.LoadRelated(unit, "StatusLookup")
            o.LoadRelated(unit, "Warehouse")
            o.LoadRelated(unit, "CondLookup")
            o.LoadRelated(unit, "User")
            
            if unit.Item != nil {
                o.LoadRelated(unit.Item, "Category")
            }
            lending.IdUnit = unit
        }
        
        // Load User data
        if lending.IdUser != nil {
            user := &models.User{Id: lending.IdUser.Id}
            o.Read(user)
            lending.IdUser = user
        }
        
        // Load Employee data
        if lending.IdEmployee != nil {
            employee := &models.Employee{IdEmployee: lending.IdEmployee.IdEmployee}
            o.Read(employee)
            lending.IdEmployee = employee
        }
        
        // Load Type data
        if lending.Type != nil {
            lendType := &models.LendingTypeLookup{IdType: lending.Type.IdType}
            o.Read(lendType)
            lending.Type = lendType
        }
    }
    
    return lendings, nil
}


// GetReturnedLoans retrieves all returned loans (Type = 2)
func (s *LendingService) GetReturnedLoans() ([]*models.Lending, error) {
    o := orm.NewOrm()
    var lendings []*models.Lending
    
    _, err := o.QueryTable(new(models.Lending)).Filter("Type__IdType", 2).All(&lendings)
    if err != nil {
        return nil, err
    }

    for _, lending := range lendings {
        // Load ItemUnit and all its relations
        if lending.IdUnit != nil {
            unit := &models.ItemUnit{IdUnit: lending.IdUnit.IdUnit}
            o.Read(unit)
            o.LoadRelated(unit, "Item")
            o.LoadRelated(unit, "StatusLookup")
            o.LoadRelated(unit, "Warehouse")
            o.LoadRelated(unit, "CondLookup")
            o.LoadRelated(unit, "User")
            
            if unit.Item != nil {
                o.LoadRelated(unit.Item, "Category")
            }
            lending.IdUnit = unit
        }
        
        // Load User data
        if lending.IdUser != nil {
            user := &models.User{Id: lending.IdUser.Id}
            o.Read(user)
            lending.IdUser = user
        }
        
        // Load Employee data
        if lending.IdEmployee != nil {
            employee := &models.Employee{IdEmployee: lending.IdEmployee.IdEmployee}
            o.Read(employee)
            lending.IdEmployee = employee
        }
        
        // Load Type data
        if lending.Type != nil {
            lendType := &models.LendingTypeLookup{IdType: lending.Type.IdType}
            o.Read(lendType)
            lending.Type = lendType
        }
    }
    
    return lendings, nil
}
