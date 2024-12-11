package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type ItemCategory struct {
    IdCategory    int    `orm:"column(id_category);auto;pk" description:"primary key"`
    CategoryName  string `orm:"column(category_name);size(60);null(false)"`
    CatCode       string `orm:"column(cat_code);size(6);null(false)"`
}

func init() {
    orm.RegisterModel(new(ItemCategory))
}

func (ic *ItemCategory) TableName() string {
    return "item_category"
}
