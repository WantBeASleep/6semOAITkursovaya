package tables

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

	"kra/data"
	"kra/tables/relation"
)

func Manager(db *sql.DB, datasetPath string, truncatePath string) error {
	fmt.Println("🤙")
	
	data, err := data.ReadCSVFile(datasetPath)
	if err != nil {
		return fmt.Errorf("cant read csv: %w", err)
	}

	// main
	if err = fillAudio(data, db); err != nil {
		return fmt.Errorf("cant fill audio: %w", err)
	}

	if err = fillGenre(data, db); err != nil {
		return fmt.Errorf("cant fill genre: %w", err)
	}

	if err = fillAlbum(db); err != nil {
		return fmt.Errorf("cant fill album: %w", err)
	}

	if err = fillAuthor(data, db); err != nil {
		return fmt.Errorf("cant fill author: %w", err)
	}

	if err = fillUser(db); err != nil {
		return fmt.Errorf("cant fill user: %w", err)
	}

	// 1:1
	if err = fillSnippet(data, db); err != nil {
		return fmt.Errorf("cant fill snippet: %w", err)
	}

	if err = fillExternal(db); err != nil {
		return fmt.Errorf("cant fill external: %w", err)
	}

	// M:N
	if err = relation.Manager(data, db); err != nil {
		return fmt.Errorf("cant fill relation: %w", err)
	}

	return nil
}