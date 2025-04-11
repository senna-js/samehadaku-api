package samehadakuepisode

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/radenrishwan/samehadaku-api/external"
)

func Fetch(slug string, fetchStreamUrl bool) (Episode, error) {
	url := external.BASE_URL + slug

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return Episode{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Episode{}, err
	}

	episode := Episode{}
	episode.Title = doc.Find("div.player-area.widget_senction > header > div > h1.entry-title").Text()

	episodeIndex, err := strconv.Atoi(doc.Find("div.player-area.widget_senction > header > div > div.sbdbti > span > span:nth-child(2)").Text())
	if err != nil {
		episodeIndex = -1
	}

	episode.Episode = episodeIndex
	// TODO: fix later, stll wrong format
	episode.ReleaseDate = doc.Find("div.player-area.widget_senction > header > div > div.sbdbti > span").Text()

	streams := []Stream{}
	doc.Find("#server > ul > li > div").Each(func(i int, s *goquery.Selection) {
		stream := Stream{}
		stream.Title = s.Find("span").Text()
		stream.Post = s.AttrOr("data-post", "-1")
		stream.Type = s.AttrOr("data-type", "-1")
		stream.Nume = s.AttrOr("data-nume", "-1")

		streams = append(streams, stream)
	})

	episode.Streams = streams

	if fetchStreamUrl {
		streamUrls := []StreamUlrl{}
		for _, ep := range episode.Streams {
			iframeUrl, _ := external.GetIFrameURL(url, external.IFrameBody{
				Post:         ep.Post,
				ResponseType: ep.Type,
				Nume:         ep.Nume,
			})

			streamUrls = append(streamUrls, StreamUlrl{
				Title:     ep.Title,
				IframeUrl: iframeUrl,
			})
		}

		episode.StreamUlrls = streamUrls
	}

	downloadUrls := []DownloadUrls{}
	doc.Find("#downloadb").Each(func(i int, s *goquery.Selection) {
		urls := DownloadUrls{}
		urls.VideoFormat = s.Find("p > b").Text()

		downloadUrl := []DownloadUrl{}
		s.Find("ul > li").Each(func(i int, s2 *goquery.Selection) {
			dd := DownloadUrl{}
			dd.Quality = s2.Find("strong").Text()

			urls := []Url{}
			s2.Find("span > a").Each(func(i int, s3 *goquery.Selection) {
				url := Url{}

				url.Host = s3.Text()
				url.Url = s3.AttrOr("href", "")

				urls = append(urls, url)
			})

			dd.Urls = urls

			downloadUrl = append(downloadUrl, dd)
		})

		urls.DownloadUrl = downloadUrl

		downloadUrls = append(downloadUrls, urls)
	})

	episode.DownloadUrls = downloadUrls

	return episode, nil
}
