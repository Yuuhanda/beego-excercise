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

// GetByID retrieves condition lookup by ID
func (s *ConditionLookupService) GetByID(id uint) (*models.ConditionLookup, error) {
    condition := &models.ConditionLookup{IdCondition: id}
    err := s.ormer.Read(condition)
    if err == orm.ErrNoRows {
        return nil, errors.New("condition not found")
    }
    return condition, err
}

// Delete deletes a condition lookup
func (s *ConditionLookupService) Delete(id uint) error {
    condition := &models.ConditionLookup{IdCondition: id}
    _, err := s.ormer.Delete(condition)
    return err
}
