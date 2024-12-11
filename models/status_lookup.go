package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type StatusLookup struct {
    IdStatus    uint8  `orm:"column(id_status);auto;pk" description:"primary key"`
    StatusName  string `orm:"column(status_name);size(30);null(false)" description:"status of the unit"`
}

func init() {
    orm.RegisterModel(new(StatusLookup))
}

func (s *StatusLookup) TableName() string {
    return "status_lookup"
}
