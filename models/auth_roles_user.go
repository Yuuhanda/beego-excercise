package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type AuthRolesUser struct {
    UserId      *User 		`orm:"rel(fk);column(user_id)"`
    RoleId   	*AuthRoles 	`orm:"rel(fk);column(roles_code)"`
}

func init() {
    orm.RegisterModel(new(AuthRolesUser))
}