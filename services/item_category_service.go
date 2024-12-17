package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "fmt"
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
func (s *ItemCategoryService) Create(category *models.ItemCategory) (string, error) {
    // Check if category with same code exists
    existing := &models.ItemCategory{CatCode: category.CatCode}
    err := s.ormer.Read(existing, "CatCode")
    if err == nil {
        return "", fmt.Errorf("failed to create category: code %s already exists", category.CatCode)
    }
    if err != orm.ErrNoRows {
        return "", err
    }
    
    _, err = s.ormer.Insert(category)
    if err != nil {
        return "", err
    }
    
    return fmt.Sprintf("category %s created successfully", category.CatCode), nil
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
    o := orm.NewOrm()
    _, err := o.Update(category)
    return err
}


// Delete deletes an item category
func (s *ItemCategoryService) Delete(category *models.ItemCategory) error {
    o := orm.NewOrm()
    _, err := o.Delete(category)
    return err
}


// List retrieves item categories with pagination
func (s *ItemCategoryService) List(page, pageSize int, name, code string) ([]*models.ItemCategory, int64, error) {
    var categories []*models.ItemCategory
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.ItemCategory))
    
    if name != "" {
        qs = qs.Filter("category_name__icontains", name)  // Using snake_case for database column
    }
    if code != "" {
        qs = qs.Filter("cat_code__icontains", code)  // Using snake_case for database column
    }
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&categories)
    return categories, total, err
}


