package samehadakudetail

import (
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/radenrishwan/samehadaku-api/utility"
)

var titlePrefix = "Nonton Anime "

type Detail struct {
	BaseUrl string
}

func (self Detail) Fetch(slug string) (DetailResult, error) {
	url := self.BaseUrl + "anime/" + slug

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return DetailResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return DetailResult{}, err
	}

	if doc.Find("#infoarea > div > div.infoanime.widget_senction > h2").Text() == "" {
		return DetailResult{}, utility.ErrNotFound
	}

	anime := DetailResult{}
	anime.Title = strings.TrimPrefix(doc.Find("#infoarea > div > div.infoanime.widget_senction > h2").Text(), titlePrefix)
	anime.Thumbnail = doc.Find("#infoarea > div > div.infoanime.widget_senction > div.thumb > img").AttrOr("src", "")
	anime.Synopsis = doc.Find("#infoarea > div > div.infoanime.widget_senction > div.infox > div.desc > div > p").Text()

	genres := []Genre{}
	doc.Find("#infoarea > div > div.infoanime.widget_senction > div.infox > div.genre-info > a").Each(func(i int, s *goquery.Selection) {
		genres = append(genres, Genre{
			Title: s.Text(),
			Href:  s.AttrOr("href", ""),
		})
	})

	anime.Genres = genres

	anime.Href = doc.Find("#infoarea > div > div:nth-child(26) > a").AttrOr("href", "")
	anime.EpisodeSlug = utility.ExtractSlug(anime.Href)

	episodes := []Episode{}
	doc.Find("#infoarea > div > div.whites.lsteps.widget_senction > div.lstepsiode.listeps > ul > li").Each(func(i int, s *goquery.Selection) {
		left := s.Find("div.epsright")
		right := s.Find("div.epsleft")

		index, err := strconv.Atoi(left.Find("span > a").Text())
		if err != nil {
			index = -1
		}

		episodes = append(episodes, Episode{
			Index:       index,
			Title:       right.Find("span.lchx > a").Text(),
			ReleasedOn:  right.Find("span.date").Text(),
			Href:        right.Find("span.lchx > a").AttrOr("href", ""),
			EpisodeSlug: utility.ExtractSlug(right.Find("span.lchx > a").AttrOr("href", "")),
		})
	})

	slices.Reverse(episodes)
	anime.Episodes = episodes

	doc.Find("#infoarea > div > div.whites.lsteps.widget_senction > div.anim-senct > div.right-senc.widget_senction > div > div > div > span").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			anime.Detail.Japanese = extractDetailText(s.Text())
		case 1:
			anime.Detail.Synonyms = extractDetailText(s.Text())
		case 2:
			anime.Detail.Status = extractDetailText(s.Text())
		case 3:
			anime.Detail.Type = extractDetailText(s.Text())
		case 4:
			anime.Detail.Source = extractDetailText(s.Text())
		case 5:
			anime.Detail.Duration = extractDetailText(s.Text())
		case 6:
			anime.Detail.TotalEpisode = extractDetailText(s.Text())
		case 7:
			anime.Detail.Season = extractDetailText(s.Text())
		case 8:
			anime.Detail.Studio = extractDetailText(s.Text())
		case 9:
			anime.Detail.Producers = extractDetailText(s.Text())
		case 10:
			anime.Detail.Released = extractDetailText(s.Text())
		}
	})

	return anime, nil
}

func extractDetailText(text string) string {
	t := strings.Split(text, " ")
	if len(t) < 2 && t[1] == "" {
		return "-"
	}

	return t[1]
}
