package relation

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	helper "kra/querryhelpers"
)

func deleteAllAlbum_Genre(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"album_genre\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate album_genre table: %w", err)
	}

	return nil
}

func fillAlbum_Genre(db *sql.DB) error {
	if err := deleteAllAlbum_Genre(db); err != nil {
		return fmt.Errorf("cant del album_audio table: %w", err)
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

		topGenreIdInAlbum := 0
		err = db.QueryRow(
			"SELECT genre.id FROM " + 
			"album_audio " +
			"WHERE \"album_audio.album id\" = $1 " + 
			"JOIN audio " + 
			"ON audio.id = \"album_audio.audio id\" " + 
			"JOIN genre " + 
			"ON genre.id = audio.id " +
			"GROUP BY genre.appellation " + 
			"ORDER BY COUNT(*) DESC " + 
			"LIMIT 1",
			albumID,
		).Scan(&topGenreIdInAlbum)
		if err != nil {
			return fmt.Errorf("cant get top genre in album: %w", err)
		}

		err := helper.InsertAlbum_Genre(
			db,
			albumID,
			topGenreIdInAlbum,
		)
		if err != nil {
			return fmt.Errorf("cant insert new album-genre: %w", err)
		}
	}
	
	return nil
}