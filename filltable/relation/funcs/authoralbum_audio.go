package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	helper "kra/querryhelpers"
)

func deleteAllAuthorAlbum_Audio(db * sql.DB) error {
	authorIDs, err := db.Query(
		"SELECT album.id FROM " + 
		"\"album_audio\" " + 
		"JOIN " + 
		"album " + 
		"ON album.id = album_audio.\"album id\" " + 
		"WHERE album.appellation LIKE 'author%'",
	)
	if err != nil {
		return fmt.Errorf("cant get auhtor albums ids: %w", err)
	}

	for authorIDs.Next() {
		albumID := 0
		err := authorIDs.Scan(&albumID)
		if err != nil {
			return fmt.Errorf("cant convert auhthor album id: %w", err)
		}

		_, err = db.Exec(
			"DELETE FROM album_audio " + 
			"WHERE \"album id\" = $1",
			albumID,
		)
		if err != nil {
			return fmt.Errorf("cant del author album relation: %w", err)
		}
	}

	return nil
}

func FillAuthorAlbum_Audio(db *sql.DB) error {
	if err := deleteAllAuthorAlbum_Audio(db); err != nil {
		return fmt.Errorf("cant del album_audio table: %w", err)
	}

	//author
	authorAlbums, err := db.Query(
		"SELECT id FROM \"author\" " +
		"WHERE appellation LIKE 'author%'",
	)
	if err != nil {
		return fmt.Errorf("cant get author albums ids: %w", err)
	}
	defer authorAlbums.Close()

	authorIDs, err := db.Query(
		"SELECT \"author id\", COUNT(*) FROM \"author_audio\" " +
		"GROUP BY \"author id\" " +
		"HAVING COUNT(*) > 3",
	)
	if err != nil {
		return fmt.Errorf("cant get authors with many tracks: %w", err)
	}
	defer authorIDs.Close()

	arrayOfAuthors := [][]int{}
	for authorIDs.Next() {
		authorID, cntAuthorTracks := 0, 0
		if err := authorIDs.Scan(&authorID, &cntAuthorTracks); err != nil {
			return fmt.Errorf("cant scan author id with many tracks: %w", err)
		}
		arrayOfAuthors = append(arrayOfAuthors, []int{authorID, cntAuthorTracks})
	}

	i := 0
	for authorAlbums.Next() {
		albumID := 0
		if err := authorAlbums.Scan(&albumID); err != nil {
			return fmt.Errorf("cant scan author album id: %w", err)
		}
		
		if i == len(arrayOfAuthors) {
			continue
		}

		authorID, cntAuthorTracks := arrayOfAuthors[i][0], arrayOfAuthors[i][1]

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
		i++
	}
	
	return nil
}