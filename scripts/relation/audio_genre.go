package relation

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllAudio_Genre(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"audio_genre\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate audio_genre table: %w", err)
	}

	return nil
}

func fillAudio_Genre(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAudio_Genre(db); err != nil {
		return fmt.Errorf("cant del audio_genre table: %w", err)
	}

	for _, track := range data {
		err := helper.InsertAudio_Genre(
			db,
			track.Audio.Id,
			track.Genre.Id,
		)
		if err != nil {
			return fmt.Errorf("cant insert new audio_genre relation: %w", err)
		}
	}
	
	return nil
}