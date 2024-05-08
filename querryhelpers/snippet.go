package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertSnippet(
	db *sql.DB,
	start int,
	end int,
	audioId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"snippet\" (\"start\", \"end\", \"audio id\") "+
			"VALUES ($1, $2, $3) RETURNING id",
		start,
		end,
		audioId,
	).Scan(&response)
	return response, err
}
