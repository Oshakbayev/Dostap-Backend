package database

import (
	"database/sql"
	"log"
)

func CreateDB(DBDriver string) *sql.DB {
	DBconnectionString := "postgres://owwkdlwj:UqnYqPkMlDPUHBntLlFYpIeLjaXZkCxR@abul.db.elephantsql.com/owwkdlwj"
	db, err := sql.Open(DBDriver, DBconnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
