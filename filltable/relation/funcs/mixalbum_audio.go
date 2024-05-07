package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"

	helper "kra/querryhelpers"
)

func deleteAllMixAlbum_Audio(db * sql.DB) error {
	mixIds, err := db.Query(
		"SELECT album.id FROM " + 
		"\"album_audio\" " + 
		"JOIN " + 
		"album " + 
		"ON album.id = album_audio.\"album id\" " + 
		"WHERE album.appellation LIKE 'mixed%'",
	)
	if err != nil {
		return fmt.Errorf("cant get mixed albums ids: %w", err)
	}

	for mixIds.Next() {
		albumID := 0
		err := mixIds.Scan(&albumID)
		if err != nil {
			return fmt.Errorf("cant convert mix album id: %w", err)
		}

		_, err = db.Exec(
			"DELETE FROM album_audio " + 
			"WHERE \"album id\" = $1",
			albumID,
		)
		if err != nil {
			return fmt.Errorf("cant del mix album relation: %w", err)
		}
	}

	return nil
}

func FillMixAlbum_Audio(db *sql.DB) error {
	if err := deleteAllMixAlbum_Audio(db); err != nil {
		return fmt.Errorf("cant del album_audio table: %w", err)
	}

	//mixed
	mixedAlbums, err := db.Query(
		"SELECT id FROM \"album\" " +
		"WHERE appellation LIKE 'mixed%'",
	)
	if err != nil {
		return fmt.Errorf("cant get mixed albums ids: %w", err)
	}
	defer mixedAlbums.Close()

	for mixedAlbums.Next() {
		albumID := 0
		if err := mixedAlbums.Scan(&albumID); err != nil {
			return fmt.Errorf("cant scan mixed album id: %w", err)
		}

		trackInAlbum := 1 + rand.Intn(constants.TopCntAudioInAlbum)
		randTracks, err := db.Query(
			"SELECT id FROM \"audio\" " +
			"ORDER BY RANDOM() " +
			"LIMIT $1",
			trackInAlbum,
		)
		if err != nil {
			return fmt.Errorf("cant parse random tracks for album")
		}
		defer randTracks.Close()

		for randTracks.Next() {
			trackID := 0
			if err := randTracks.Scan(&trackID); err != nil {
				return fmt.Errorf("cant random track ID: %w", err)
			}

			err = helper.InsertAlbum_Audio(
				db,
				albumID,
				trackID,
			)
			if err != nil {
				return fmt.Errorf("cant insert new album-audio: %w", err)
			}
		}
	}

	return nil
}