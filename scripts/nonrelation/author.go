package nonrelation

import (
	"database/sql"
	"fmt"

	"encoding/csv"
	"kra/constants"
	"os"
	"path"

	_ "github.com/lib/pq"

	"kra/data"
	helper "kra/querryhelpers"
)

func deleteAllAuthor(db * sql.DB) error {
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
			"DELETE FROM \"author\" " + 
			"WHERE id = $1",
			authorId,
		)
		if err != nil {
			return fmt.Errorf("cant del author by id: %w", err)
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
		"TRUNCATE TABLE \"author\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate author table: %w", err)
	}

	return nil
}

func fillAuthor(data []data.TrackInfo, db *sql.DB) error {
	if err := deleteAllAuthor(db); err != nil {
		return fmt.Errorf("cant del author table: %w", err)
	}

	curDir, _ := os.Getwd()
	idsCsvFile, _ := os.Create(path.Join(curDir, constants.CostylAudio))
	defer idsCsvFile.Close()
	w := csv.NewWriter(idsCsvFile)
	defer w.Flush()
	idsCsv := [][]string{}

	for _, track := range data {
		cntAuthorByName := 0
		err := db.QueryRow(
			"SELECT COUNT(*) FROM \"author\" " +
			"WHERE appellation = $1",
			track.Author.Appellation,
		).Scan(&cntAuthorByName)
		if err != nil {
			return fmt.Errorf("cant check uniq author: %w", err)
		}

		authorID := 0
		isUniq := cntAuthorByName == 0
		if isUniq {
			metricId, err := createMetric(db)
			if err != nil {
				return fmt.Errorf("cant insert new author metric: %w", err)
			}
			authorID, err = helper.InsertAuthor(
				db,
				track.Author.Appellation,
				track.Author.Description,
				metricId,
			)
			if err != nil {
				return fmt.Errorf("cant insert new author: %w", err)
			}
		} else {
			err := db.QueryRow(
				"SELECT id FROM \"author\" " +
				"WHERE appellation = $1",
				track.Author.Appellation,
			).Scan(&authorID)
			if err != nil {
				return fmt.Errorf("cant parse author ID: %w", err)
			}
		}
		track.Author.Id = authorID
		idsCsv = append(idsCsv, []string{fmt.Sprint(authorID)})
	}

	w.WriteAll(idsCsv)
	
	return nil
}