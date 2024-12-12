package database

import (
    "fmt"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    configFile := "conf/app.conf"
    err := web.LoadAppConfig("ini", configFile)
    if err != nil {
        panic("Failed to load config file: " + configFile)
    }
    
    dbUser, err := web.AppConfig.String("mysqluser")
    if err != nil {
        panic("mysqluser not found in config")
    }
    
    dbPass, err := web.AppConfig.String("mysqlpass")
    if err != nil {
        panic("mysqlpass not found in config") 
    }

    dbName, err := web.AppConfig.String("mysqldb")
    if err != nil {
        panic("mysqldb not found in config")
    }

    dbHost, err := web.AppConfig.String("mysqlhost") 
    if err != nil {
        panic("mysqlhost not found in config")
    }

    dbPort, err := web.AppConfig.String("mysqlport")
    if err != nil {
        panic("mysqlport not found in config")
    }

    dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", 
        dbUser, dbPass, dbHost, dbPort, dbName)

    orm.RegisterDriver("mysql", orm.DRMySQL)
    if err := orm.RegisterDataBase("default", "mysql", dataSource); err != nil {
        panic(err)
    }
}