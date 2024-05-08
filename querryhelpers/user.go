package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// authorId OPTIONAL set -1 if user not author
func InsertUser(
	db *sql.DB,
	login string,
	email string,
	password string,
	role int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"user\" (\"login\", \"email\", \"password\", \"role\") "+
			"VALUES ($1, $2, $3, $4) RETURNING id",
		login,
		email,
		password,
		role,
	).Scan(&response)
	return response, err

}

func InsertUser_Album(
	db *sql.DB,
	userId int,
	albumId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"user_album\" (\"user id\", \"album id\") "+
			"VALUES ($1, $2)",
		userId,
		albumId,
	)
	return err
}

func InsertUser_Audio(
	db *sql.DB,
	userId int,
	audioId int,
) error {
	_, err := db.Exec(
		"INSERT INTO \"user_audio\" (\"user id\", \"audio id\") "+
			"VALUES ($1, $2)",
		userId,
		audioId,
	)
	return err
}
