package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type ItemUnitService struct {
    ormer orm.Ormer
}

func NewItemUnitService() *ItemUnitService {
    return &ItemUnitService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new item unit
func (s *ItemUnitService) Create(itemUnit *models.ItemUnit) error {
    _, err := s.ormer.Insert(itemUnit)
    return err
}

// GetByID retrieves item unit by ID
func (s *ItemUnitService) GetByID(id uint) (*models.ItemUnit, error) {
    itemUnit := &models.ItemUnit{IdUnit: id}
    err := s.ormer.Read(itemUnit)
    if err == orm.ErrNoRows {
        return nil, errors.New("item unit not found")
    }
    return itemUnit, err
}

// GetBySerialNumber retrieves item unit by serial number
func (s *ItemUnitService) GetBySerialNumber(serialNumber string) (*models.ItemUnit, error) {
    itemUnit := &models.ItemUnit{SerialNumber: serialNumber}
    err := s.ormer.Read(itemUnit, "SerialNumber")
    if err == orm.ErrNoRows {
        return nil, errors.New("item unit not found")
    }
    return itemUnit, err
}

// Update updates item unit information
func (s *ItemUnitService) Update(itemUnit *models.ItemUnit) error {
    if itemUnit.IdUnit == 0 {
        return errors.New("item unit ID is required")
    }
    _, err := s.ormer.Update(itemUnit)
    return err
}

// Delete deletes an item unit
func (s *ItemUnitService) Delete(id uint) error {
    itemUnit := &models.ItemUnit{IdUnit: id}
    _, err := s.ormer.Delete(itemUnit)
    return err
}

// List retrieves item units with pagination
func (s *ItemUnitService) List(page, pageSize int) ([]*models.ItemUnit, int64, error) {
    var itemUnits []*models.ItemUnit
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.ItemUnit))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&itemUnits)
    return itemUnits, total, err
}

// GetByWarehouse retrieves item units by warehouse ID
func (s *ItemUnitService) GetByWarehouse(warehouseID uint) ([]*models.ItemUnit, error) {
    var itemUnits []*models.ItemUnit
    _, err := s.ormer.QueryTable(new(models.ItemUnit)).Filter("id_wh", warehouseID).All(&itemUnits)
    return itemUnits, err
}

// GetByItem retrieves item units by item ID
func (s *ItemUnitService) GetByItem(itemID uint) ([]*models.ItemUnit, error) {
    var itemUnits []*models.ItemUnit
    _, err := s.ormer.QueryTable(new(models.ItemUnit)).Filter("id_item", itemID).All(&itemUnits)
    return itemUnits, err
}

// GetByStatus retrieves item units by status
func (s *ItemUnitService) GetByStatus(status uint8) ([]*models.ItemUnit, error) {
    var itemUnits []*models.ItemUnit
    _, err := s.ormer.QueryTable(new(models.ItemUnit)).Filter("status", status).All(&itemUnits)
    return itemUnits, err
}
