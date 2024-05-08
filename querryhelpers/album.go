package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertAlbum(
	db *sql.DB,
	appellation string,
	release string,
	metricId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"album\" (\"appellation\", \"release data\", \"metric id\") "+
			"VALUES ($1, $2, $3) RETURNING id",
		appellation,
		release,
		metricId,
	).Scan(&response)
	return response, err
}

func InsertAlbum_Genre(
	db *sql.DB,
	albumId int,
	genreId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"album_genre\" (\"album id\", \"genre id\") "+
			"VALUES ($1, $2)",
		albumId,
		genreId,
	)
	return err
}

func InsertAlbum_Audio(
	db *sql.DB,
	albumId int,
	audioId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"album_audio\" (\"album id\", \"audio id\") "+
			"VALUES ($1, $2)",
		albumId,
		audioId,
	)
	return err
}
