package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    //filter 
)

type AuthItemService struct {
    ormer orm.Ormer
}

func NewAuthItemService() *AuthItemService {
    return &AuthItemService{
        ormer: orm.NewOrm(),
    }
}

func (s *AuthItemService) Create(authItem *models.AuthItem) error {
    // Check for existing identical entry
    count, err := s.ormer.QueryTable(new(models.AuthItem)).
        Filter("role", authItem.Role).
        Filter("path", authItem.Path).
        Count()
    if err != nil {
        return err
    }
    if count > 0 {
        return errors.New("this role already has access to this path")
    }

    // Verify role exists
    count, err = s.ormer.QueryTable("auth_roles").Filter("code", authItem.Role).Count()
    if err != nil || count == 0 {
        return errors.New("role not found")
    }

    // Verify path exists
    count, err = s.ormer.QueryTable("api_route").Filter("path", authItem.Path).Count()
    if err != nil || count == 0 {
        return errors.New("path not found")
    }

    _, err = s.ormer.Insert(authItem)
    return err
}

func (s *AuthItemService) GetByID(id int) (*models.AuthItem, error) {
    authItem := &models.AuthItem{Id: id}
    err := s.ormer.Read(authItem)
    if err == orm.ErrNoRows {
        return nil, errors.New("auth item not found")
    }
    return authItem, err
}

func (s *AuthItemService) List(page, pageSize int) ([]*models.AuthItem, int64, error) {
    var authItems []*models.AuthItem
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.AuthItem))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&authItems)
    return authItems, total, err
}

func (s *AuthItemService) Update(authItem *models.AuthItem) error {
    if authItem.Id == 0 {
        return errors.New("auth item ID is required")
    }

    // Verify role exists
    var count int64
    count, err := s.ormer.QueryTable("auth_roles").Filter("code", authItem.Role).Count()
    if err != nil || count == 0 {
        return errors.New("role not found")
    }

    // Verify path exists
    count, err = s.ormer.QueryTable("api_route").Filter("path", authItem.Path).Count()
    if err != nil || count == 0 {
        return errors.New("path not found")
    }

    _, err = s.ormer.Update(authItem)
    return err
}

func (s *AuthItemService) Delete(id int) error {
    // Check if auth item exists
    authItem := &models.AuthItem{Id: id}
    err := s.ormer.Read(authItem)
    if err == orm.ErrNoRows {
        return errors.New("no auth item found with this id")
    }
    if err != nil {
        return err
    }

    // Proceed with deletion
    _, err = s.ormer.Delete(authItem)
    return err
}

func (s *AuthItemService) CreateBulk(role string, paths []string) error {
    // Verify role exists first
    count, err := s.ormer.QueryTable("auth_roles").Filter("code", role).Count()
    if err != nil || count == 0 {
        return errors.New("role not found")
    }

    // Use single ormer instance
    o := orm.NewOrm()
    
    for _, path := range paths {
        // Check if combination already exists
        exists := o.QueryTable(new(models.AuthItem)).
            Filter("role", role).
            Filter("path", path).
            Exist()
        if exists {
            continue
        }

        // Create new auth item
        authItem := &models.AuthItem{
            Role: role,
            Path: path,
        }
        _, err = o.Insert(authItem)
        if err != nil {
            return err
        }
    }

    return nil
}

