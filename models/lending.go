package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type Lending struct {
    IdLending  uint             `orm:"column(id_lending);auto;pk" description:"primary key"`
    IdUnit     *ItemUnit        `orm:"rel(fk);column(id_unit)" description:"foreign key - item_unit"`
    IdUser     *User            `orm:"rel(fk);column(user_id)" description:"foreign key - user"`
    IdEmployee *Employee        `orm:"rel(fk);column(id_employee)" description:"foreign key - employee"`
    Type       *LendingTypeLookup `orm:"rel(fk);column(type)" description:"1 = in-use 2 = returned"`
    Date       time.Time        `orm:"column(date);type(date);null(false)" description:"date of data"`
    PicLoan    string           `orm:"column(pic_loan);size(255);null(false)" description:"pic before being used"`
    PicReturn  string           `orm:"column(pic_return);size(255);null(false)" description:"pic on return"`
}



func init() {
    orm.RegisterModel(new(Lending))
}

func (l *Lending) TableName() string {
    return "lending"
}
