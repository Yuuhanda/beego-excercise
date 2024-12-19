package services

import (
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "fmt"
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
    // Get the role first
    role := &models.AuthRoles{Code: roleUser.RoleId.Code}
    if err := s.ormer.Read(role); err != nil {
        return fmt.Errorf("role with Code %s not found", roleUser.RoleId.Code)
    }
    roleUser.RoleId = role

    // Get the user
    user := &models.User{Id: roleUser.UserId.Id}
    if err := s.ormer.Read(user); err != nil {
        return fmt.Errorf("user with ID %d not found", roleUser.UserId.Id)
    }
    roleUser.UserId = user

    // Create the assignment
    _, err := s.ormer.Insert(roleUser)
    return err
}



func (s *AuthRolesUserService) GetRolesByUserId(userId int) ([]*models.AuthRoles, error) {
    var roles []*models.AuthRoles
    _, err := s.ormer.Raw("SELECT ar.* FROM auth_roles ar "+
        "JOIN auth_roles_user aru ON ar.code = aru.roles_code "+
        "WHERE aru.user_id = ?", userId).QueryRows(&roles)
    return roles, err
}

func (s *AuthRolesUserService) GetUsersByRoleId(roleId string) ([]*models.User, error) {
    var users []*models.User
    _, err := s.ormer.Raw("SELECT u.* FROM user u "+
        "JOIN auth_roles_user ru ON u.id = ru.user_id "+
        "WHERE ru.role_id = ?", roleId).QueryRows(&users)
    return users, err
}


func (s *AuthRolesUserService) Delete(userId int) error {
    _, err := s.ormer.Raw("DELETE FROM auth_roles_user WHERE user_id = ?", 
        userId).Exec()
    return err
}
