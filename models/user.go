package models

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
)

type User struct {
    Id                int       `orm:"column(id);auto;pk"`
    Username          string    `orm:"column(username);size(255);null(false)"`
    PasswordHash      string    `orm:"column(password_hash);size(255);null(false)"`
    Status            int       `orm:"column(status);default(1)"`
    Superadmin        int8      `orm:"column(superadmin);default(0)"`
    CreatedAt         time.Time `orm:"column(created_at);type(datetime);null(false)"`
    UpdatedAt         time.Time `orm:"column(updated_at);type(datetime);null(false)"`
    RegistrationIP    string    `orm:"column(registration_ip);size(15);null"`
    Email             string    `orm:"column(email);size(128);null"`
    AuthKey           string    `orm:"column(auth_key);size(255);null"`
    BindToIP          string    `orm:"column(bind_to_ip);size(255);null"`
    EmailConfirmed    int       `orm:"column(email_confirmed);null(false)"`
    ConfirmationToken string    `orm:"column(confirmation_token);size(255);null"`
    IdWh              *Warehouse `orm:"column(id_wh);null;rel(fk)" description:"assigned warehouse if this account is a warehouse admin"` // Changed to *Warehouse
    UserLang          string    `orm:"column(user_lang);size(3);null" description:"preferred ui language"`
}

func init() {
    // Register model
    orm.RegisterModel(new(User))
}

// TableName specifies the table name
func (u *User) TableName() string {
    return "user"
}
