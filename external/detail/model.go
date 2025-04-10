package samehadakudetail

type Anime struct {
	Thumbnail string    `json:"thumbnail"`
	Title     string    `json:"title"`
	Synopsis  string    `json:"synopsis"`
	Genres    []Genre   `json:"genres"`
	Episodes  []Episode `json:"episodes"`
	Detail    Detail    `json:"detail"`
	// url to newest or latest episode
	Href string `json:"href"`
}

type Genre struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type Episode struct {
	Index      int    `json:"index"`
	Title      string `json:"title"`
	ReleasedOn string `json:"released_on"`
	Href       string `json:"href"`
}

type Detail struct {
	Japanese     string `json:"japanese"`
	Status       string `json:"status"`
	Source       string `json:"source"`
	TotalEpisode string `json:"total_episode"`
	Studio       string `json:"studio"`
	Released     string `json:"released"`
	Synonyms     string `json:"synonyms"`
	Type         string `json:"type"`
	Duration     string `json:"duration"`
	Season       string `json:"season"`
	Producers    string `json:"producers"`
}
