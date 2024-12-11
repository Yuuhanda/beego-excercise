package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type RepairLog struct {
    IdRepair  int       `orm:"column(id_repair);auto;pk" description:"primary key"`
    IdUnit    int       `orm:"column(id_unit);null(false)" description:"foreign key - item_unit"`
    Comment   string    `orm:"column(comment);size(120);null(false)" description:"comments content"`
    RepType   int8      `orm:"column(rep_type);null(false);rel(fk)" description:"type of repair log"`
    Datetime  time.Time `orm:"column(datetime);type(datetime);precision(2);null(false)" description:"date & time of which the log is written"`
    TypeLookup *RepTypeLookup `orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(RepairLog))
}

func (r *RepairLog) TableName() string {
    return "repair_log"
}
