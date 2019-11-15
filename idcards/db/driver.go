package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectSQL ...
func ConnectSQL() (*sql.DB, error) {

	dbName := os.Getenv("MYSQL_DATABASE")
	dbRoot := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	dbSource := fmt.Sprintf(
		dbRoot+":%s@tcp(%s:%s)/%s?charset=utf8",
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	return d, err
}

//test_tb
//name

