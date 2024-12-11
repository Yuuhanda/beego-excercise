package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type ItemCategoryService struct {
    ormer orm.Ormer
}

func NewItemCategoryService() *ItemCategoryService {
    return &ItemCategoryService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new item category
func (s *ItemCategoryService) Create(category *models.ItemCategory) error {
    _, err := s.ormer.Insert(category)
    return err
}

// GetByID retrieves item category by ID
func (s *ItemCategoryService) GetByID(id int) (*models.ItemCategory, error) {
    category := &models.ItemCategory{IdCategory: id}
    err := s.ormer.Read(category)
    if err == orm.ErrNoRows {
        return nil, errors.New("category not found")
    }
    return category, err
}

// GetByCatCode retrieves item category by category code
func (s *ItemCategoryService) GetByCatCode(code string) (*models.ItemCategory, error) {
    category := &models.ItemCategory{CatCode: code}
    err := s.ormer.Read(category, "CatCode")
    if err == orm.ErrNoRows {
        return nil, errors.New("category not found")
    }
    return category, err
}

// Update updates item category information
func (s *ItemCategoryService) Update(category *models.ItemCategory) error {
    if category.IdCategory == 0 {
        return errors.New("category ID is required")
    }
    _, err := s.ormer.Update(category)
    return err
}

// Delete deletes an item category
func (s *ItemCategoryService) Delete(id int) error {
    category := &models.ItemCategory{IdCategory: id}
    _, err := s.ormer.Delete(category)
    return err
}

// List retrieves item categories with pagination
func (s *ItemCategoryService) List(page, pageSize int) ([]*models.ItemCategory, int64, error) {
    var categories []*models.ItemCategory
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.ItemCategory))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&categories)
    return categories, total, err
}
