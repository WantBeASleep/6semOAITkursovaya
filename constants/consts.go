package constants

// относительно main.go
const (
	DATASET_PATH  = "data/tcc_ceds_music.csv"
	TRUNCATE_PATH = "sql/truncateTables.sql"
)

const (
	CntAudio    = -1    // parse from csv
	CntGenre    = -1    // parse from csv
	CntAuthor   = -1    //parse from csv
	CntAlbum    = -1    // parse from csv*
	CntAlbumMix = 3000  //40000
	CntUser     = 15000 // 5000000
)

// * - авторские/mix альбомы, авторские - по альбому на исполнителя, mix - const

const (
	PercentExternal = 0.6
	PercentSnippets = 0.3
)

const (
	TopCntAudioInAlbum = 7
	TopCntUserAudio    = 10
	TopCntUserAlbum    = 3
)

const (
	LinkMaxId = 1000000000

	LoginMaxGenLen    = 20
	EmailMaxGenLen    = 20
	PasswordMaxGenLen = 20

	AlbumName = 20
)

const (
	MetricViews          = 1000000
	MetricLikes          = 500000
	MetricReposts        = 300000
	MetricRetention      = 100 // [0, 100] float
	MetricDownloads      = 500000
	MetricYearPopularity = 100 // [0, 100] float
)
