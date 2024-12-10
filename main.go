package main

import (
    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
    _ "github.com/go-sql-driver/mysql"
    _ "myproject/models"
    _ "myproject/routers"
	"os"
    "os/signal"
    "syscall"
	"fmt"
)
func main() {
    orm.Debug = true
    orm.RunSyncdb("default", false, true)
    
    // Setup graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        logs.Info("Shutting down server...")
        // Perform cleanup
        os.Exit(0)
    }()
    
    beego.Run()
}


func init() {
    user, _ := beego.AppConfig.String("mysqluser")
    password, _ := beego.AppConfig.String("mysqlpass")
    dbName, _ := beego.AppConfig.String("mysqldb")
    host, _ := beego.AppConfig.String("mysqlhost")
    port, _ := beego.AppConfig.String("mysqlport")

    // Register the MySQL database
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user, password, host, port, dbName)
    if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
        panic(err)
    }

    orm.SetMaxIdleConns("default", 10)
    orm.SetMaxOpenConns("default", 100)
}