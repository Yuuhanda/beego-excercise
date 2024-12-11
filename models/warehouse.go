package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type Warehouse struct {
    IdWh      uint   `orm:"column(id_wh);auto;pk" description:"primary key"`
    WhName    string `orm:"column(wh_name);size(255);null(false)" description:"warehouse name"`
    WhAddress string `orm:"column(wh_address);size(255);null(false)" description:"warehouse address"`
    Users     []*User `orm:"reverse(many)"` // Reverse relationship to User model
}

func init() {
    orm.RegisterModel(new(Warehouse))
}

func (w *Warehouse) TableName() string {
    return "warehouse"
}
