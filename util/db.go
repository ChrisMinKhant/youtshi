package util

import (
	"database/sql"
	"log"
	"v1/model"

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

var ErrorReponse = model.NewError()

// Connect to the datbase.
func OpenConnection() *sql.DB {

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Panicf("Error connecting database >>> %s", err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Panicf("Error pinging database >>> %s", pingErr.Error())
	}

	log.Printf("Connnected to a database >>> %v", db.Stats().OpenConnections)

	return db
}

// Build create query for inserting new data to the table.
func BuildCreateQuery(tableName string, columns []string, values []any) {
	db := OpenConnection()

	defer db.Close()

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
		log.Panicf("Error building create query >>> %v", err.Error())
	}

	log.Printf("Fetched query resul >>> %v", result)
}

// Build select query for retriving data from database
func BuildSelectQuery(tableName string, indentifier string, condition []any) (*sql.Rows, *model.Error) {
	db := OpenConnection()

	defer db.Close()

	defer func() (*sql.Rows, *model.Error) {
		if recoveryStatus := recover(); recoveryStatus != nil {
			log.Printf("System was recovered from error >>> %v", recoveryStatus)
			return nil, ErrorReponse.Set(model.I500, 500, recoveryStatus.(string))
		}

		log.Printf("System cannot be recovered.")
		return nil, ErrorReponse.Set(model.I500, 500, "Unknown error occured.")
	}()

	result, err := db.Query("SELECT * FRO "+tableName+" WHERE "+indentifier+"= ?", condition...)

	if err != nil {
		ErrorReponse.Set(model.I500, 500, "Unknown error occured.")
		log.Panicf("Error building select query >>> %v", err.Error())

		return nil, ErrorReponse
	}

	return result, nil
}

// Build update query for updating data to the database
func BuildUpdateQuery(tableName string, columns []string, indentifier string, condition []any, values []any) {
	db := OpenConnection()

	defer db.Close()

	var columnString string

	for i := range columns {
		if i != 0 {
			columnString = columnString + ", " + columns[i] + "= \"" + values[i].(string) + "\""
		} else {
			columnString = columnString + columns[i] + "=\"" + values[i].(string) + "\""
		}
	}

	result, err := db.Query("UPDATE "+tableName+" SET "+columnString+" WHERE "+indentifier+"=?", condition...)

	if err != nil {
		log.Fatalf("Error building update query >>> %v", err.Error())
	}

	log.Printf("Fetched query result >>> %v", result)

}
