package samehadakujadwalrilis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/radenrishwan/samehadaku-api/external"
)

func Fetch(day string) ([]JadwalRilis, error) {

	// https://samehadaku.mba/wp-json/custom/v1/all-schedule?perpage=20&day=wednesday&type=schtml
	url := external.BASE_URL + fmt.Sprintf("wp-json/custom/v1/all-schedule?perpage=20&day=%s&type=schtml", day)

	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return []JadwalRilis{}, err
	}

	if resp.StatusCode != 200 {
		return []JadwalRilis{}, external.ErrNotFound
	}

	body, _ := io.ReadAll(resp.Body)

	var response []animeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return []JadwalRilis{}, err
	}

	jadwalRilis := []JadwalRilis{}
	for _, anime := range response {
		genres := strings.Split(anime.Genre, ", ")

		jadwalRilis = append(jadwalRilis, JadwalRilis{
			Thumbnail: anime.FeaturedImg,
			Title:     anime.Title,
			Type:      anime.Type,
			Rating:    anime.EastScore,
			Time:      anime.EastTime,
			Genre:     genres,
			Href:      anime.URL,
			Slug:      external.ExtractSlug(anime.URL),
		})
	}

	return jadwalRilis, nil
}
