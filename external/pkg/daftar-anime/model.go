package samehadakudaftaranime

type DaftarAnime struct {
	Alphabets []Alphabet `json:"alphabets,omitempty"`
	Animes    []Anime    `json:"animes,omitempty"`
}

type Alphabet struct {
	Alphabet string  `json:"alphabet"`
	Animes   []Anime `json:"animes"`
}

type Anime struct {
	Title string `json:"title"`
	Href  string `json:"href"`
	Slug  string `json:"slug"`
}
