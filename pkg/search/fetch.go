package samehadakusearch

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/senna-js/samehadaku-api/utility"
)

type Search struct {
	BaseUrl string
}

func (self Search) Fetch(query string) (SearchResult, error) {
	// https://samehadaku.mba/daftar-anime-2/?title=naruto&status=&type=&order=title
	// TODO: add more query params
	url := self.BaseUrl + "/daftar-anime-2/?title=" + query

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return SearchResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return SearchResult{}, err
	}

	search := SearchResult{}

	animes := []AnimeCard{}
	doc.Find("div > div.animposx > a").Each(func(i int, s *goquery.Selection) {
		anime := AnimeCard{}
		anime.Href = s.AttrOr("href", "")
		anime.Slug = utility.ExtractSlug(anime.Href)
		anime.Title = s.AttrOr("title", "")

		anime.Thumbnail = s.Find("div.content-thumb > img").AttrOr("src", "")
		anime.Type = s.Find("div.content-thumb > div:nth-child(3)").Text()
		anime.Rating = strings.TrimSpace(s.Find(".score").Text())

		anime.Status = s.Find("div.data > div.type").Text()

		animes = append(animes, anime)
	})

	search.Result = animes

	return search, nil
}
