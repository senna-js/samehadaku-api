package samehadakuanimeterbaru

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/radenrishwan/mcp-server-samehadaku/external"
)

func Fetch(page int) (AnimeTerbaru, error) {
	if page < 0 {
		return AnimeTerbaru{}, external.ErrNotFound
	}

	url := external.BASE_URL + "anime-terbaru/page/" + strconv.Itoa(page)

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return AnimeTerbaru{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return AnimeTerbaru{}, err
	}

	animeTerbaru := AnimeTerbaru{}

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
			Slug:       external.ExtractSlug(dlta.Find("h2 > a").AttrOr("href", "")),
		})
	})

	animeTerbaru.Page = page
	animeTerbaru.Anime = animes

	return animeTerbaru, nil
}
