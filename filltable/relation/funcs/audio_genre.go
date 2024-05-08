package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllAudio_Genre(db *sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"audio_genre\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate audio_genre table: %w", err)
	}

	return nil
}

func FillAudio_Genre(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAudio_Genre(db); err != nil {
		return fmt.Errorf("cant del audio_genre table: %w", err)
	}

	for _, track := range data {

		audioId := 0
		err := db.QueryRow(
			"SELECT id FROM audio "+
				"WHERE lyric = $1",
			track.Audio.Lyric,
		).Scan(&audioId)
		if err != nil {
			return fmt.Errorf("cant get audioID: %w", err)
		}

		genreID := 0
		err = db.QueryRow(
			"SELECT id FROM genre "+
				"WHERE appellation = $1",
			track.Genre.Appellation,
		).Scan(&genreID)
		if err != nil {
			return fmt.Errorf("cant get genreID: %w", err)
		}

		err = helper.InsertAudio_Genre(
			db,
			audioId,
			genreID,
		)
		if err != nil {
			return fmt.Errorf("cant insert new audio_genre relation: %w", err)
		}
	}

	return nil
}
