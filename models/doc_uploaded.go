package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type DocUploaded struct {
    IdDoc     int       `orm:"column(id_doc);auto;pk" description:"primary key"`
    FileName  string    `orm:"column(file_name);size(255);null(false)"`
    Datetime  time.Time `orm:"column(datetime);type(datetime);precision(2);null(false)"`
    UserId    *User `orm:"rel(fk);column(user_id)"`
}

func init() {
    orm.RegisterModel(new(DocUploaded))
}

func (d *DocUploaded) TableName() string {
    return "doc_uploaded"
}
