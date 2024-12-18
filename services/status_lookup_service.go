package services

import (
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "errors"
)

type StatusLookupService struct {
    ormer orm.Ormer
}

func NewStatusLookupService() *StatusLookupService {
    return &StatusLookupService{
        ormer: orm.NewOrm(),
    }
}

func (s *StatusLookupService) GetByID(id uint) (*models.StatusLookup, error) {
    status := &models.StatusLookup{IdStatus: id}
    err := s.ormer.Read(status)
    if err == orm.ErrNoRows {
        return nil, errors.New("status not found")
    }
    return status, err
}

func (s *StatusLookupService) GetAll() ([]*models.StatusLookup, error) {
    var statuses []*models.StatusLookup
    _, err := s.ormer.QueryTable(new(models.StatusLookup)).All(&statuses)
    return statuses, err
}

func (s *StatusLookupService) Create(status *models.StatusLookup) error {
    _, err := s.ormer.Insert(status)
    return err
}

func (s *StatusLookupService) Update(status *models.StatusLookup) error {
    _, err := s.ormer.Update(status, "StatusName")
    return err
}

func (s *StatusLookupService) Delete(id uint) error {
    status := &models.StatusLookup{IdStatus: id}
    _, err := s.ormer.Delete(status)
    return err
}
