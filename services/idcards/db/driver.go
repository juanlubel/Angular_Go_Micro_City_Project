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
		"surname VARCHAR(255)," +
		"slug VARCHAR(255)," +
		"email VARCHAR(255)," +
		"nCard DOUBLE NOT NULL," +
		"pass VARCHAR(255)" +
		")")


	if err != nil {
		fmt.Print("error1")
		return err
	}
	//var res string
	//err = conn.QueryRow("select id from Users").Scan(res)
	//if err != nil {
	//	fmt.Print("error2")
	//}

	//for rows.Next() {
	//	fmt.Print(res)
	//	rows.Scan(&res)
	//}
	return nil
}

//test_tb
//name

