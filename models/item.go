package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type Item struct {
    IdItem      uint          `orm:"column(id_item);auto;pk" description:"primary key"`
    ItemName    string        `orm:"column(item_name);size(60);null(false)"`
    SKU         string        `orm:"column(SKU);size(50);null(false);unique" description:"Stock Keeping Unit Code"`
    Imagefile   string        `orm:"column(imagefile);size(255);null(false)" description:"item pic"`
    IdCategory  *int          `orm:"column(id_category);null;rel(fk)" description:"foreign key"`
    Category    *ItemCategory `orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(Item))
}

func (i *Item) TableName() string {
    return "item"
}
