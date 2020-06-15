package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:Aa123456@/dbshopee")

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
