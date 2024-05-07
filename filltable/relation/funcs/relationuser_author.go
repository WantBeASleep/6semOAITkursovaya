package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
)

func RelationUser_Author(data []data.TrackInfo, db *sql.DB) error {

	authorIDs, err := db.Query(
		"SELECT id FROM author",
	)
	if err != nil {
		return fmt.Errorf("cant get authors ids: %w", err)
	}

	seenUsers := map[int]bool{-1: true}

	for authorIDs.Next() {
		authorID := 0
		err := authorIDs.Scan(&authorID)
		if err != nil {
			return fmt.Errorf("cant scan author id: %w", err)
		}

		randUserID := -1
		for seenUsers[randUserID] {
			err := db.QueryRow(
				"SELECT id FROM \"user\" " + 
				"ORDER BY RANDOM() " + 
				"LIMIT 1",
			).Scan(&randUserID)
			if err != nil {
				return fmt.Errorf("cant get random user: %w", err)
			}
		}
		seenUsers[randUserID] = true

		_, err = db.Exec(
			"UPDATE \"user\" " + 
			"SET " +
			"\"author id\" = $1, " + 
			"role = 1 " +
			"WHERE id = $2 ",
			authorID,
			randUserID,
		)
		if err != nil {
			return fmt.Errorf("cant update authorID and role on user: %w", err)
		}
	}

	return nil
}