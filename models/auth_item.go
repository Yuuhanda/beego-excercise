package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type AuthItem struct {
    Id          int   		`orm:"column(id);auto"`
	RoleId      *AuthRoles 	`orm:"rel(fk);column(role_id)"`
	RouteId     *ApiRoute  	`orm:"rel(fk);column(route_id)"`
}

func init() {
    orm.RegisterModel(new(AuthItem))
}