package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type DocUploadedService struct {
    ormer orm.Ormer
}

func NewDocUploadedService() *DocUploadedService {
    return &DocUploadedService{
        ormer: orm.NewOrm(),
    }
}

func (s *DocUploadedService) Create(doc *models.DocUploaded) error {
    o := orm.NewOrm()

    // Verify and load the user
    user := &models.User{Id: doc.UserId.Id}
    if err := o.Read(user); err != nil {
        return errors.New("invalid user ID")
    }
    doc.UserId = user

    _, err := o.Insert(doc)
    if err != nil {
        return err
    }

    // Load related data
    o.LoadRelated(doc, "UserId")
    return nil
}

func (s *DocUploadedService) GetByID(id int) (*models.DocUploaded, error) {
    o := orm.NewOrm()
    doc := &models.DocUploaded{IdDoc: id}

    err := o.QueryTable(new(models.DocUploaded)).Filter("IdDoc", id).RelatedSel().One(doc)
    if err == orm.ErrNoRows {
        return nil, errors.New("document not found")
    }
    if err != nil {
        return nil, err
    }

    // Load related user data
    if doc.UserId != nil {
        o.LoadRelated(doc, "UserId")
    }

    return doc, nil
}


func (s *DocUploadedService) List(page, pageSize int, filters map[string]string) ([]*models.DocUploaded, int64, error) {
    var docs []*models.DocUploaded
    offset := (page - 1) * pageSize

    o := orm.NewOrm()
    qs := o.QueryTable(new(models.DocUploaded)).RelatedSel()

    // Apply filters
    if fileName := filters["file_name"]; fileName != "" {
        qs = qs.Filter("FileName__icontains", fileName)
    }

    if userId := filters["user_id"]; userId != "" {
        qs = qs.Filter("UserId__Id", userId)
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

    _, err = qs.OrderBy("-Datetime").Offset(offset).Limit(pageSize).All(&docs)
    if err != nil {
        return nil, 0, err
    }

    return docs, total, nil
}

func (s *DocUploadedService) Delete(id int) error {
    o := orm.NewOrm()
    doc := &models.DocUploaded{IdDoc: id}
    
    if err := o.Read(doc); err != nil {
        return errors.New("document not found")
    }
    
    _, err := o.Delete(doc)
    return err
}
