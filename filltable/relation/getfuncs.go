package relation

import (
	"database/sql"
	helpers "kra/filltable/helpers"
	"kra/filltable/model"
	funcs "kra/filltable/relation/funcs"
)

func GetFuncs(db *sql.DB) []model.TableFunc {
	return []model.TableFunc{
		{
			TableName:   "snippet",
			TableStatus: helpers.GetTableStatus(db, "snippet"),
			Efunc:       funcs.FillSnippet,
		},
		{
			TableName:   "external resource",
			TableStatus: helpers.GetTableStatus(db, "external resource"),
			Efunc:       funcs.FillExternal,
		},
		{
			TableName:   "audio_genre",
			TableStatus: helpers.GetTableStatus(db, "audio_genre"),
			Efunc:       funcs.FillAudio_Genre,
		},
		{
			TableName:   "author_audio",
			TableStatus: helpers.GetTableStatus(db, "author_audio"),
			Efunc:       funcs.FillAuthor_Audio,
		},
		{
			TableName:   "user_audio",
			TableStatus: helpers.GetTableStatus(db, "user_audio"),
			Efunc:       funcs.FillUser_Audio,
		},
		{
			TableName:   "relation user-author",
			TableStatus: helpers.GetTableStatus(db, "user"),
			Efunc:       funcs.RelationUser_Author,
		},
		{
			TableName:   "mix album_audio",
			TableStatus: helpers.GetTableStatus(db, "album_audio"),
			Efunc:       funcs.FillMixAlbum_Audio,
		},
		{
			TableName:   "album_genre",
			TableStatus: helpers.GetTableStatus(db, "album_genre"),
			Efunc:       funcs.FillAlbum_Genre,
		},
		{
			TableName:   "author_album",
			TableStatus: helpers.GetTableStatus(db, "author_album"),
			Efunc:       funcs.FillAuthor_Album,
		},
		{
			TableName:   "user_album",
			TableStatus: helpers.GetTableStatus(db, "user_album"),
			Efunc:       funcs.FillUser_Album,
		},
	}
}
