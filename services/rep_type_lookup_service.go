package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type RepTypeLookupService struct {
    ormer orm.Ormer
}

func NewRepTypeLookupService() *RepTypeLookupService {
    return &RepTypeLookupService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new repair type lookup
func (s *RepTypeLookupService) Create(repType *models.RepTypeLookup) error {
    _, err := s.ormer.Insert(repType)
    return err
}

// GetByID retrieves repair type by ID
func (s *RepTypeLookupService) GetByID(id int8) (*models.RepTypeLookup, error) {
    repType := &models.RepTypeLookup{IdRepT: id}
    err := s.ormer.Read(repType)
    if err == orm.ErrNoRows {
        return nil, errors.New("repair type not found")
    }
    return repType, err
}

// Update updates repair type information
func (s *RepTypeLookupService) Update(repType *models.RepTypeLookup) error {
    if repType.IdRepT == 0 {
        return errors.New("repair type ID is required")
    }
    _, err := s.ormer.Update(repType)
    return err
}

// Delete deletes a repair type
func (s *RepTypeLookupService) Delete(id int8) error {
    repType := &models.RepTypeLookup{IdRepT: id}
    _, err := s.ormer.Delete(repType)
    return err
}

// List retrieves all repair types
func (s *RepTypeLookupService) List() ([]*models.RepTypeLookup, error) {
    var repTypes []*models.RepTypeLookup
    _, err := s.ormer.QueryTable(new(models.RepTypeLookup)).All(&repTypes)
    return repTypes, err
}

// GetByRepType retrieves repair type by name
func (s *RepTypeLookupService) GetByRepType(repType string) (*models.RepTypeLookup, error) {
    repTypeLookup := &models.RepTypeLookup{RepType: repType}
    err := s.ormer.Read(repTypeLookup, "RepType")
    if err == orm.ErrNoRows {
        return nil, errors.New("repair type not found")
    }
    return repTypeLookup, err
}
