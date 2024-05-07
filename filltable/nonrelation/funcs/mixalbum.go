package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/lib"
	"kra/constants"
	helper "kra/querryhelpers"
	funcshelpers "kra/filltable/helpers"
)

func deleteAllMixAlbum(db * sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"metric id\" FROM \"album\" " + 
		"WHERE appellation LIKE 'mixed%'",
	)
	if err != nil {
		return fmt.Errorf("cant get album && metric id: %w", err)
	}
	rows.Close()

	for rows.Next() {
		albumID, metricId := 0, 0
		if err := rows.Scan(&albumID, &metricId); err != nil {
			return fmt.Errorf("cant convert albumID && metric id: %w", err)
		}

		_, err := db.Exec(
			"DELETE FROM \"album\" " + 
			"WHERE id = $1",
			albumID,
		)
		if err != nil {
			return fmt.Errorf("cant del album by id: %w", err)
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

	return nil
}

func genAlbumByPrefix(db *sql.DB, albumPrefix string) error {
	metricId, err := funcshelpers.CreateMetric(db)
	if err != nil {
		return fmt.Errorf("cant insert new album metric: %w", err)
	}
	_, err = helper.InsertAlbum(
		db, 
		albumPrefix + "@" + lib.GetRandString(constants.AlbumName),
		fmt.Sprint(1950 + rand.Intn(2020 - 1950)),
		metricId,
	)
	if err != nil {
		return fmt.Errorf("cant insert new %s album: %w", albumPrefix, err)
	}
	
	return nil
}

func FillMixAlbum(db *sql.DB) error {
	if err := deleteAllMixAlbum(db); err != nil {
		return fmt.Errorf("cant del album table: %w", err)
	}

	// mixes
	for range constants.CntAlbumMix {
		if err := genAlbumByPrefix(db, "mixed"); err != nil {
			return err
		}
	}

	return nil
}