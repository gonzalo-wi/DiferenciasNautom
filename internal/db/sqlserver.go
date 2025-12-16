package db

import (
	"database/sql"
	"fmt"
	"os"
)

func NewSqlServerDB() (*sql.DB, error) {
	connString := fmt.Sprintf(

		"sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return sql.Open("sqlserver", connString)
}
