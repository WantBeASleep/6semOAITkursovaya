package relation

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"

	helper "kra/querryhelpers"
)

func deleteAllAlbum_Audio(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"album_audio\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate album_audio table: %w", err)
	}

	return nil
}

func fillAlbum_Audio(db *sql.DB) error {
	if err := deleteAllAlbum_Audio(db); err != nil {
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
			if err := mixedAlbums.Scan(&trackID); err != nil {
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

	//author
	authorAlbums, err := db.Query(
		"SELECT id FROM \"author\" " +
		"WHERE appellation LIKE 'author%'",
	)
	if err != nil {
		return fmt.Errorf("cant get author albums ids: %w", err)
	}
	defer mixedAlbums.Close()

	authorIDs, err := db.Query(
		"SELECT id, COUNT(*) FROM \"author_audio\" " +
		"GROUP BY \"author id\" " +
		"HAVING COUNT(*) > 3",
	)
	if err != nil {
		return fmt.Errorf("cant get authors with many tracks: %w", err)
	}
	defer authorIDs.Close()

	for authorAlbums.Next() {
		albumID := 0
		if err := mixedAlbums.Scan(&albumID); err != nil {
			return fmt.Errorf("cant scan author album id: %w", err)
		}
		authorID, cntAuthorTracks := 0, 0
		if err := authorIDs.Scan(&authorID, &cntAuthorTracks); err != nil {
			return fmt.Errorf("cant scan author id with many tracks: %w", err)
		}

		tracksInAlbum, err := db.Query(
			"SELECT \"audio id\" FROM \"author_audio\" " +
			"WHERE \"author id\" = $1" + 
			"ORDER BY RANDOM() " +
			"LIMIT $2",
			authorID,
			1 + rand.Intn(cntAuthorTracks),
		)
		if err != nil {
			return fmt.Errorf("cant get tracks in author album: %w", err)
		}
		defer tracksInAlbum.Close()

		for tracksInAlbum.Next() {
			audioID := 0
			if err := tracksInAlbum.Scan(&audioID); err != nil {
				return fmt.Errorf("cant scan audio id in author album: %w", err)
			}

			err := helper.InsertAlbum_Audio(
				db,
				albumID,
				audioID,
			)
			if err != nil {
				return fmt.Errorf("cant insert new author album-audio: %w", err)
			}
		}

		authorIDs.Next()
	}
	
	return nil
}