package services

import (
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "errors"
)

type AuthService struct {
    ormer orm.Ormer
}

func NewAuthService() *AuthService {
    return &AuthService{
        ormer: orm.NewOrm(),
    }
}

func (s *AuthService) GetUserFromToken(token string) (*models.User, error) {
    user := &models.User{AuthKey: token}
    err := s.ormer.Read(user, "AuthKey")
    if err == orm.ErrNoRows {
        return nil, errors.New("invalid token")
    }
    return user, err
}
