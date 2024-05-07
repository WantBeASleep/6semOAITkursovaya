package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	helper "kra/querryhelpers"
)

func deleteAllAuthor_Album(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"author_album\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate author_album table: %w", err)
	}

	return nil
}

func FillAuthor_Album(db *sql.DB) error {
	if err := deleteAllAuthor_Album(db); err != nil {
		return fmt.Errorf("cant del author_album table: %w", err)
	}

	albumIds, err := db.Query(
		"SELECT id FROM \"album\"",
	)
	if err != nil {
		return fmt.Errorf("cant get albums ids: %w", err)
	}
	defer albumIds.Close()

	for albumIds.Next() {
		albumID := 0
		if err := albumIds.Scan(&albumID); err != nil {
			return fmt.Errorf("cant scan album id: %w", err)
		}

		authorsInAlbum, err := db.Query(
			"SELECT author_audio.\"author id\" FROM " + 
			"(SELECT * FROM album_audio WHERE album_audio.\"album id\" = $1) as album_audio " +
			"JOIN author_audio " + 
			"ON author_audio.\"audio id\" = album_audio.\"audio id\" " + 
			"GROUP BY author_audio.\"author id\"",
			albumID,
		)
		if err != nil {
			return fmt.Errorf("cant get authors of album: %w", err)
		}
		defer albumIds.Close()

		for authorsInAlbum.Next() {
			authorID := 0
			if err := authorsInAlbum.Scan(&authorID); err != nil {
				return fmt.Errorf("cant scan author id: %w", err)
			}

			err := helper.InsertAuthor_Album(
				db,
				authorID,
				albumID,
			)
			if err != nil {
				return fmt.Errorf("cant insert new authior-album: %w", err)
			}
		}
	}
	
	return nil
}