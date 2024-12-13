package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "myproject/database"
    "github.com/beego/beego/v2/core/logs"
    "fmt"
    "strings"
)

type ItemUnitService struct {
    ormer orm.Ormer
}


func NewItemUnitService() *ItemUnitService {
    return &ItemUnitService{
        ormer: database.GetOrmer(),
    }
}
// Create creates a new item unit
func (s *ItemUnitService) Create(itemUnit *models.ItemUnit) error {
    o := orm.NewOrm()
    _, err := o.Insert(itemUnit)
    return err
}





func (s *ItemUnitService) Get(id int) (*models.ItemUnit, error) {
    if id <= 0 {
        return nil, errors.New("invalid item unit ID")
    }

    itemUnit := &models.ItemUnit{IdUnit: uint(id)}
    err := s.ormer.Read(itemUnit)
    
    if err == orm.ErrNoRows {
        return nil, errors.New("item unit not found")
    }
    
    if err != nil {
        return nil, fmt.Errorf("database error: %v", err)
    }
    
    s.ormer.LoadRelated(itemUnit, "Item")
    s.ormer.LoadRelated(itemUnit, "StatusLookup")
    s.ormer.LoadRelated(itemUnit, "Warehouse")
    s.ormer.LoadRelated(itemUnit, "CondLookup")
    s.ormer.LoadRelated(itemUnit, "User")
    
    return itemUnit, nil
}


func (s *ItemUnitService) GetBySerialNumber(serialNumber string) (*models.ItemUnit, error) {
    itemUnit := &models.ItemUnit{SerialNumber: serialNumber}
    err := s.ormer.Read(itemUnit, "SerialNumber")
    if err == orm.ErrNoRows {
        return nil, errors.New("item unit not found")
    }
    
    s.ormer.LoadRelated(itemUnit, "Item")
    s.ormer.LoadRelated(itemUnit, "StatusLookup")
    s.ormer.LoadRelated(itemUnit, "Warehouse")
    s.ormer.LoadRelated(itemUnit, "CondLookup")
    s.ormer.LoadRelated(itemUnit, "User")
    
    return itemUnit, err
}

func (s *ItemUnitService) Update(itemUnit *models.ItemUnit) error {
    o := orm.NewOrm()
    
    // Get existing item unit
    existing := &models.ItemUnit{IdUnit: itemUnit.IdUnit}
    if err := o.Read(existing); err != nil {
        return err
    }

    // Update only non-empty fields
    if itemUnit.Comment != "" {
        existing.Comment = itemUnit.Comment
    }
    if itemUnit.Item != nil && itemUnit.Item.IdItem > 0 {
        existing.Item = itemUnit.Item
    }
    if itemUnit.StatusLookup != nil && itemUnit.StatusLookup.IdStatus > 0 {
        existing.StatusLookup = itemUnit.StatusLookup
    }
    if itemUnit.Warehouse != nil && itemUnit.Warehouse.IdWh > 0 {
        existing.Warehouse = itemUnit.Warehouse
    }
    if itemUnit.CondLookup != nil && itemUnit.CondLookup.IdCondition > 0 {
        existing.CondLookup = itemUnit.CondLookup
    }
    if itemUnit.User != nil && itemUnit.User.Id > 0 {
        existing.User = itemUnit.User
    }

    // Update only the fields that were provided
    cols := []string{}
    if itemUnit.Comment != "" {
        cols = append(cols, "Comment")
    }
    if itemUnit.Item != nil && itemUnit.Item.IdItem > 0 {
        cols = append(cols, "Item")
    }
    if itemUnit.StatusLookup != nil && itemUnit.StatusLookup.IdStatus > 0 {
        cols = append(cols, "StatusLookup")
    }
    if itemUnit.Warehouse != nil && itemUnit.Warehouse.IdWh > 0 {
        cols = append(cols, "Warehouse")
    }
    if itemUnit.CondLookup != nil && itemUnit.CondLookup.IdCondition > 0 {
        cols = append(cols, "CondLookup")
    }
    if itemUnit.User != nil && itemUnit.User.Id > 0 {
        cols = append(cols, "User")
    }

    _, err := o.Update(existing, cols...)
    return err
}


func (s *ItemUnitService) Delete(id uint) error {
    itemUnit := &models.ItemUnit{IdUnit: id}
    _, err := s.ormer.Delete(itemUnit)
    return err
}


// List retrieves item units with pagination
func (s *ItemUnitService) List(page, pageSize int, filters map[string]interface{}) ([]*models.ItemUnit, int64, error) {
    var items []*models.ItemUnit
    o := orm.NewOrm()
    
    qs := o.QueryTable("item_unit")
    
    // Apply filters with correct column names
    if filters != nil {
        if itemName, ok := filters["itemName"].(string); ok && itemName != "" {
            qs = qs.Filter("Item__item_name__icontains", itemName)
        }
        if serialNumber, ok := filters["serialNumber"].(string); ok && serialNumber != "" {
            qs = qs.Filter("serial_number__icontains", serialNumber)
        }
        if warehouseId, ok := filters["warehouseId"].(uint); ok && warehouseId > 0 {
            qs = qs.Filter("id_wh", warehouseId)
        }
        if status, ok := filters["status"].(string); ok && status != "" {
            statusValues := strings.Split(status, ",")
            qs = qs.Filter("status__in", statusValues)
        }
        if condition, ok := filters["condition"].(string); ok && condition != "" {
            conditionValues := strings.Split(condition, ",")
            qs = qs.Filter("condition__in", conditionValues)
        }
        if userId, ok := filters["userId"].(uint); ok && userId > 0 {
            qs = qs.Filter("updated_by", userId)
        }
    }
    
    total, _ := qs.Count()
    _, err := qs.Limit(pageSize).Offset((page - 1) * pageSize).All(&items)
    
    if len(items) > 0 {
        for _, item := range items {
            o.LoadRelated(item, "Item")
            o.LoadRelated(item, "Warehouse")
            o.LoadRelated(item, "StatusLookup")
            o.LoadRelated(item, "CondLookup")
            o.LoadRelated(item, "User")
        }
    }
    
    return items, total, err
}






// GetByWarehouse retrieves item units by warehouse ID
func (s *ItemUnitService) GetByWarehouse(warehouseId uint) ([]*models.ItemUnit, error) {
    var items []*models.ItemUnit
    
    nums, err := s.ormer.QueryTable(new(models.ItemUnit)).
        Filter("Warehouse__IdWh", warehouseId).
        RelatedSel().
        All(&items)
        
    logs.Info("Found %d items for warehouse %d", nums, warehouseId)
    
    return items, err
}

func (s *ItemUnitService) GetByItem(itemID uint) ([]*models.ItemUnit, error) {
    var itemUnits []*models.ItemUnit
    
    _, err := s.ormer.QueryTable(new(models.ItemUnit)).
        Filter("Item__IdItem", itemID).
        RelatedSel().
        All(&itemUnits)
        
    return itemUnits, err
}

func (s *ItemUnitService) GetByStatus(status uint8) ([]*models.ItemUnit, error) {
    var itemUnits []*models.ItemUnit
    
    _, err := s.ormer.QueryTable(new(models.ItemUnit)).
        Filter("StatusLookup__IdStatus", status).
        RelatedSel().
        All(&itemUnits)
        
    return itemUnits, err
}
