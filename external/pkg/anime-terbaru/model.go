package samehadakuanimeterbaru

type AnimeTerbaru struct {
	Page  int         `json:"page"`
	Anime []AnimeCard `json:"anime"`
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
