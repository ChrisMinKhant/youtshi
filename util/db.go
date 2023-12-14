package util

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

/*
* This is where all the database operation
* are done.
 */

var conf = mysql.Config{
	User:                 GetEvnValue("db.username"),
	Passwd:               GetEvnValue("db.password"),
	Net:                  GetEvnValue("db.net"),
	Addr:                 GetEvnValue("db.address"),
	DBName:               GetEvnValue("db.DbName"),
	AllowNativePasswords: true,
}

// Connect to the datbase.
func OpenConnection() *sql.DB {

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatalf("Error at database connection >>> %s", err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatalf("Error at database ping >>> %s", pingErr.Error())
	}

	log.Printf("Connnected to a database >>> %v", db.Stats().OpenConnections)

	return db
}

// Build create query for inserting new data to the table.
func BuildCreateQuery(tableName string, columns []string, values []any) {
	db := OpenConnection()
	var columnString string

	for i := range columns {
		if i != 0 {
			columnString = columnString + "," + columns[i]
		} else {
			columnString = columnString + columns[i]
		}
	}

	result, err := db.Query("INSERT INTO "+tableName+"("+columnString+") VALUES(?,?,?)", values...)

	if err != nil {
		log.Fatalf("There is an error at quering >>> %v", err.Error())
	}

	log.Printf("Query Result >>> %v", result)
}
