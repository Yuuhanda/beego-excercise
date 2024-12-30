package services

import (
    "errors"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "strings"
    "strconv"
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
        Filter("method", authItem.Method).
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

    // Verify path and method combination exists
    count, err = s.ormer.QueryTable("api_route").
        Filter("path", authItem.Path).
        Filter("method", authItem.Method).
        Count()
    if err != nil || count == 0 {
        return errors.New("path and method combination not found")
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

func (s *AuthItemService) CreateBulk(role string, paths []models.PathMethod) error {
    // Verify role exists first
    count, err := s.ormer.QueryTable("auth_roles").Filter("code", role).Count()
    if err != nil || count == 0 {
        return errors.New("role not found")
    }

    o := orm.NewOrm()
    
    // Insert new permissions
    for _, pathData := range paths {
        // Verify path and method combination exists
        pathExists := o.QueryTable("api_route").
            Filter("path", pathData.Path).
            Filter("method", pathData.Method).
            Exist()
        
        // Skip if path-method combination doesn't exist
        if !pathExists {
            continue
        }

        // Check if combination already exists
        exists := o.QueryTable(new(models.AuthItem)).
            Filter("role", role).
            Filter("path", pathData.Path).
            Filter("method", pathData.Method).
            Exist()
        
        if exists {
            continue // Skip this combination and continue with next one
        }

        authItem := &models.AuthItem{
            Role:   role,
            Path:   pathData.Path,
            Method: pathData.Method,
        }
        _, err = o.Insert(authItem)
        if err != nil {
            return err
        }
    }

    return nil
}

func (s *AuthItemService) CheckPermission(role, path, method string) (bool, error) {
    // First try exact match
    var count int
    sql := "SELECT COUNT(*) FROM auth_item WHERE role = ? AND path = ? AND method = ?"
    err := s.ormer.Raw(sql, role, path, method).QueryRow(&count)
    
    if count > 0 {
        return true, err
    }

    // Handle parameterized paths
    segments := strings.Split(path, "/")
    for i, segment := range segments {
        // Convert segments after specific keywords to their parameter names
        if i > 0 {
            switch segments[i-1] {
            case "warehouse":
                segments[i] = ":warehouseId"
            case "serial":
                segments[i] = ":serialNumber"
            case "visits":
                segments[i] = ":id"
            case "unit":
                segments[i] = ":unitId"
            default:
                // For numeric segments use :id
                if _, err := strconv.Atoi(segment); err == nil {
                    segments[i] = ":id"
                }
            }
        }
    }
    normalizedPath := strings.Join(segments, "/")
    
    // Try again with normalized path
    err = s.ormer.Raw(sql, role, normalizedPath, method).QueryRow(&count)
    
    return count > 0, err
}