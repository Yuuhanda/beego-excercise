package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type WarehouseService struct {
    ormer orm.Ormer
}

func NewWarehouseService() *WarehouseService {
    return &WarehouseService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new warehouse
func (s *WarehouseService) Create(warehouse *models.Warehouse) error {
    _, err := s.ormer.Insert(warehouse)
    return err
}

// GetByID retrieves warehouse by ID
func (s *WarehouseService) GetByID(id uint) (*models.Warehouse, error) {
    warehouse := &models.Warehouse{IdWh: id}
    err := s.ormer.Read(warehouse)
    if err == orm.ErrNoRows {
        return nil, errors.New("warehouse not found")
    }
    return warehouse, err
}

// GetByName retrieves warehouse by name
func (s *WarehouseService) GetByName(name string) (*models.Warehouse, error) {
    warehouse := &models.Warehouse{WhName: name}
    err := s.ormer.Read(warehouse, "WhName")
    if err == orm.ErrNoRows {
        return nil, errors.New("warehouse not found")
    }
    return warehouse, err
}

// Update updates warehouse information
func (s *WarehouseService) Update(warehouse *models.Warehouse) error {
    if warehouse.IdWh == 0 {
        return errors.New("warehouse ID is required")
    }
    _, err := s.ormer.Update(warehouse)
    return err
}

// Delete deletes a warehouse
func (s *WarehouseService) Delete(id uint) error {
    warehouse := &models.Warehouse{IdWh: id}
    _, err := s.ormer.Delete(warehouse)
    return err
}

// List retrieves warehouses with pagination
// List retrieves warehouses with pagination and filters
func (s *WarehouseService) List(page, pageSize int, whName, whAddress string) ([]*models.Warehouse, int64, error) {
    var warehouses []*models.Warehouse
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.Warehouse))
    
    // Apply filters if provided
    if whName != "" {
        qs = qs.Filter("wh_name__icontains", whName)
    }
    if whAddress != "" {
        qs = qs.Filter("wh_address__icontains", whAddress)
    }
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&warehouses)
    return warehouses, total, err
}

// GetWarehouseUsers retrieves all users associated with a warehouse
func (s *WarehouseService) GetWarehouseUsers(warehouseID uint) ([]*models.User, error) {
    warehouse := &models.Warehouse{IdWh: warehouseID}
    err := s.ormer.Read(warehouse)
    if err != nil {
        return nil, err
    }
    
    _, err = s.ormer.LoadRelated(warehouse, "Users")
    if err != nil {
        return nil, err
    }
    
    return warehouse.Users, nil
}
