package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type AuthRoles struct {
    Code        string    `orm:"column(code);size(32);pk"`
    Name        string    `orm:"column(name);size(100)"`
    Description string    `orm:"column(description);size(255);null"`
    CreatedAt   time.Time `orm:"column(created_at);type(datetime)"`
    UpdatedAt   time.Time `orm:"column(updated_at);type(datetime)"`
}

func init() {
    orm.RegisterModel(new(AuthRoles))
}