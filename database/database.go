package database

import (
	"database/sql"
	"fmt"

	"echo-rest/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	conf := config.LoadConfiguration()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	fmt.Println(connectionString)
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		println(err.Error())
		panic("connectionString Error...")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
