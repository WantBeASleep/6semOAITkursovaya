package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	helper "kra/querryhelpers"
	"kra/constants"
)

func deleteAllUser_Album(db * sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"user_album\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate user_album table: %w", err)
	}

	return nil
}

func FillUser_Album(db *sql.DB) error {
	if err := deleteAllUser_Album(db); err != nil {
		return fmt.Errorf("cant del user_audio table: %w", err)
	}

	usersId, err := db.Query(
		"SELECT id FROM \"user\"",
	)
	if err != nil {
		return fmt.Errorf("cant get users id: %w", err)
	}
	defer usersId.Close()

	for usersId.Next() {
		userId := 0
		err := usersId.Scan(&userId)
		if err != nil {
			return fmt.Errorf("cant convert userid: %w", err)
		}

		albumIds, err := db.Query(
			"SELECT id FROM \"album\" " +
			"ORDER BY RANDOM() " +
			"LIMIT $1",
			1 + rand.Intn(constants.TopCntUserAlbum),
		)
		if err != nil {
			return fmt.Errorf("cant get album ids: %w", err)
		}
		defer albumIds.Close()

		for albumIds.Next() {
			albumID := 0
			err := albumIds.Scan(&albumID)
			if err != nil {
				return fmt.Errorf("cant convert albumID: %w", err)
			}

			err = helper.InsertUser_Album(
				db,
				userId,
				albumID,
			)
			if err != nil {
				return fmt.Errorf("cant insert user album: %w", err)
			}
		}
	}
	
	return nil
}