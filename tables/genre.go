package tables

import (
	"database/sql"
	"fmt"

	pq "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllGenre(db * sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"metric id\" FROM \"genre\" ",
	)
	if err != nil {
		return fmt.Errorf("cant get genre && metric id: %w", err)
	}
	rows.Close()

	for rows.Next() {
		genreId, metricId := 0, 0
		if err := rows.Scan(&genreId, &metricId); err != nil {
			return fmt.Errorf("cant convert genre && metric id: %w", err)
		}

		_, err := db.Exec(
			"DELETE FROM \"genre\" " + 
			"WHERE id = $1",
			genreId,
		)
		if err != nil {
			return fmt.Errorf("cant del genre by id: %w", err)
		}

		_, err = db.Exec(
			"DELETE FROM \"metric\" " + 
			"WHERE id = $1",
			metricId,
		)
		if err != nil {
			return fmt.Errorf("cant del metric by id: %w", err)
		}
	}

	_, err = db.Exec(
		"TRUNCATE TABLE \"genre\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate genre table: %w", err)
	}

	return nil
}

func fillGenre(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllGenre(db); err != nil {
		return fmt.Errorf("cant del genre table: %w", err)
	}

	for _, track := range data {
		metricId, err := createMetric(db)
		if err != nil {
			return fmt.Errorf("cant insert new genre metric: %w", err)
		}
		genreId, err := helper.InsertGenre(
			db,
			track.Genre.Appellation,
			track.Genre.Description,
			metricId,
		)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); !ok || pqErr.Code != "23505" {
				return fmt.Errorf("cant insert new genre: %w", err)
			}
			err = db.QueryRow(
				"SELECT id FROM genre WHERE appellation = $1",
				track.Genre.Appellation,
			).Scan(&genreId)
			if err != nil {
				return fmt.Errorf("cant get non uniq genre id: %w", err)
			}
		}
		track.Genre.Id = genreId
	}
	
	return nil
}