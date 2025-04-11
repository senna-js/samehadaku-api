package samehadakuhome

type Home struct {
	AnimeTerbaru   []AnimeCard        `json:"anime_terbaru"`
	ProjectMovie   []ProjectMovieCard `json:"project_movie"`
	DonghuaDanFilm []DonghuaDanFilm   `json:"donghua_dan_film,omitempty"`
	BatchAnime     []AnimeCard        `json:"batch_anime"`
}

type AnimeCard struct {
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	Episode    string `json:"episode"`
	PostedBy   string `json:"posted_by"`
	ReleasedOn string `json:"released_on"`
	Href       string `json:"href"`
	Slug       string `json:"slug"`
}

type ProjectMovieCard struct {
	Thumbnail  string  `json:"thumbnail"`
	Title      string  `json:"title"`
	Genres     []Genre `json:"genres"`
	ReleasedOn string  `json:"released_on"`
	Href       string  `json:"href"`
	Slug       string  `json:"slug"`
}

type Genre struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

// maybe we don't use this. because the url goes to outside samehadaku
type DonghuaDanFilm struct {
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Episode   string `json:"episode"`
	Rating    string `json:"rating"`
	Href      string `json:"href"`
	Slug      string `json:"slug"`
}
