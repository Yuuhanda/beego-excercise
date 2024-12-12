package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type LendingTypeLookup struct {
    IdType    uint  `orm:"column(id_type);auto;pk" description:"primary key"`
    TypeName  string `orm:"column(type_name);size(30);null(false)" description:"type of lending data be it lending out or returned"`
}

func init() {
    orm.RegisterModel(new(LendingTypeLookup))
}

func (l *LendingTypeLookup) TableName() string {
    return "lending_type_lookup"
}
