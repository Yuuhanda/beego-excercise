package services

import (
    "errors"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type UnitLogService struct {
    ormer orm.Ormer
}

func NewUnitLogService() *UnitLogService {
    return &UnitLogService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new unit log
func (s *UnitLogService) Create(unitLog *models.UnitLog) error {
    unitLog.UpdateAt = time.Now()
    _, err := s.ormer.Insert(unitLog)
    return err
}

// GetByID retrieves unit log by ID
func (s *UnitLogService) GetByID(id int) (*models.UnitLog, error) {
    unitLog := &models.UnitLog{IdLog: id}
    err := s.ormer.Read(unitLog)
    if err == orm.ErrNoRows {
        return nil, errors.New("unit log not found")
    }
    return unitLog, err
}

// GetByUnitID retrieves all logs for a specific unit
func (s *UnitLogService) GetByUnitID(unitID int) ([]*models.UnitLog, error) {
    var logs []*models.UnitLog
    _, err := s.ormer.QueryTable(new(models.UnitLog)).Filter("IdUnit", unitID).All(&logs)
    return logs, err
}

// Update updates unit log information
func (s *UnitLogService) Update(unitLog *models.UnitLog) error {
    if unitLog.IdLog == 0 {
        return errors.New("unit log ID is required")
    }
    unitLog.UpdateAt = time.Now()
    _, err := s.ormer.Update(unitLog)
    return err
}

// Delete deletes a unit log
func (s *UnitLogService) Delete(id int) error {
    unitLog := &models.UnitLog{IdLog: id}
    _, err := s.ormer.Delete(unitLog)
    return err
}

// List retrieves unit logs with pagination
func (s *UnitLogService) List(page, pageSize int) ([]*models.UnitLog, int64, error) {
    var logs []*models.UnitLog
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.UnitLog))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&logs)
    return logs, total, err
}
