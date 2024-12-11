package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type ConditionLookupService struct {
    ormer orm.Ormer
}

func NewConditionLookupService() *ConditionLookupService {
    return &ConditionLookupService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new condition lookup
func (s *ConditionLookupService) Create(condition *models.ConditionLookup) error {
    _, err := s.ormer.Insert(condition)
    return err
}

// GetByID retrieves condition lookup by ID
func (s *ConditionLookupService) GetByID(id uint8) (*models.ConditionLookup, error) {
    condition := &models.ConditionLookup{IdCondition: id}
    err := s.ormer.Read(condition)
    if err == orm.ErrNoRows {
        return nil, errors.New("condition lookup not found")
    }
    return condition, err
}

// GetByName retrieves condition lookup by name
func (s *ConditionLookupService) GetByName(name string) (*models.ConditionLookup, error) {
    condition := &models.ConditionLookup{ConditionName: name}
    err := s.ormer.Read(condition, "ConditionName")
    if err == orm.ErrNoRows {
        return nil, errors.New("condition lookup not found")
    }
    return condition, err
}

// Update updates condition lookup information
func (s *ConditionLookupService) Update(condition *models.ConditionLookup) error {
    if condition.IdCondition == 0 {
        return errors.New("condition ID is required")
    }
    _, err := s.ormer.Update(condition)
    return err
}

// Delete deletes a condition lookup
func (s *ConditionLookupService) Delete(id uint8) error {
    condition := &models.ConditionLookup{IdCondition: id}
    _, err := s.ormer.Delete(condition)
    return err
}

// List retrieves all condition lookups with pagination
func (s *ConditionLookupService) List(page, pageSize int) ([]*models.ConditionLookup, int64, error) {
    var conditions []*models.ConditionLookup
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.ConditionLookup))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&conditions)
    return conditions, total, err
}
