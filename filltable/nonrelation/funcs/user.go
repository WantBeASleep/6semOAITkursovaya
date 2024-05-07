package funcs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"kra/constants"
	"kra/lib"
	helper "kra/querryhelpers"
)

func deleteAllUser(db * sql.DB) error {
	rows, err := db.Query(
		"SELECT id, \"author id\" FROM \"user\" ",
	)
	if err != nil {
		return fmt.Errorf("cant get user && author id id: %w", err)
	}
	rows.Close()

	for rows.Next() {
		userId, authorID := 0, 0
		if err := rows.Scan(&userId, &authorID); err != nil {
			return fmt.Errorf("cant convert userId && authorID id: %w", err)
		}

		_, err := db.Exec(
			"DELETE FROM \"user\" " + 
			"WHERE id = $1",
			userId,
		)
		if err != nil {
			return fmt.Errorf("cant del user by id: %w", err)
		}

		_, err = db.Exec(
			"DELETE FROM \"author\" " + 
			"WHERE id = $1",
			authorID,
		)
		if err != nil {
			return fmt.Errorf("cant del user by id: %w", err)
		}
	}

	_, err = db.Exec(
		"TRUNCATE TABLE \"user\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate user table: %w", err)
	}

	return nil
}

func FillUser(db *sql.DB) error {
	if err := deleteAllUser(db); err != nil {
		return fmt.Errorf("cant del user table: %w", err)
	}

	for range constants.CntUser {
		user := lib.GenUser()
		_, err := helper.InsertUser(
			db,
			user.Login,
			user.Email,
			user.Password,
			user.Role,
		)
		if err != nil {
			return fmt.Errorf("cant insert new user: %w", err)
		}
	}

	return nil
}