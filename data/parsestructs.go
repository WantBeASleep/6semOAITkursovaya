package data

type Audio struct {
	Id int
	Appellation string
	Lyric string
	Release string
}

type Genre struct {
	Id int
	Appellation string
	Description string
}

type Author struct {
	Id int
	Appellation string
	Description string
}

type Snippet struct {
	Start int
	End int
}

type TrackInfo struct {
	Audio Audio
	Genre Genre
	Author Author
	Snippet Snippet
}