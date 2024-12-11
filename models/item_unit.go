package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type ItemUnit struct {
    IdUnit       uint             `orm:"column(id_unit);auto;pk" description:"primary key"`
    IdItem       uint             `orm:"column(id_item);null(false);rel(fk)" description:"foreign key - item"`
    Status       uint8            `orm:"column(status);null(false);default(0);rel(fk)" description:"1 = available 2 = in-use 3 = in-repair 4 = lost or destroyed"`
    IdWh         *uint            `orm:"column(id_wh);null;rel(fk)" description:"where is the unit stored"`
    Comment      string           `orm:"column(comment);size(60);null" description:"additional info on unit"`
    SerialNumber string           `orm:"column(serial_number);size(60);null(false);unique"`
    Condition    uint8            `orm:"column(condition);null(false);rel(fk)" description:"refer to condition_lookup"`
    UpdatedBy    *int             `orm:"column(updated_by);null;rel(fk)" description:"last user that interacted with the data"`
    Item         *Item            `orm:"rel(one)"`
    StatusLookup *StatusLookup    `orm:"rel(one)"`
    Warehouse    *Warehouse       `orm:"rel(one)"`
    CondLookup   *ConditionLookup `orm:"rel(one)"`
    User         *User            `orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(ItemUnit))
}

func (iu *ItemUnit) TableName() string {
    return "item_unit"
}
