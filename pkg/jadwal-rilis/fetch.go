package samehadakujadwalrilis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/radenrishwan/samehadaku-api/utility"
)

type JadwalRilis struct {
	BaseUrl string
}

func (self JadwalRilis) Fetch(day string) ([]JadwalRilisResult, error) {
	url := self.BaseUrl + fmt.Sprintf("wp-json/custom/v1/all-schedule?perpage=20&day=%s&type=schtml", day)

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return []JadwalRilisResult{}, err
	}

	if resp.StatusCode != 200 {
		return []JadwalRilisResult{}, utility.ErrNotFound
	}

	body, _ := io.ReadAll(resp.Body)

	var response []animeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return []JadwalRilisResult{}, err
	}

	jadwalRilis := []JadwalRilisResult{}
	for _, anime := range response {
		genres := strings.Split(anime.Genre, ", ")

		jadwalRilis = append(jadwalRilis, JadwalRilisResult{
			Thumbnail: anime.FeaturedImg,
			Title:     anime.Title,
			Type:      anime.Type,
			Rating:    anime.EastScore,
			Time:      anime.EastTime,
			Genre:     genres,
			Href:      anime.URL,
			Slug:      utility.ExtractSlug(anime.URL),
		})
	}

	return jadwalRilis, nil
}
