package main

import (
	"database/sql"
	"fmt"
	"kra/scripts"

	_ "github.com/lib/pq"
)

// вхуярить сюда env
func openConnection() (*sql.DB, error) {
	connStr := "user=postgres password=1234 dbname=6sem sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cant sql open: %w", err)
	}
	return db, nil
}

func main() {
	db, err := openConnection()
	if err != nil {
		panic(fmt.Errorf("cant connect to postgres!: %w", err))
	}
	defer db.Close()

	if err := scripts.Manager(db); err != nil {
		panic(fmt.Errorf("cant manipulate with tables: %w", err))
	}
}