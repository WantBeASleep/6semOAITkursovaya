package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Используем драйвер PostgreSQL
)

func TruncateAllTables(db *sql.DB, path string) error {
	query, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cant read SQL file: %v", err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("cant truncate tables: %v", err)
	}

	return nil
}
