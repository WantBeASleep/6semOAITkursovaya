package nonrelation

import (
	"database/sql"
	helpers "kra/filltable/helpers"
	"kra/filltable/model"
	funcs "kra/filltable/nonrelation/funcs"
)

func GetFuncs(db *sql.DB) []model.TableFunc {
	return []model.TableFunc{
		{
			TableName: "audio",
			TableStatus: helpers.GetTableStatus(db, "audio"),
			Efunc: funcs.FillAudio,
		},
		{
			TableName: "author",
			TableStatus: helpers.GetTableStatus(db, "author"),
			Efunc: funcs.FillAuthor,
		},
		{
			TableName: "genre",
			TableStatus: helpers.GetTableStatus(db, "genre"),
			Efunc: funcs.FillGenre,
		},
		{
			TableName: "mixalbum",
			TableStatus: helpers.GetTableStatus(db, "album"),
			Efunc: funcs.FillMixAlbum,
		},
		{
			TableName: "user",
			TableStatus: helpers.GetTableStatus(db, "user"),
			Efunc: funcs.FillUser,
		},
	}
}