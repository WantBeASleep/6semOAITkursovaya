package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"
	helper "kra/querryhelpers"
)

func deleteAllUser_Audio(db *sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"user_audio\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate user_audio table: %w", err)
	}

	return nil
}

func FillUser_Audio(db *sql.DB) error {
	if err := deleteAllUser_Audio(db); err != nil {
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

		audioIds, err := db.Query(
			"SELECT id FROM \"audio\" "+
				"ORDER BY RANDOM() "+
				"LIMIT $1",
			1+rand.Intn(constants.TopCntUserAudio),
		)
		if err != nil {
			return fmt.Errorf("cant get audio ids: %w", err)
		}
		defer audioIds.Close()

		for audioIds.Next() {
			audioID := 0
			err := audioIds.Scan(&audioID)
			if err != nil {
				return fmt.Errorf("cant convert audioId: %w", err)
			}

			err = helper.InsertUser_Audio(
				db,
				userId,
				audioID,
			)
			if err != nil {
				return fmt.Errorf("cant insert user audio: %w", err)
			}
		}
	}

	return nil
}
