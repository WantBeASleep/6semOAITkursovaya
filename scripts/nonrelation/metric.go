package nonrelation

import (
	"fmt"
	"math/rand"

	"database/sql"
	"kra/constants"
	helper "kra/querryhelpers"

	_ "github.com/lib/pq"
)

func createMetric(db *sql.DB) (int, error) {
	metricId, err := helper.InsertMetric(
		db,
		rand.Intn(constants.MetricViews),
		rand.Intn(constants.MetricLikes),
		rand.Intn(constants.MetricReposts),
		rand.Float64() * constants.MetricRetention,
		rand.Intn(constants.MetricDownloads),
		rand.Float64() * constants.MetricYearPopularity,
	)
	if err != nil {
		return metricId, fmt.Errorf("cant insert new metric: %w", err)
	}
	return metricId, nil
}