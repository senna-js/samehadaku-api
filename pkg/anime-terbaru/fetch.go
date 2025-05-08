package samehadakuanimeterbaru

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/senna-js/samehadaku-api/utility"
)

type AnimeTerbaru struct {
	BaseUrl string
}

func (self AnimeTerbaru) Fetch(page int) (AnimeTerbaruResult, error) {
	if page < 0 {
		return AnimeTerbaruResult{}, utility.ErrNotFound
	}

	url := self.BaseUrl + "anime-terbaru/page/" + strconv.Itoa(page)

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return AnimeTerbaruResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return AnimeTerbaruResult{}, err
	}

	animeTerbaru := AnimeTerbaruResult{}

	animes := []AnimeCard{}
	doc.Find("#main > div.post-show > ul > li").Each(func(i int, s *goquery.Selection) {
		thumb := s.Find("div.thumb > a > img")
		dlta := s.Find("div.dtla")

		animes = append(animes, AnimeCard{
			Thumbnail:  thumb.AttrOr("src", ""),
			Title:      dlta.Find("h2 > a").Text(),
			Episode:    dlta.Find("span:nth-child(2) > author").Text(),
			PostedBy:   dlta.Find("span:nth-child(3) > author").Text(),
			ReleasedOn: dlta.Find("span:nth-child(4) > author").Text(),
			Href:       dlta.Find("h2 > a").AttrOr("href", ""),
			Slug:       utility.ExtractSlug(dlta.Find("h2 > a").AttrOr("href", "")),
		})
	})

	animeTerbaru.Page = page
	animeTerbaru.Anime = animes

	return animeTerbaru, nil
}
