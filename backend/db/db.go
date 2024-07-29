package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabaseConnection() *sql.DB {

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DB")

	connection_string := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, db_name)

	connection, err := sql.Open("mysql", connection_string)
	connection.SetConnMaxLifetime(time.Minute * 3)

	if err != nil {
		panic("Could not connect to database" + err.Error())
	}

	return connection
}

func CloseDatabaseConnection(db *sql.DB) {
	db.Close()
}
