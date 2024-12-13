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
    o := orm.NewOrm()
    repairLog := &models.RepairLog{IdRepair: id}
    
    err := o.QueryTable(repairLog).Filter("IdRepair", id).RelatedSel().One(repairLog)
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
func (s *RepairLogService) List(page, pageSize int, filters map[string]string) ([]*models.RepairLog, int64, error) {
    var logs []*models.RepairLog
    offset := (page - 1) * pageSize

    o := orm.NewOrm()
    qs := o.QueryTable(new(models.RepairLog)).RelatedSel()

    // Apply filters
    if itemName := filters["item_name"]; itemName != "" {
        qs = qs.Filter("IdUnit__Item__ItemName__icontains", itemName)
    }

    if serialNumber := filters["serial_number"]; serialNumber != "" {
        qs = qs.Filter("IdUnit__SerialNumber__icontains", serialNumber)
    }

    if sku := filters["sku"]; sku != "" {
        qs = qs.Filter("IdUnit__Item__SKU__icontains", sku)
    }

    if warehouse := filters["warehouse"]; warehouse != "" {
        qs = qs.Filter("IdUnit__Warehouse__IdWarehouse", warehouse)
    }

    if repType := filters["rep_type"]; repType != "" {
        qs = qs.Filter("RepType__IdRepT", repType)
    }

    if startDate := filters["start_date"]; startDate != "" {
        qs = qs.Filter("Datetime__gte", startDate)
    }

    if endDate := filters["end_date"]; endDate != "" {
        qs = qs.Filter("Datetime__lte", endDate)
    }

    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }

    _, err = qs.OrderBy("-Datetime").Offset(offset).Limit(pageSize).All(&logs)
    if err != nil {
        return nil, 0, err
    }

    // Load related data
    for _, log := range logs {
        if log.IdUnit != nil {
            o.LoadRelated(log.IdUnit, "Item")
            o.LoadRelated(log.IdUnit, "StatusLookup")
            o.LoadRelated(log.IdUnit, "Warehouse")
            o.LoadRelated(log.IdUnit, "CondLookup")
            o.LoadRelated(log.IdUnit, "User")

            if log.IdUnit.Item != nil {
                o.LoadRelated(log.IdUnit.Item, "Category")
            }
        }
    }

    return logs, total, nil
}


