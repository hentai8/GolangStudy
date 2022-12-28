package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

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

func create() *sqlx.DB {

	databaseString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", "root", "zhujia1007", "127.0.0.1", "3306", "dxfufu", "charset=utf8mb4&parseTime=True&loc=Local")
	database, err := sqlx.Open("mysql", databaseString)

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}
	return database
}
