package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type RepairLog struct {
    IdRepair   int            `orm:"column(id_repair);auto;pk" description:"primary key"`
    IdUnit     *uint          `orm:"column(id_unit);null(false)" description:"foreign key - item_unit"`
    Comment    string         `orm:"column(comment);size(120);null(false)" description:"comments content"`
    RepType    *RepTypeLookup `orm:"rel(fk);column(rep_type)" description:"type of repair log"`
    Datetime   time.Time      `orm:"column(datetime);type(datetime);precision(2);null(false)" description:"date & time of which the log is written"`
    Unit       *ItemUnit      `orm:"rel(one)"`
}


func init() {
    orm.RegisterModel(new(RepairLog))
}

func (r *RepairLog) TableName() string {
    return "repair_log"
}
