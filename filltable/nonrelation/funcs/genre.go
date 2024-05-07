package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
	funcshelpers "kra/filltable/helpers"
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

func FillGenre(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllGenre(db); err != nil {
		return fmt.Errorf("cant del genre table: %w", err)
	}

	seenGenres := map[string]bool{}

	for _, track := range data {
		if seenGenres[track.Genre.Appellation] {
			continue
		}
		seenGenres[track.Genre.Appellation] = true

		metricId, err := funcshelpers.CreateMetric(db)
		if err != nil {
			return fmt.Errorf("cant insert new genre metric: %w", err)
		}
		_, err = helper.InsertGenre(
			db,
			track.Genre.Appellation,
			track.Genre.Description,
			metricId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new genre: %w", err)
		}
	}

	return nil
}