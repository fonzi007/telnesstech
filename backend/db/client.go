package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDbClient() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:changeme@tcp(localhost:3306)/testDB?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %w", err)
	}
	return db, err
}
