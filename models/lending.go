package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type Lending struct {
    IdLending  uint              `orm:"column(id_lending);auto;pk" description:"primary key"`
    IdUnit     uint              `orm:"column(id_unit);null(false);rel(fk)" description:"foreign key - item_unit"`
    UserId     int               `orm:"column(user_id);null(false);rel(fk)" description:"foreign key - user 'id'"`
    IdEmployee uint              `orm:"column(id_employee);null(false);rel(fk)" description:"foreign key - employee"`
    Type       uint8             `orm:"column(type);null(false);rel(fk)" description:"1 = in-use 2 = returned"`
    Date       time.Time         `orm:"column(date);type(date);null(false)" description:"date of data"`
    PicLoan    string           `orm:"column(pic_loan);size(255);null(false)" description:"pic before being used"`
    PicReturn  string           `orm:"column(pic_return);size(255);null(false)" description:"pic on return"`
    ItemUnit   *ItemUnit        `orm:"rel(one)"`
    User       *User            `orm:"rel(one)"`
    Employee   *Employee        `orm:"rel(one)"`
    LendType   *LendingTypeLookup `orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(Lending))
}

func (l *Lending) TableName() string {
    return "lending"
}
