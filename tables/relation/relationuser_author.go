package relation

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
)

func relationUser_Author(data []data.TrackInfo, db *sql.DB) error {
	
	for _, track := range data {
		seenUsers := map[int]bool{-1: true}

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

		_, err := db.Exec(
			"UPDATE \"user\" " + 
			"SET " +
			"\"author id\" = $1, " + 
			"role = 1",
			track.Author.Id,
		)
		if err != nil {
			return fmt.Errorf("cant update authorID and role on user: %w", err)
		}
	}

	return nil
}