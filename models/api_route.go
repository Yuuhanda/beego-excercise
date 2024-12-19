package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type ApiRoute struct {
    Id          int       `orm:"column(id);auto;pk"`
    Path        string    `orm:"column(path);size(255)"`
    Method      string    `orm:"column(method);size(10)"`
    Controller  string    `orm:"column(controller);size(100)"`
    Action      string    `orm:"column(action);size(100)"`
    Description string    `orm:"column(description);size(255);null"`
    CreatedAt   time.Time `orm:"column(created_at);type(datetime)"`
    UpdatedAt   time.Time `orm:"column(updated_at);type(datetime)"`
}

func init() {
    orm.RegisterModel(new(ApiRoute))
}
