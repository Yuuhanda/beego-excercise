package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type RepTypeLookup struct {
    IdRepT   uint   `orm:"column(id_rep_t);auto;pk" description:"primary key"`
    RepType  string `orm:"column(rep_type);size(255);null(false)" description:"type of log be it repair initiated or closed"`
}

func init() {
    orm.RegisterModel(new(RepTypeLookup))
}

func (r *RepTypeLookup) TableName() string {
    return "rep_type_lookup"
}
