package samehadakusearch

type Search struct {
	Result []AnimeCard `json:"result"`
}

type AnimeCard struct {
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Rating    string `json:"rating"`
	Href      string `json:"href"`
	Slug      string `json:"slug"`
}
