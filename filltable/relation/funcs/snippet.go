package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"
	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllSnippet(db *sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"snippet\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate snippet table: %w", err)
	}

	return nil
}

func FillSnippet(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllSnippet(db); err != nil {
		return fmt.Errorf("cant del snippet table: %w", err)
	}

	audioIDS, err := db.Query(
		"SELECT id from audio",
	)
	if err != nil {
		return fmt.Errorf("cant get audio IDS: %w", err)
	}
	defer audioIDS.Close()
	i := 0

	for audioIDS.Next() {
		if i == len(data) {
			break
		}

		if rand.Float64() > constants.PercentSnippets {
			continue
		}

		audioId := 0
		err := audioIDS.Scan(&audioId)
		if err != nil {
			return fmt.Errorf("cant convert audio id: %w", err)
		}

		_, err = helper.InsertSnippet(
			db,
			data[i].Snippet.Start,
			data[i].Snippet.End,
			audioId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new snippet: %w", err)
		}

		i++
	}

	return nil
}
