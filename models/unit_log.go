package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type UnitLog struct {
    IdLog     int       `orm:"column(id_log);auto;pk" description:"primary key"`
    IdUnit    int       `orm:"column(id_unit);null(false)" description:"foreign key - item_unit"`
    Content   string    `orm:"column(content);size(255);null(false)" description:"content of the log, text"`
    UpdateAt  time.Time `orm:"column(update_at);type(datetime);precision(6);null(false)" description:"date time of which the log is written"`
}

func init() {
    orm.RegisterModel(new(UnitLog))
}

func (ul *UnitLog) TableName() string {
    return "unit_log"
}
