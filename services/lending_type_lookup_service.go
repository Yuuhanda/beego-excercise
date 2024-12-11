package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type LendingTypeLookupService struct {
    ormer orm.Ormer
}

func NewLendingTypeLookupService() *LendingTypeLookupService {
    return &LendingTypeLookupService{
        ormer: orm.NewOrm(),
    }
}

func (s *LendingTypeLookupService) Create(lendingType *models.LendingTypeLookup) error {
    _, err := s.ormer.Insert(lendingType)
    return err
}

func (s *LendingTypeLookupService) GetByID(id uint8) (*models.LendingTypeLookup, error) {
    lendingType := &models.LendingTypeLookup{IdType: id}
    err := s.ormer.Read(lendingType)
    if err == orm.ErrNoRows {
        return nil, errors.New("lending type not found")
    }
    return lendingType, err
}

func (s *LendingTypeLookupService) GetByTypeName(typeName string) (*models.LendingTypeLookup, error) {
    lendingType := &models.LendingTypeLookup{TypeName: typeName}
    err := s.ormer.Read(lendingType, "TypeName")
    if err == orm.ErrNoRows {
        return nil, errors.New("lending type not found")
    }
    return lendingType, err
}

func (s *LendingTypeLookupService) Update(lendingType *models.LendingTypeLookup) error {
    if lendingType.IdType == 0 {
        return errors.New("lending type ID is required")
    }
    _, err := s.ormer.Update(lendingType)
    return err
}

func (s *LendingTypeLookupService) Delete(id uint8) error {
    lendingType := &models.LendingTypeLookup{IdType: id}
    _, err := s.ormer.Delete(lendingType)
    return err
}

func (s *LendingTypeLookupService) List() ([]*models.LendingTypeLookup, error) {
    var lendingTypes []*models.LendingTypeLookup
    _, err := s.ormer.QueryTable(new(models.LendingTypeLookup)).All(&lendingTypes)
    return lendingTypes, err
}
