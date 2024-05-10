package reports

import (
	"database/sql"
	helpers "kra/filltable/helpers"
	"kra/filltable/model"
	funcs "kra/filltable/reports/funcs"
)

func GetFuncs(db *sql.DB) []model.TableFunc {
	return []model.TableFunc{
		{
			TableName:   "do bar genre hist",
			TableStatus: helpers.GetTableStatus(db, "audio"),
			Efunc:       funcs.DoHistogramUser_Audio,
		},
		{
			TableName:   "do bar years hist",
			TableStatus: helpers.GetTableStatus(db, "audio"),
			Efunc:       funcs.DoHistogrammAudio_Years,
		},
	}
}
