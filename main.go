package main

import (
    "myproject/database"
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/server/web"
    "os"
    "os/signal"
    "syscall"
    "time"
    _ "myproject/routers"
    _ "myproject/database"
    _ "myproject/services"
    _ "myproject/controllers"
    _ "myproject/models"
)

func main() {
    database.GetInstance() // This is the correct function name
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        logs.Info("Shutting down server...")
        os.Exit(0)
    }()
    
    web.Run()
}

func init() {
    os.Setenv("TZ", "Asia/Jakarta")
    time.Local = time.FixedZone("WIB", 7*60*60) // WIB = UTC+7
    database.InitDatabase()
    //logs.SetLogger(logs.AdapterConsole, "")
    //logs.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}
