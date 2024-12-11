package services

import (
    "errors"
    "time"
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

// Create creates a new document upload record
func (s *DocUploadedService) Create(doc *models.DocUploaded) error {
    doc.Datetime = time.Now()
    _, err := s.ormer.Insert(doc)
    return err
}

// GetByID retrieves document by ID
func (s *DocUploadedService) GetByID(id int) (*models.DocUploaded, error) {
    doc := &models.DocUploaded{IdDoc: id}
    err := s.ormer.Read(doc)
    if err == orm.ErrNoRows {
        return nil, errors.New("document not found")
    }
    return doc, err
}

// GetByUserID retrieves all documents for a specific user
func (s *DocUploadedService) GetByUserID(userID int) ([]*models.DocUploaded, error) {
    var docs []*models.DocUploaded
    _, err := s.ormer.QueryTable(new(models.DocUploaded)).Filter("UserId", userID).All(&docs)
    return docs, err
}

// Update updates document information
func (s *DocUploadedService) Update(doc *models.DocUploaded) error {
    if doc.IdDoc == 0 {
        return errors.New("document ID is required")
    }
    _, err := s.ormer.Update(doc)
    return err
}

// Delete deletes a document
func (s *DocUploadedService) Delete(id int) error {
    doc := &models.DocUploaded{IdDoc: id}
    _, err := s.ormer.Delete(doc)
    return err
}

// List retrieves documents with pagination
func (s *DocUploadedService) List(page, pageSize int) ([]*models.DocUploaded, int64, error) {
    var docs []*models.DocUploaded
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.DocUploaded))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.OrderBy("-Datetime").Offset(offset).Limit(pageSize).All(&docs)
    return docs, total, err
}

// GetDocumentsByDateRange retrieves documents within a specific date range
func (s *DocUploadedService) GetDocumentsByDateRange(startDate, endDate time.Time) ([]*models.DocUploaded, error) {
    var docs []*models.DocUploaded
    _, err := s.ormer.QueryTable(new(models.DocUploaded)).
        Filter("Datetime__gte", startDate).
        Filter("Datetime__lte", endDate).
        OrderBy("-Datetime").
        All(&docs)
    return docs, err
}
