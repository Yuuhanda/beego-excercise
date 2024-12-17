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

        if lending.IdUnit.User != nil {
            lending.IdUnit.User = &models.User{
                Id:       lending.IdUser.Id,
                Username: lending.IdUser.Username,
                Email:    lending.IdUser.Email,
            }
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
        o.LoadRelated(lending.IdUnit, "User")
        if lending.IdUnit.Item != nil {
            o.LoadRelated(lending.IdUnit.Item, "Category")
        }
    }

    // Create simplified User data
    if lending.IdUser != nil {
        lending.IdUser = &models.User{
            Id:       lending.IdUser.Id,
            Username: lending.IdUser.Username,
            Email:    lending.IdUser.Email,
        }
    }
    if lending.IdUnit != nil {
        lending.IdUnit.User = &models.User{
            Id:       lending.IdUnit.User.Id,
            Username: lending.IdUnit.User.Username,
            Email:    lending.IdUnit.User.Email,
        }
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

            // Create simplified User data
            if lending.IdUnit.User != nil {
                lending.IdUnit.User = &models.User{
                    Id:       lending.IdUnit.User.Id,
                    Username: lending.IdUnit.User.Username,
                    Email:    lending.IdUnit.User.Email,
                }
            }


        }
        
        // Simplify IdUser data
        if lending.IdUser != nil {
            userData := lending.IdUser
            lending.IdUser = &models.User{
                Id:       userData.Id,
                Username: userData.Username,
                Email:    userData.Email,
            }
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
func (s *LendingService) GetActiveLoans(filters map[string]string) ([]*models.Lending, error) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(models.Lending)).Filter("Type__IdType", 1)
    
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
    if startDate := filters["start_date"]; startDate != "" {
        qs = qs.Filter("Date__gte", startDate)
    }
    if endDate := filters["end_date"]; endDate != "" {
        qs = qs.Filter("Date__lte", endDate)
    }

    var lendings []*models.Lending
    _, err := qs.All(&lendings)
    if err != nil {
        return nil, err
    }

    // Load all related data
    for _, lending := range lendings {
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

            // Create simplified User data
            if unit.User != nil {
                unit.User = &models.User{
                    Id:       unit.User.Id,
                    Username: unit.User.Username,
                    Email:    unit.User.Email,
                }
            }

            lending.IdUnit = unit
        }
        o.LoadRelated(lending, "IdUser")
        // Simplify IdUser data
        if lending.IdUser != nil {
            userData := lending.IdUser
            lending.IdUser = &models.User{
                Id:       userData.Id,
                Username: userData.Username,
                Email:    userData.Email,
            }
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
    
    return lendings, nil
}

func (s *LendingService) GetReturnedLoans(filters map[string]string) ([]*models.Lending, error) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(models.Lending)).Filter("Type__IdType", 2)
    
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
    if startDate := filters["start_date"]; startDate != "" {
        qs = qs.Filter("Date__gte", startDate)
    }
    if endDate := filters["end_date"]; endDate != "" {
        qs = qs.Filter("Date__lte", endDate)
    }

    var lendings []*models.Lending
    _, err := qs.All(&lendings)
    if err != nil {
        return nil, err
    }

    // Load all related data
    for _, lending := range lendings {
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

            // Create simplified User data
            if unit.User != nil {
                unit.User = &models.User{
                    Id:       unit.User.Id,
                    Username: unit.User.Username,
                    Email:    unit.User.Email,
                }
            }

            lending.IdUnit = unit
        }
        o.LoadRelated(lending, "IdUser")
        // Simplify IdUser data
        if lending.IdUser != nil {
            userData := lending.IdUser
            lending.IdUser = &models.User{
                Id:       userData.Id,
                Username: userData.Username,
                Email:    userData.Email,
            }
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
    
    return lendings, nil
}


func (s *LendingService) SearchItemReport(page, pageSize int, filters map[string]string) ([]map[string]interface{}, int64, error) {
    o := orm.NewOrm()
    qb, _ := orm.NewQueryBuilder("mysql")

    // Base query for item report
    qb.Select("i.SKU",
        "i.item_name",
        "i.id_item",
        "COUNT(CASE WHEN l.type = 1 THEN l.id_lending END) as total_item_lent").
        From("item_unit iu").
        LeftJoin("lending l ON iu.id_unit = l.id_unit").
        LeftJoin("item i ON i.id_item = iu.id_item")

    // Add filters
    if itemName := filters["item_name"]; itemName != "" {
        qb.Where("i.item_name LIKE ?")
    }
    if sku := filters["SKU"]; sku != "" {
        qb.And("i.SKU LIKE ?")
    }
    if idItem := filters["id_item"]; idItem != "" {
        qb.And("i.id_item = ?")
    }

    qb.GroupBy("i.id_item").
        OrderBy("total_item_lent DESC").
        Limit(pageSize).
        Offset((page - 1) * pageSize)

    sql := qb.String()
    var maps []orm.Params
    var args []interface{}

    // Add filter values to args
    if itemName := filters["item_name"]; itemName != "" {
        args = append(args, "%"+itemName+"%")
    }
    if sku := filters["SKU"]; sku != "" {
        args = append(args, "%"+sku+"%")
    }
    if idItem := filters["id_item"]; idItem != "" {
        args = append(args, idItem)
    }

    num, err := o.Raw(sql, args...).Values(&maps)
    if err != nil {
        return nil, 0, err
    }

    result := make([]map[string]interface{}, len(maps))
    for i, m := range maps {
        result[i] = make(map[string]interface{})
        for k, v := range m {
            result[i][k] = v
        }
    }

    return result, num, nil
}

func (s *LendingService) SearchUnitReport(page, pageSize int, filters map[string]string) ([]map[string]interface{}, int64, error) {
    o := orm.NewOrm()
    qb, _ := orm.NewQueryBuilder("mysql")

    // Base query for unit report
    qb.Select("iu.serial_number",
        "i.item_name",
        "i.id_item",
        "l.id_unit",
        "COUNT(l.id_unit) as number_of_times_unit_is_lent").
        From("item_unit iu").
        LeftJoin("lending l ON iu.id_unit = l.id_unit").
        LeftJoin("item i ON i.id_item = iu.id_item")

    // Add filters
    if itemName := filters["item_name"]; itemName != "" {
        qb.Where("i.item_name LIKE ?")
    }
    if serialNumber := filters["serial_number"]; serialNumber != "" {
        qb.And("iu.serial_number LIKE ?")
    }

    qb.GroupBy("l.id_unit").
        Having("COUNT(l.id_unit) > 0").
        OrderBy("number_of_times_unit_is_lent DESC").
        Limit(pageSize).
        Offset((page - 1) * pageSize)

    sql := qb.String()
    var maps []orm.Params
    var args []interface{}

    // Add filter values to args
    if itemName := filters["item_name"]; itemName != "" {
        args = append(args, "%"+itemName+"%")
    }
    if serialNumber := filters["serial_number"]; serialNumber != "" {
        args = append(args, "%"+serialNumber+"%")
    }

    num, err := o.Raw(sql, args...).Values(&maps)
    if err != nil {
        return nil, 0, err
    }

    result := make([]map[string]interface{}, len(maps))
    for i, m := range maps {
        result[i] = make(map[string]interface{})
        for k, v := range m {
            result[i][k] = v
        }
    }

    return result, num, nil
}
