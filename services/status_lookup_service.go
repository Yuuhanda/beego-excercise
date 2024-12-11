package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type StatusLookupService struct {
    ormer orm.Ormer
}

func NewStatusLookupService() *StatusLookupService {
    return &StatusLookupService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new status
func (s *StatusLookupService) Create(status *models.StatusLookup) error {
    _, err := s.ormer.Insert(status)
    return err
}

// GetByID retrieves status by ID
func (s *StatusLookupService) GetByID(id uint8) (*models.StatusLookup, error) {
    status := &models.StatusLookup{IdStatus: id}
    err := s.ormer.Read(status)
    if err == orm.ErrNoRows {
        return nil, errors.New("status not found")
    }
    return status, err
}

// GetByName retrieves status by name
func (s *StatusLookupService) GetByName(name string) (*models.StatusLookup, error) {
    status := &models.StatusLookup{StatusName: name}
    err := s.ormer.Read(status, "StatusName")
    if err == orm.ErrNoRows {
        return nil, errors.New("status not found")
    }
    return status, err
}

// Update updates status information
func (s *StatusLookupService) Update(status *models.StatusLookup) error {
    if status.IdStatus == 0 {
        return errors.New("status ID is required")
    }
    _, err := s.ormer.Update(status)
    return err
}

// Delete deletes a status
func (s *StatusLookupService) Delete(id uint8) error {
    status := &models.StatusLookup{IdStatus: id}
    _, err := s.ormer.Delete(status)
    return err
}

// List retrieves all statuses
func (s *StatusLookupService) List() ([]*models.StatusLookup, error) {
    var statuses []*models.StatusLookup
    _, err := s.ormer.QueryTable(new(models.StatusLookup)).All(&statuses)
    return statuses, err
}
