package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

var (
	DB_HOST     = os.Getenv("MSSQL_DB_HOST")
	DB_PORT     = os.Getenv("MSSQL_DB_PORT")
	DB_USER     = os.Getenv("MSSQL_DB_USER")
	DB_PASSWORD = os.Getenv("MSSQL_DB_PASSWORD")
	DB_DATABASE = os.Getenv("MSSQL_DB_DATABASE")
)

var lock = &sync.Mutex{}

var DatabaseInstance *sql.DB

func GetDatabaseInstance() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable;", DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_DATABASE)

	if DatabaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if DatabaseInstance == nil {
			db, err := sql.Open("sqlserver", connString)

			if err != nil {
				log.Fatal("Error creating connection pool: ", err.Error())
			}

			ctx := context.Background()

			err = db.PingContext(ctx)

			if err != nil {
				log.Fatal(err.Error())
			}
			DatabaseInstance = db
		}
	}

	return DatabaseInstance
}
