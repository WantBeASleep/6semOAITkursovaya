package nonrelation

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/lib"
	"kra/constants"
	helper "kra/querryhelpers"
)

func deleteAllAlbum(db * sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"metric id\" FROM \"album\" ",
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

	_, err = db.Exec(
		"TRUNCATE TABLE \"album\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate album table: %w", err)
	}

	return nil
}

func genAlbumByPrefix(db *sql.DB, albumPrefix string) error {
	metricId, err := createMetric(db)
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

func fillAlbum(db *sql.DB) error {
	if err := deleteAllAlbum(db); err != nil {
		return fmt.Errorf("cant del album table: %w", err)
	}

	// mixes
	for range constants.CntAlbumMix {
		if err := genAlbumByPrefix(db, "mixed"); err != nil {
			return err
		}
	}

	//authors
	authorWithManyTracksCnt := 0
	err := db.QueryRow(
		"SELECT COUNT(*), FROM \"author_audio\" " +
		"GROUP BY \"author id\" " +
		"HAVING COUNT(*) > 3",
	).Scan(&authorWithManyTracksCnt)
	if err != nil {
		return fmt.Errorf("cant get authors with many tracks: %w", err)
	}

	for range authorWithManyTracksCnt {
		if err := genAlbumByPrefix(db, "author"); err != nil {
			return err
		}
	}

	return nil
}