package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertExternal(
	db *sql.DB,
	link string,
	rtype string,
	authorId int,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"external resource\" (\"link\", \"type\", \"author id\") " + 
			"VALUES ($1, $2, $3) RETURNING id", 
		link,
		rtype, 
		authorId,
	).Scan(&response)
	return response, err
}

