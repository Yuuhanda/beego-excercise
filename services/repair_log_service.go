package services

import (
    "errors"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type RepairLogService struct {
    ormer orm.Ormer
}

func NewRepairLogService() *RepairLogService {
    return &RepairLogService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new repair log
func (s *RepairLogService) Create(repairLog *models.RepairLog) error {
    repairLog.Datetime = time.Now()
    _, err := s.ormer.Insert(repairLog)
    return err
}

// GetByID retrieves repair log by ID
func (s *RepairLogService) GetByID(id int) (*models.RepairLog, error) {
    repairLog := &models.RepairLog{IdRepair: id}
    err := s.ormer.Read(repairLog)
    if err == orm.ErrNoRows {
        return nil, errors.New("repair log not found")
    }
    return repairLog, err
}

// GetByUnit retrieves repair logs by unit ID
func (s *RepairLogService) GetByUnit(unitID int) ([]*models.RepairLog, error) {
    var logs []*models.RepairLog
    _, err := s.ormer.QueryTable(new(models.RepairLog)).Filter("id_unit", unitID).All(&logs)
    return logs, err
}

// Update updates repair log information
func (s *RepairLogService) Update(repairLog *models.RepairLog) error {
    if repairLog.IdRepair == 0 {
        return errors.New("repair log ID is required")
    }
    _, err := s.ormer.Update(repairLog)
    return err
}

// Delete deletes a repair log
func (s *RepairLogService) Delete(id int) error {
    repairLog := &models.RepairLog{IdRepair: id}
    _, err := s.ormer.Delete(repairLog)
    return err
}

// List retrieves repair logs with pagination
func (s *RepairLogService) List(page, pageSize int) ([]*models.RepairLog, int64, error) {
    var logs []*models.RepairLog
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.RepairLog))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.OrderBy("-datetime").Offset(offset).Limit(pageSize).All(&logs)
    return logs, total, err
}
