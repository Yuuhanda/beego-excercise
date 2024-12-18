package models

import (
    "github.com/beego/beego/v2/client/orm"
)

type UserVisitLog struct {
    Id        int    `orm:"column(id);auto;pk"`
    Token     string `orm:"column(token);size(255)"`
    IP        string `orm:"column(ip);size(15)"`
    Language  string `orm:"column(language);size(2)"`
    UserAgent string `orm:"column(user_agent);size(255)"`
    UserId    *User  `orm:"column(user_id);rel(fk);null"`
    VisitTime int    `orm:"column(visit_time)"`
    Browser   string `orm:"column(browser);size(30);null"`
    OS        string `orm:"column(os);size(20);null"`
}

func init() {
    orm.RegisterModel(new(UserVisitLog))
}

func (u *UserVisitLog) TableName() string {
    return "user_visit_log"
}
