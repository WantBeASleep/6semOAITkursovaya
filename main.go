package main

import (
	"database/sql"
	"fmt"
	"kra/constants"
	"kra/tables"

	_ "github.com/lib/pq"
)

// вхуярить сюда env
func openConnection() (*sql.DB, error) {
	connStr := "user=postgres password=1234 dbname=6sem sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cant sql open: %w", err)
	}
	defer db.Close()

	return db, nil
}

func main() {
	db, err := openConnection()
	if err != nil {
		panic(fmt.Errorf("cant connect to postgres!: %w", err))
	}

	if err := tables.Manager(db, constants.DATASET_PATH, constants.TRUNCATE_PATH); err != nil {
		panic(fmt.Errorf("cant manipulate with tables: %w", err))
	}
}