package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	cfg := mysql.Config{
		User:                 "physician",
		Passwd:               "health",
		Net:                  "tcp",
		Addr:                 "mysql:3306",
		DBName:               "notable_health",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		panic(pingErr)
	}

	fmt.Println("Connected to mysql...")
}

func GetDB() *sql.DB {
	return db
}
