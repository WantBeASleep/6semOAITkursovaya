package relation

import (
	"kra/scripts/model"
)

func GetFuncs() []model.FuncInfo {
	return []model.FuncInfo{
		{
			Name: "fill audio_genre",
			Efunc: fillAudio_Genre,
		},
		{
			Name: "fill audio_user",
			Efunc: fillAudio_User,
		},
		{
			Name: "fill author_audio",
			Efunc: fillAuthor_Audio,
		},
		{
			Name: "fill album_audio",
			Efunc: fillAlbum_Audio,
		},
		{
			Name: "fill album_genre",
			Efunc: fillAlbum_Genre,
		},
		{
			Name: "fill album_user",
			Efunc: fillAlbum_User,
		},
		{
			Name: "fill album_author",
			Efunc: fillAlbum_Author,
		},
		{
			Name: "fill relation user-author",
			Efunc: relationUser_Author,
		},
	}
}