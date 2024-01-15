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
	User: GetEvnValue("db.username"),
	// Passwd:               GetEvnValue("db.password"),
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
	var valueString string

	for i := range columns {
		if i != 0 {
			columnString = columnString + "," + columns[i]
			valueString = valueString + ",?"
		} else {
			columnString = columnString + columns[i]
			valueString = valueString + "?"
		}
	}

	result, err := db.Query("INSERT INTO "+tableName+"("+columnString+") VALUES("+valueString+")", values...)

	if err != nil {
		log.Fatalf("There is an error at quering >>> %v", err.Error())
	}

	log.Printf("Query Result >>> %v", result)
}

func BuildSelectQuery(tableName string, indentifier string, condition []any) *sql.Rows {
	db := OpenConnection()
	result, err := db.Query("SELECT * FROM "+tableName+" WHERE "+indentifier+"= ?", condition...)

	if err != nil {
		log.Fatalf("There is an error at quering >>> %v", err.Error())
	}

	return result
}

func BuildUpdateQuery(tableName string, columns []string, indentifier string, condition []any, values []any) {
	db := OpenConnection()
	var columnString string

	for i := range columns {
		if i != 0 {
			columnString = columnString + ", " + columns[i] + "= \"" + values[i].(string) + "\""
		} else {
			columnString = columnString + columns[i] + "=\"" + values[i].(string) + "\""
		}
	}

	log.Printf(columnString)

	// UPDATE tablename SET column=value WHERE indentifier=condition
	result, err := db.Query("UPDATE "+tableName+" SET "+columnString+" WHERE "+indentifier+"=?", condition...)

	if err != nil {
		log.Fatalf("There is an error at preparing query >>> %v", err.Error())
	}

	log.Printf("Fetched update query result >>> %v", result)

}
