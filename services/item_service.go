package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type ItemService struct {
    ormer orm.Ormer
}

func NewItemService() *ItemService {
    return &ItemService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new item
func (s *ItemService) Create(item *models.Item) error {
    _, err := s.ormer.Insert(item)
    return err
}

// GetByID retrieves item by ID
func (s *ItemService) GetByID(id uint) (*models.Item, error) {
    item := &models.Item{IdItem: id}
    err := s.ormer.Read(item)
    if err == orm.ErrNoRows {
        return nil, errors.New("item not found")
    }
    return item, err
}

// GetBySKU retrieves item by SKU
func (s *ItemService) GetBySKU(sku string) (*models.Item, error) {
    item := &models.Item{SKU: sku}
    err := s.ormer.Read(item, "SKU")
    if err == orm.ErrNoRows {
        return nil, errors.New("item not found")
    }
    return item, err
}

// Update updates item information
func (s *ItemService) Update(item *models.Item) error {
    if item.IdItem == 0 {
        return errors.New("item ID is required")
    }
    _, err := s.ormer.Update(item)
    return err
}

// Delete deletes an item
func (s *ItemService) Delete(id uint) error {
    item := &models.Item{IdItem: id}
    _, err := s.ormer.Delete(item)
    return err
}

// List retrieves items with pagination
func (s *ItemService) List(page, pageSize int) ([]*models.Item, int64, error) {
    var items []*models.Item
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.Item))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&items)
    return items, total, err
}

// ListByCategory retrieves items by category ID with pagination
func (s *ItemService) ListByCategory(categoryID int, page, pageSize int) ([]*models.Item, int64, error) {
    var items []*models.Item
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.Item)).Filter("id_category", categoryID)
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&items)
    return items, total, err
}
