package querryhelpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InsertMetric(
	db *sql.DB,
	views int,
	likes int,
	reposts int,
	retention float64,
	downloads int,
	year_popularity float64,
) (int, error) {
	response := 0
	err := db.QueryRow(
		"INSERT INTO \"metric\" (\"views\", \"likes\", \"reposts\", \"retention\", \"downloads\", \"year-popularity\") " + 
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", 
		views,
		likes, 
		reposts,
		retention,
		downloads, 
		year_popularity,
	).Scan(&response)
	return response, err
}