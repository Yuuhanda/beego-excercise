package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type ItemUnit struct {
    IdUnit       uint             `orm:"column(id_unit);auto;pk"`

    Comment      string           `orm:"column(comment);size(60);null"`
    SerialNumber string           `orm:"column(serial_number);size(60);null(false);unique"`


    // Relationships with proper column specifications
    Item         *Item            `orm:"rel(fk);column(id_item)"`
    StatusLookup *StatusLookup    `orm:"rel(fk);column(status)"`
    Warehouse    *Warehouse       `orm:"rel(fk);column(id_wh)"`
    CondLookup   *ConditionLookup `orm:"rel(fk);column(condition)"`
    User         *User            `orm:"rel(fk);column(updated_by)"`

}


func init() {
    orm.RegisterModel(new(ItemUnit))
}

func (iu *ItemUnit) TableName() string {
    return "item_unit"
}
