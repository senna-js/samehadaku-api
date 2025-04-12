package samehadakujadwalrilis

type JadwalRilisResult struct {
	Thumbnail string   `json:"thumbnail"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`
	Rating    string   `json:"rating"`
	Time      string   `json:"time"`
	Genre     []string `json:"genre"`
	Href      string   `json:"href"`
	Slug      string   `json:"slug"`
}

// this struct used to parse a api response
type animeResponse struct {
	ID           int    `json:"id"`
	Slug         string `json:"slug"`
	Date         string `json:"date"`
	Author       string `json:"author"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Content      string `json:"content"`
	FeaturedImg  string `json:"featured_img_src"`
	Genre        string `json:"genre"`
	EastScore    string `json:"east_score"`
	EastType     string `json:"east_type"`
	EastSchedule string `json:"east_schedule"`
	EastTime     string `json:"east_time"`
}
