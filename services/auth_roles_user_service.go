package services

import (
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
)

type AuthRolesUserService struct {
    ormer orm.Ormer
}

func NewAuthRolesUserService() *AuthRolesUserService {
    return &AuthRolesUserService{
        ormer: orm.NewOrm(),
    }
}

func (s *AuthRolesUserService) Create(roleUser *models.AuthRolesUser) error {
    _, err := s.ormer.Insert(roleUser)
    return err
}

func (s *AuthRolesUserService) GetRolesByUserId(userId int) ([]*models.AuthRoles, error) {
    var roles []*models.AuthRoles
    _, err := s.ormer.Raw("SELECT r.* FROM auth_roles r "+
        "JOIN auth_roles_user ru ON r.id = ru.roles_id "+
        "WHERE ru.user_id = ?", userId).QueryRows(&roles)
    return roles, err
}

func (s *AuthRolesUserService) GetUsersByRoleId(roleId int) ([]*models.User, error) {
    var users []*models.User
    _, err := s.ormer.Raw("SELECT u.* FROM user u "+
        "JOIN auth_roles_user ru ON u.id = ru.user_id "+
        "WHERE ru.roles_id = ?", roleId).QueryRows(&users)
    return users, err
}

func (s *AuthRolesUserService) Delete(userId int, roleId int) error {
    _, err := s.ormer.Raw("DELETE FROM auth_roles_user WHERE user_id = ? AND roles_id = ?", 
        userId, roleId).Exec()
    return err
}
