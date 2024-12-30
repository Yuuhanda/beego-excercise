package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type AuthItem struct {
    Id      int         `orm:"column(id);auto"`
    Role    string      `orm:"column(role)" description:"references auth_roles.code"`
    Path    string      `orm:"column(path)" description:"references api_route.path"`
    Method      string    `orm:"column(method);size(10)"`
}

func init() {
    orm.RegisterModel(new(AuthItem))
}

type PathMethod struct {
    Path   string `json:"path"`
    Method string `json:"method"`
}
