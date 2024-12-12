package services

import (
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

func (s *ItemService) List(page, pageSize int, filters map[string]interface{}) ([]models.Item, int64, error) {
    var items []models.Item
    o := orm.NewOrm()
    
    qs := o.QueryTable(new(models.Item))
    
    // Apply filters using exact field names from Item model
    if name, ok := filters["name"].(string); ok && name != "" {
        qs = qs.Filter("ItemName__icontains", name)
    }
    if category, ok := filters["category"].(string); ok && category != "" {
        qs = qs.Filter("Category__CategoryName__icontains", category)
    }
    if sku, ok := filters["sku"].(string); ok && sku != "" {
        qs = qs.Filter("SKU__icontains", sku)
    }
    if status, ok := filters["status"].(uint); ok && status > 0 {
        qs = qs.Filter("Status", status)
    }
    
    total, _ := qs.Count()
    _, err := qs.RelatedSel("Category").Limit(pageSize).Offset((page - 1) * pageSize).All(&items)
    
    return items, total, err
}

func (s *ItemService) ListWithCategories(page, pageSize int) ([]models.Item, int64, error) {
    var items []models.Item
    var total int64
    
    offset := (page - 1) * pageSize
    
    // Join items with categories table
    qs := s.ormer.QueryTable("item").RelatedSel("Category")
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Limit(pageSize, offset).All(&items)
    if err != nil {
        return nil, 0, err
    }
    
    return items, total, nil
}

func (s *ItemService) GetByID(id uint) (*models.Item, error) {
    var item models.Item
    o := orm.NewOrm()
    
    qs := o.QueryTable(new(models.Item))
    err := qs.Filter("IdItem", id).RelatedSel("Category").One(&item)
    
    return &item, err
}


func (s *ItemService) Create(item *models.Item) error {
    o := orm.NewOrm()
    // Create the item
    _, err := o.Insert(item)
    if err != nil {
        return err
    }
    // Load the complete category data
    o.LoadRelated(item, "Category")
    return nil
}


func (s *ItemService) Update(item *models.Item) error {
    _, err := s.ormer.Update(item)
    return err
}

func (s *ItemService) Delete(id uint) error {
    _, err := s.ormer.Delete(&models.Item{IdItem: id})
    return err
}
