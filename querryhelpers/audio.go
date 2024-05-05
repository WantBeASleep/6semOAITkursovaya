package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertAudio(
	db *sql.DB,
	appellation string,
	lyric string,
	release string,
	metricId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"audio\" (\"appellation\", \"lyric\", \"release data\", \"metric id\") " + 
			"VALUES ($1, $2, $3, $4) RETURNING id", 
		appellation,
		lyric, 
		release,
		metricId,
	).Scan(&response)
	return response, err
}

func InsertAudio_Genre(
	db *sql.DB,
	audioId int,
	genreId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"audio_genre\" (\"audio id\", \"genre id\") " + 
			"VALUES ($1, $2)", 
		audioId,
		genreId,
	)
	return err
}