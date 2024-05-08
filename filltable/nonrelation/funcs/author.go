package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	funcshelpers "kra/filltable/helpers"
	helper "kra/querryhelpers"
)

func deleteAllAuthor(db *sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"metric id\" FROM \"author\" ",
	)
	if err != nil {
		return fmt.Errorf("cant get author && metric id: %w", err)
	}
	rows.Close()

	for rows.Next() {
		authorId, metricId := 0, 0
		if err := rows.Scan(&authorId, &metricId); err != nil {
			return fmt.Errorf("cant convert author && metric id: %w", err)
		}

		_, err := db.Exec(
			"DELETE FROM \"author\" "+
				"WHERE id = $1",
			authorId,
		)
		if err != nil {
			return fmt.Errorf("cant del author by id: %w", err)
		}

		_, err = db.Exec(
			"DELETE FROM \"metric\" "+
				"WHERE id = $1",
			metricId,
		)
		if err != nil {
			return fmt.Errorf("cant del metric by id: %w", err)
		}
	}

	_, err = db.Exec(
		"TRUNCATE TABLE \"author\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate author table: %w", err)
	}

	return nil
}

func FillAuthor(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAuthor(db); err != nil {
		return fmt.Errorf("cant del author table: %w", err)
	}

	seenAuthors := map[string]bool{}

	for _, track := range data {
		if seenAuthors[track.Author.Appellation] {
			continue
		}
		seenAuthors[track.Author.Appellation] = true

		metricId, err := funcshelpers.CreateMetric(db)
		if err != nil {
			return fmt.Errorf("cant insert new author metric: %w", err)
		}
		_, err = helper.InsertAuthor(
			db,
			track.Author.Appellation,
			track.Author.Description,
			metricId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new author: %w", err)
		}
	}

	return nil
}
