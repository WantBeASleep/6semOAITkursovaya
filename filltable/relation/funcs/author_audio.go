package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllAuthor_Audio(db *sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"author_audio\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate author_audio table: %w", err)
	}

	return nil
}

func FillAuthor_Audio(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAuthor_Audio(db); err != nil {
		return fmt.Errorf("cant del author_audio table: %w", err)
	}

	for _, track := range data {

		authorID := 0
		err := db.QueryRow(
			"SELECT id FROM author "+
				"WHERE appellation = $1",
			track.Author.Appellation,
		).Scan(&authorID)
		if err != nil {
			return fmt.Errorf("cant get authorID: %w", err)
		}

		audioId := 0
		err = db.QueryRow(
			"SELECT id FROM audio "+
				"WHERE lyric = $1",
			track.Audio.Lyric,
		).Scan(&audioId)
		if err != nil {
			return fmt.Errorf("cant get audioID: %w", err)
		}

		err = helper.InsertAuthor_Audio(
			db,
			authorID,
			audioId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new author_audio relation: %w", err)
		}
	}

	return nil
}
