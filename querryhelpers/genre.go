package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertGenre(
	db *sql.DB,
	appellation string,
	description string,
	metricId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"genre\" (\"appellation\", \"description\", \"metric id\") " + 
			"VALUES ($1, $2, $3) RETURNING id", 
		appellation,
		description, 
		metricId,
	).Scan(&response)
	return response, err
}
