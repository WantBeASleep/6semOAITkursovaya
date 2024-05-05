package nonrelation

import (
	"kra/scripts/model"
)

func GetFuncs() []model.FuncInfo {
	return []model.FuncInfo{
		{
			Name: "fill audio",
			Efunc: fillAudio,
		},
		{
			Name: "fill genre",
			Efunc: fillGenre,
		},
		{
			Name: "fill album",
			Efunc: fillAlbum,
		},
		{
			Name: "fill author",
			Efunc: fillAuthor,
		},
		{
			Name: "fill user",
			Efunc: fillUser,
		},
		{
			Name: "fill snippet",
			Efunc: fillSnippet,
		},
		{
			Name: "fill external",
			Efunc: fillExternal,
		},
	}
}