package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertAuthor(
	db *sql.DB,
	appellation string,
	description string,
	metricId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"author\" (\"appellation\", \"description\", \"metric id\") " + 
			"VALUES ($1, $2, $3) RETURNING id", 
		appellation,
		description, 
		metricId,
	).Scan(&response)
	return response, err
}

func InsertAuthor_Album(
	db *sql.DB,
	authorId int,
	albumId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"author_album\" (\"author id\", \"album id\") " + 
			"VALUES ($1, $2)", 
		authorId,
		albumId,
	)
	return err
}

func InsertAuthor_Audio(
	db *sql.DB,
	authorId int,
	audioId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"author_audio\" (\"author id\", \"audio id\") " + 
			"VALUES ($1, $2)", 
		authorId,
		audioId,
	)
	return err
}

