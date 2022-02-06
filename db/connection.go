package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbURL = "root:@tcp(localhost:3306)/petshopgo"
var db *sql.DB

//Func to open the connection with the DB
func Open() {
	connect, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	db = connect
}

//Func to verify is the connection is still open
func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

//Func to close the connection with the DB
func Close() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

//func to create a table
func CreateTable(schemaTable string) {
	if _, err := Exec(schemaTable); err != nil {
		panic(err)
	} else {
		fmt.Println("Table created")
	}
}

//func to truncate a table
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("Truncate table %s", tableName)
	if _, err := Exec(sql); err != nil {
		panic(err)
	} else {
		fmt.Println("Table truncated")
	}
}

//func to execute statements to the DB
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Open()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		panic(err)
	}

	return result, err
}

// func to execute some database query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Open()
	rows, err := db.Query(query, args...)
	Close()

	if err != nil {
		panic(err)
	}

	return rows, err
}
