package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type ConditionLookup struct {
    IdCondition    uint8  `orm:"column(id_condition);auto;pk" description:"primary key"`
    ConditionName  string `orm:"column(condition_name);size(255);null(false)"`
}

func init() {
    orm.RegisterModel(new(ConditionLookup))
}

func (c *ConditionLookup) TableName() string {
    return "condition_lookup"
}
