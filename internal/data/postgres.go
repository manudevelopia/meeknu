package data

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func getDataBaseConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", uri)
}
