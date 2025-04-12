package samehadakuepisode

type EpisodeResult struct {
	Title       string `json:"title"`
	Episode     int    `json:"episode"`
	ReleaseDate string `json:"release_date"`
	// use this data to get a iframe url using function `samehadakuapi.GetIframeUrl`
	Streams      []Stream       `json:"stream"`
	StreamUlrls  []StreamUlrl   `json:"stream_url,omitempty"`
	DownloadUrls []DownloadUrls `json:"download_urls"`
}

type Stream struct {
	Title string `json:"title"`
	// use `Post` to get a iframe url using function `samehadakuapi.GetIframeUrl`
	Post string `json:"post"`
	// use `Type` to get a iframe url using function `samehadakuapi.GetIframeUrl`
	Type string `json:"type"`
	// use `Nume` to get a iframe url using function `samehadakuapi.GetIframeUrl`
	Nume string `json:"nume"`
}

type StreamUlrl struct {
	Title     string `json:"title"`
	IframeUrl string `json:"iframe_url"`
}

type DownloadUrls struct {
	VideoFormat string        `json:"video_format"`
	DownloadUrl []DownloadUrl `json:"download_url"`
}

type DownloadUrl struct {
	Quality string `json:"quality"`
	Urls    []Url  `json:"urls"`
}

type Url struct {
	Host string `json:"host"`
	Url  string `json:"url"`
}
