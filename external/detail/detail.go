package samehadakudetail

import (
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/radenrishwan/mcp-server-samehadaku/external"
)

var titlePrefix = "Nonton Anime "

func Fetch(slug string) (Anime, error) {
	url := external.BASE_URL + "anime/" + slug

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return Anime{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Anime{}, err
	}

	if doc.Find("#infoarea > div > div.infoanime.widget_senction > h2").Text() == "" {
		return Anime{}, external.ErrNotFound
	}

	anime := Anime{}
	anime.Title = strings.TrimPrefix(doc.Find("#infoarea > div > div.infoanime.widget_senction > h2").Text(), titlePrefix)
	anime.Thumbnail = doc.Find("#infoarea > div > div.infoanime.widget_senction > div.thumb > img").AttrOr("src", "")
	anime.Synopsis = doc.Find("#infoarea > div > div.infoanime.widget_senction > div.infox > div.desc > div > p").Text()

	// TODO: fix later, the genre are incorrect
	genres := []Genre{}
	doc.Find("#infoarea > div > div.infoanime.widget_senction > div.infox > div.genre-info").Each(func(i int, s *goquery.Selection) {
		genres = append(genres, Genre{
			Title: s.Find("a").Text(),
			Href:  s.Find("a").AttrOr("href", ""),
		})
	})

	anime.Genres = genres

	anime.Href = doc.Find("#infoarea > div > div:nth-child(26) > a").AttrOr("href", "")

	episodes := []Episode{}
	doc.Find("#infoarea > div > div.whites.lsteps.widget_senction > div.lstepsiode.listeps > ul > li").Each(func(i int, s *goquery.Selection) {
		left := s.Find("div.epsright")
		right := s.Find("div.epsleft")

		index, err := strconv.Atoi(left.Find("span > a").Text())
		if err != nil {
			index = -1
		}

		episodes = append(episodes, Episode{
			Index:      index,
			Title:      right.Find("span.lchx > a").Text(),
			ReleasedOn: right.Find("span.date").Text(),
			Href:       right.Find("span.lchx > a").AttrOr("href", ""),
		})
	})

	slices.Reverse(episodes)
	anime.Episodes = episodes

	// TODO: fix later, the index is still wrong
	doc.Find("#infoarea > div > div.whites.lsteps.widget_senction > div.anim-senct > div.right-senc.widget_senction > div > div > div > span").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			anime.Detail.Japanese = s.Text()
		case 1:
			anime.Detail.Status = s.Text()
		case 2:
			anime.Detail.Source = s.Text()
		case 3:
			anime.Detail.TotalEpisode = s.Text()
		case 4:
			anime.Detail.Studio = s.Text()
		case 5:
			anime.Detail.Released = s.Text()
		case 6:
			anime.Detail.Synonyms = s.Text()
		case 7:
			anime.Detail.Type = s.Text()
		case 8:
			anime.Detail.Duration = s.Text()
		case 9:
			anime.Detail.Season = s.Text()
		case 10:
			anime.Detail.Producers = s.Text()
		}
	})

	return anime, nil
}
