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

func CreateTable() error {
	conn, err := ConnectSQL()

	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Query("CREATE TABLE IF NOT EXISTS Users(" +
		"id INT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"name VARCHAR(255) NOT NULL UNIQUE," +
		"email VARCHAR(255)" +
		")")
	if err != nil {
		return err
	}
	return nil
}

//test_tb
//name

