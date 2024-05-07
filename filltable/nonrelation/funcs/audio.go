package funcs

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
	funcshelpers "kra/filltable/helpers"
)

func deleteAllAudio(db * sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"metric id\" FROM \"audio\" ",
	)
	if err != nil {
		return fmt.Errorf("cant get audio && metric id: %w", err)
	}
	rows.Close()

	for rows.Next() {
		audioId, metricId := 0, 0
		if err := rows.Scan(&audioId, &metricId); err != nil {
			return fmt.Errorf("cant convert audio && metric id: %w", err)
		}

		_, err := db.Exec(
			"DELETE FROM \"audio\" " + 
			"WHERE id = $1",
			audioId,
		)
		if err != nil {
			return fmt.Errorf("cant del audio by id: %w", err)
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
		"TRUNCATE TABLE \"audio\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate audio table: %w", err)
	}

	return nil
}

func FillAudio(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAudio(db); err != nil {
		return fmt.Errorf("cant del audio table: %w", err)
	}

	for _, track := range data {
		metricId, err := funcshelpers.CreateMetric(db)
		if err != nil {
			return fmt.Errorf("cant insert new audio metric: %w", err)
		}
		_, err = helper.InsertAudio(
			db,
			track.Audio.Appellation,
			track.Audio.Lyric,
			track.Audio.Release,
			metricId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new audio: %w", err)
		}
	}
	
	return nil
}