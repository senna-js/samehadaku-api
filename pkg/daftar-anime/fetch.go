package samehadakudaftaranime

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/senna-js/samehadaku-api/utility"
)

type DaftarAnime struct {
	BaseUrl string
}

func (self DaftarAnime) Fetch(seperate bool) (DaftarAnimeResult, error) {
	url := self.BaseUrl + "/daftar-anime-2/?list"

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return DaftarAnimeResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return DaftarAnimeResult{}, err
	}

	daftarAnime := DaftarAnimeResult{}

	alphabets := []Alphabet{}
	doc.Find("#main > div.listpst > div").Each(func(i int, s *goquery.Selection) {
		alphabet := Alphabet{}
		alphabet.Alphabet = s.Find("div.listabj > a").Text()

		animes := []Anime{}
		s.Find("div.listttl > ul > li > a").Each(func(i int, s *goquery.Selection) {
			anime := Anime{}

			anime.Title = s.Text()
			anime.Href = s.AttrOr("href", "")
			anime.Slug = utility.ExtractSlug(anime.Href)

			animes = append(animes, anime)
		})

		alphabet.Animes = animes

		if alphabet.Alphabet != "" {
			alphabets = append(alphabets, alphabet)
		}
	})

	if seperate {
		daftarAnime.Alphabets = alphabets
	} else {
		animes := []Anime{}
		for _, alphabet := range alphabets {
			animes = append(animes, alphabet.Animes...)
		}

		daftarAnime.Animes = animes
	}

	return daftarAnime, nil
}
