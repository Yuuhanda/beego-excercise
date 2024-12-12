package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type Item struct {
    IdItem      uint          `orm:"column(id_item);auto;pk" description:"primary key"`
    ItemName    string        `orm:"column(item_name);size(60);null(false)"`
    SKU         string        `orm:"column(SKU);size(50);null(false);unique" description:"Stock Keeping Unit Code"`
    Imagefile   string        `orm:"column(imagefile);size(255);null(false)" description:"item pic"`
    Category    *ItemCategory `orm:"rel(fk);column(id_category)" description:"foreign key - category"` // Correctly defined as a pointer to ItemCategory

}

func init() {
    orm.RegisterModel(new(Item)) // Ensure both models are registered
}

func (i *Item) TableName() string {
    return "item"
}
