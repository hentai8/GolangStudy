package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/configor"
    "github.com/jmoiron/sqlx"
    "sync"
    "time"
)

type Config struct {
    Mysql Mysql   `json:"mysql"`
    Coins []Coins `json:"coins"`
}

type Coins struct {
    Name       string `json:"name"`
    Port       string `json:"port"`
    TargetPool struct {
        Name      string `json:"name"`
        Host      string `json:"host"`
        Port      string `json:"port"`
        Worker    string `json:"worker"`
        Enable    bool   `json:"enable"`
        ReportUrl string `json:"report_url"`
    } `json:"target_pool"`
    ListeningPools []struct {
        Name   string `json:"name"`
        Host   string `json:"host"`
        Port   string `json:"port"`
        Worker string `json:"worker"`
        Enable bool   `json:"enable"`
    } `json:"listening_pools"`
}

type Mysql struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Network  string `json:"network"`
    Server   string `json:"server"`
    Port     int    `json:"port"`
    Database string `json:"database"`
}


var sqlxDB *sqlx.DB
var once sync.Once

func GetInstance() *sqlx.DB {
    once.Do(func() {
        sqlxDB = create()
    })

    return sqlxDB
}

func Close() {
    if sqlxDB != nil {
        sqlxDB.Close()
    }
}

type Conf struct {
    Mysql struct{
        MysqlIp string `yaml:"MysqlIp"`
        MySqlPort int `yaml:"MySqlPort"`
        MySqlUser string `yaml:"MySqlUser"`
        MySqlPwd string `yaml:"MySqlPwd"`
        MySqlTab string `yaml:"MySqlTab"`
    } `yaml:"Mysql"`
}

var cfg Config

func create() *sqlx.DB {
    err := configor.Load(&cfg, "config.json")
    if err != nil {
        fmt.Println("read config err=", err)
        return nil
    }
    dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
        cfg.Mysql.Username,
        cfg.Mysql.Password,
        cfg.Mysql.Network,
        cfg.Mysql.Server,
        cfg.Mysql.Port,
        cfg.Mysql.Database)
    database, err := sqlx.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println("open mysql failed,", err)
        return nil
    }

    return database
}

func Record(height int64, url string) error {
    db := GetInstance()
    now := time.Now().Unix()
    sql := `insert into ltc (height, url, start_time) values (?,?,?)`
    _, err := db.Exec(sql, height, url, now)
    return err
}


func main() {
    err := Record(111, "127.0.0.1")
    if err != nil {
        fmt.Println("record to mysql err=", err)
    }
}
