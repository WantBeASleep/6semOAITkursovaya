package tables

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"
	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllSnippet(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"snippet\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate snippet table: %w", err)
	}

	return nil
}

func fillSnippet(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllSnippet(db); err != nil {
		return fmt.Errorf("cant del snippet table: %w", err)
	}

	for _, track := range data {
		if rand.Float64() > constants.PercentSnippets {
			continue
		}
		_, err := helper.InsertSnippet(
			db,
			track.Snippet.Start,
			track.Snippet.End,
			track.Audio.Id,
		)
		if err != nil {
			return fmt.Errorf("cant insert new snippet: %w", err)
		}
	}
	
	return nil
}