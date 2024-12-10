package services

import (
    "errors"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type UserService struct {
    ormer orm.Ormer
}

func NewUserService() *UserService {
    return &UserService{
        ormer: orm.NewOrm(),
    }
}

// Create creates a new user
func (s *UserService) Create(user *models.User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    if user.Status == 0 {
        user.Status = 1 // Default status
    }
    
    _, err := s.ormer.Insert(user)
    return err
}

// GetByID retrieves user by ID
func (s *UserService) GetByID(id int) (*models.User, error) {
    user := &models.User{Id: id}
    err := s.ormer.Read(user)
    if err == orm.ErrNoRows {
        return nil, errors.New("user not found")
    }
    return user, err
}

// GetByEmail retrieves user by email
func (s *UserService) GetByEmail(email string) (*models.User, error) {
    user := &models.User{Email: email}
    err := s.ormer.Read(user, "Email")
    if err == orm.ErrNoRows {
        return nil, errors.New("user not found")
    }
    return user, err
}

// Update updates user information
func (s *UserService) Update(user *models.User) error {
    if user.Id == 0 {
        return errors.New("user ID is required")
    }
    user.UpdatedAt = time.Now()
    _, err := s.ormer.Update(user)
    return err
}

// Delete deletes a user
func (s *UserService) Delete(id int) error {
    user := &models.User{Id: id}
    _, err := s.ormer.Delete(user)
    return err
}

// List retrieves users with pagination
func (s *UserService) List(page, pageSize int) ([]*models.User, int64, error) {
    var users []*models.User
    
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.User))
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Offset(offset).Limit(pageSize).All(&users)
    return users, total, err
}

// UpdateStatus updates user status
func (s *UserService) UpdateStatus(id int, status int) error {
    user := &models.User{Id: id}
    if err := s.ormer.Read(user); err != nil {
        return err
    }
    user.Status = status
    user.UpdatedAt = time.Now()
    _, err := s.ormer.Update(user, "Status", "UpdatedAt")
    return err
}
