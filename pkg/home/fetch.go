package samehadakuhome

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/senna-js/samehadaku-api/utility"
)

type Home struct {
	BaseUrl string
}

// fetch `homepage`
func (self Home) Fetch() (HomeResult, error) {
	url := self.BaseUrl

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return HomeResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return HomeResult{}, err
	}

	// find `anime terbaru`
	animeTerbaru := []AnimeCard{}
	doc.Find("#main > div:nth-child(4) > div.post-show > ul > li").Each(func(i int, s *goquery.Selection) {
		thumb := s.Find("div.thumb > a > img")
		dlta := s.Find("div.dtla")

		animeTerbaru = append(animeTerbaru, AnimeCard{
			Thumbnail:  thumb.AttrOr("src", ""),
			Title:      dlta.Find("h2 > a").Text(),
			Episode:    dlta.Find("span:nth-child(2) > author").Text(),
			PostedBy:   dlta.Find("span:nth-child(3) > author").Text(),
			ReleasedOn: dlta.Find("span:nth-child(4) > author").Text(),
			Href:       dlta.Find("h2 > a").AttrOr("href", ""),
			Slug:       utility.ExtractSlug(dlta.Find("h2 > a").AttrOr("href", "")),
		})
	})

	// find `project movie samehadaku`
	projectMovie := []ProjectMovieCard{}
	doc.Find("#sidebar > div > div > div.widgetseries > ul > li").Each(func(i int, s *goquery.Selection) {
		imgSeries := s.Find("div.imgseries > a > img")
		lftInfo := s.Find("div.lftinfo")

		genres := []Genre{}
		lftInfo.Find("span:nth-child(2) > a").Each(func(i int, s *goquery.Selection) {
			genres = append(genres, Genre{
				Title: s.Text(),
				Href:  s.AttrOr("href", ""),
			})
		})

		projectMovie = append(projectMovie, ProjectMovieCard{
			Thumbnail:  imgSeries.AttrOr("src", ""),
			Title:      lftInfo.Find("h2 > a").Text(),
			Genres:     genres,
			ReleasedOn: lftInfo.Find("span:nth-child(3)").Text(),
			Href:       lftInfo.Find("h2 > a").AttrOr("href", ""),
			Slug:       utility.ExtractSlug(lftInfo.Find("h2 > a").AttrOr("href", "")),
		})
	})

	// TODO: find `donghua dan film`

	// find `batch anime`
	batchAnime := []AnimeCard{}
	doc.Find("#main > div:nth-child(7) > div.post-show > ul > li").Each(func(i int, s *goquery.Selection) {
		thumb := s.Find("div.thumb > a > img")
		dlta := s.Find("div.dtla")

		batchAnime = append(batchAnime, AnimeCard{
			Thumbnail:  thumb.AttrOr("src", ""),
			Title:      dlta.Find("h2 > a").Text(),
			Episode:    dlta.Find("span:nth-child(2) > author").Text(),
			PostedBy:   dlta.Find("span:nth-child(3) > author").Text(),
			ReleasedOn: dlta.Find("span:nth-child(4) > author").Text(),
			Href:       dlta.Find("h2 > a").AttrOr("href", ""),
			Slug:       utility.ExtractSlug(dlta.Find("h2 > a").AttrOr("href", "")),
		})
	})

	return HomeResult{
		AnimeTerbaru: animeTerbaru,
		ProjectMovie: projectMovie,
		// DonghuaDanFilm: donghuaDanFilm,
		BatchAnime: batchAnime,
	}, nil
}
