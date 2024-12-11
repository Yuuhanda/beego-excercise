package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type Employee struct {
    IdEmployee uint   `orm:"column(id_employee);auto;pk" description:"primary key"`
    EmpName    string `orm:"column(emp_name);size(80);null(false)" description:"employee name"`
    Phone      string `orm:"column(phone);size(20);null(false)"`
    Email      string `orm:"column(email);size(60);null(false);unique"`
    Address    string `orm:"column(address);size(255);null(false)"`
}

func init() {
    orm.RegisterModel(new(Employee))
}

func (e *Employee) TableName() string {
    return "employee"
}
