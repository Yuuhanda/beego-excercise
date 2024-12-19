package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type AuthItem struct {
    Id      int         `orm:"column(id);auto"`
    Role    string      `orm:"column(role)" description:"references auth_roles.code"`
    Path    string      `orm:"column(path)" description:"references api_route.path"`
}

func init() {
    orm.RegisterModel(new(AuthItem))
}
