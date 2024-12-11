package services

import (
    "errors"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type LendingService struct {
    ormer orm.Ormer
}

func NewLendingService() *LendingService {
    return &LendingService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new lending record
func (s *LendingService) Create(lending *models.Lending) error {
    lending.Date = time.Now()
    _, err := s.ormer.Insert(lending)
    return err
}

// GetByID retrieves lending by ID
func (s *LendingService) GetByID(id uint) (*models.Lending, error) {
    lending := &models.Lending{IdLending: id}
    err := s.ormer.Read(lending)
    if err == orm.ErrNoRows {
        return nil, errors.New("lending record not found")
    }
    return lending, err
}

// GetByUserID retrieves lendings by user ID
func (s *LendingService) GetByUserID(userID int) ([]*models.Lending, error) {
    var lendings []*models.Lending
    _, err := s.ormer.QueryTable(new(models.Lending)).Filter("UserId", userID).All(&lendings)
    return lendings, err
}

// Update updates lending information
func (s *LendingService) Update(lending *models.Lending) error {
    if lending.IdLending == 0 {
        return errors.New("lending ID is required")
    }
    _, err := s.ormer.Update(lending)
    return err
}

// Delete deletes a lending record
func (s *LendingService) Delete(id uint) error {
    lending := &models.Lending{IdLending: id}
    _, err := s.ormer.Delete(lending)
    return err
}

// List retrieves lendings with pagination
func (s *LendingService) List(page, pageSize int) ([]*models.Lending, int64, error) {
    var lendings []*models.Lending
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.Lending))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&lendings)
    return lendings, total, err
}

// GetActiveLoans retrieves all active loans (Type = 1)
func (s *LendingService) GetActiveLoans() ([]*models.Lending, error) {
    var lendings []*models.Lending
    _, err := s.ormer.QueryTable(new(models.Lending)).Filter("Type", 1).All(&lendings)
    return lendings, err
}

// GetReturnedLoans retrieves all returned loans (Type = 2)
func (s *LendingService) GetReturnedLoans() ([]*models.Lending, error) {
    var lendings []*models.Lending
    _, err := s.ormer.QueryTable(new(models.Lending)).Filter("Type", 2).All(&lendings)
    return lendings, err
}
