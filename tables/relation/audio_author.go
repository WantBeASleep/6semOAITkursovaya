package relation

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllAuthor_Audio(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"author_audio\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate author_audio table: %w", err)
	}

	return nil
}

func fillAuthor_Audio(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAuthor_Audio(db); err != nil {
		return fmt.Errorf("cant del author_audio table: %w", err)
	}

	for _, track := range data {
		err := helper.InsertAuthor_Audio(
			db,
			track.Author.Id,
			track.Audio.Id,
		)
		if err != nil {
			return fmt.Errorf("cant insert new author_audio relation: %w", err)
		}
	}
	
	return nil
}