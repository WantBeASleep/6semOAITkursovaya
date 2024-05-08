package funcs

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"

	"kra/constants"
	"kra/lib"
	helper "kra/querryhelpers"
)

func deleteAllExternal(db *sql.DB) error {
	_, err := db.Exec(
		"TRUNCATE TABLE \"external resource\" RESTART IDENTITY CASCADE",
	)
	if err != nil {
		return fmt.Errorf("cant truncate external resource table: %w", err)
	}

	return nil
}

func FillExternal(db *sql.DB) error {
	if err := deleteAllExternal(db); err != nil {
		return fmt.Errorf("cant del snippet table: %w", err)
	}

	authorIds, err := db.Query(
		"SELECT id FROM author",
	)
	if err != nil {
		return fmt.Errorf("cant select id authors: %w", err)
	}
	defer authorIds.Close()

	for authorIds.Next() {
		if rand.Float64() > constants.PercentExternal {
			continue
		}

		authorId := 0
		err = authorIds.Scan(&authorId)
		if err != nil {
			return fmt.Errorf("cant scan author ID: %w", err)
		}
		res := lib.GenResource()
		_, err = helper.InsertExternal(
			db,
			res.Link,
			res.RType,
			authorId,
		)
		if err != nil {
			return fmt.Errorf("cant insert new external: %w", err)
		}
	}

	return nil
}
