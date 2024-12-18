package services

import (
    "errors"
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "golang.org/x/crypto/bcrypt"
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
    
    // Add OrderBy to ensure consistent ordering
    _, err = qs.OrderBy("id").Offset(offset).Limit(pageSize).All(&users)
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
// Add to existing UserService



func (s *UserService) Authenticate(email, password string) (*models.User, error) {
    user, err := s.GetByEmail(email)
    if err != nil {
        return nil, err
    }
    
    if !verifyPassword(password, user.PasswordHash) {
        return nil, errors.New("invalid credentials")
    }
    
    return user, nil
}

func (s *UserService) GenerateAuthToken(user *models.User) error {
    authKey := generateAuthKey()
    user.AuthKey = authKey
    user.UpdatedAt = time.Now()
    
    _, err := s.ormer.Update(user, "AuthKey", "UpdatedAt")
    return err
}

func verifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func generateAuthKey() string {
    // Generate a random 32-byte key
    key := make([]byte, 32)
    rand.Read(key)
    
    // Convert to base64 string
    authKey := base64.StdEncoding.EncodeToString(key)
    
    // Add timestamp for uniqueness
    timestamp := time.Now().Format("20060102150405")
    return fmt.Sprintf("%s_%s", timestamp, authKey)
}

func (s *UserService) GetByAuthKey(authKey string) (*models.User, error) {
    user := &models.User{AuthKey: authKey}
    err := s.ormer.Read(user, "AuthKey")
    if err == orm.ErrNoRows {
        return nil, errors.New("user not found")
    }
    return user, err
}
