package data

type Audio struct {
	Appellation string
	Lyric string
	Release string
}

type Genre struct {
	Appellation string
	Description string
}

type Author struct {
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