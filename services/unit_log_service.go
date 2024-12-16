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
    o := orm.NewOrm()

    // Verify and load the unit
    unit := &models.ItemUnit{IdUnit: unitLog.IdUnit.IdUnit}
    if err := o.Read(unit); err != nil {
        return errors.New("invalid unit ID")
    }
    unitLog.IdUnit = unit

    // Insert the log
    _, err := o.Insert(unitLog)
    if err != nil {
        return err
    }

    // Load related data
    o.LoadRelated(unitLog.IdUnit, "Item")
    o.LoadRelated(unitLog.IdUnit, "StatusLookup")
    o.LoadRelated(unitLog.IdUnit, "Warehouse")
    o.LoadRelated(unitLog.IdUnit, "CondLookup")
    o.LoadRelated(unitLog.IdUnit, "User")

    if unitLog.IdUnit.Item != nil {
        o.LoadRelated(unitLog.IdUnit.Item, "Category")
    }

    return nil
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
func (s *UnitLogService) List(page, pageSize int, filters map[string]string) ([]*models.UnitLog, int64, error) {
    var unitLogs []*models.UnitLog
    offset := (page - 1) * pageSize

    o := orm.NewOrm()
    qs := o.QueryTable(new(models.UnitLog)).RelatedSel()

    // Apply filters
    if serialNumber := filters["serial_number"]; serialNumber != "" {
        qs = qs.Filter("IdUnit__SerialNumber__icontains", serialNumber)
    }

    if content := filters["content"]; content != "" {
        qs = qs.Filter("Content__icontains", content)
    }

    if startDate := filters["start_date"]; startDate != "" {
        qs = qs.Filter("UpdateAt__gte", startDate)
    }

    if endDate := filters["end_date"]; endDate != "" {
        qs = qs.Filter("UpdateAt__lte", endDate)
    }

    if actorsAction := filters["actors_action"]; actorsAction != "" {
        qs = qs.Filter("ActorsAction__icontains", actorsAction)
    }

    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }

    _, err = qs.Offset(offset).Limit(pageSize).All(&unitLogs)
    if err != nil {
        return nil, 0, err
    }

    // Load related data
    for _, unitLog := range unitLogs {
        if unitLog.IdUnit != nil {
            o.LoadRelated(unitLog.IdUnit, "Item")
            o.LoadRelated(unitLog.IdUnit, "StatusLookup")
            o.LoadRelated(unitLog.IdUnit, "Warehouse")
            o.LoadRelated(unitLog.IdUnit, "CondLookup")
            o.LoadRelated(unitLog.IdUnit, "User")

            if unitLog.IdUnit.Item != nil {
                o.LoadRelated(unitLog.IdUnit.Item, "Category")
            }
        }
    }

    return unitLogs, total, nil
}


