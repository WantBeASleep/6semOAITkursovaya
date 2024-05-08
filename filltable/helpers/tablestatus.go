package helpers

import (
	"database/sql"
	"fmt"
	color "kra/cli/color"
)

var c = color.GetColorSprintFuncs()

func GetTableStatus(db *sql.DB, tableName string) string {
	countInTable := 0
	db.QueryRow(
		fmt.Sprintf("SELECT count(*) FROM \"%s\"", tableName),
	).Scan(&countInTable)

	res := fmt.Sprintf("Table %s have %d zapisey", tableName, countInTable)
	if countInTable == 0 {
		return c.Red(res)
	} else {
		return c.Green(res)
	}
}
