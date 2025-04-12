package samehadakuapi

import (
	samehadakuanimeterbaru "github.com/radenrishwan/samehadaku-api/pkg/anime-terbaru"
	samehadakudaftaranime "github.com/radenrishwan/samehadaku-api/pkg/daftar-anime"
	samehadakudetail "github.com/radenrishwan/samehadaku-api/pkg/detail"
	samehadakuepisode "github.com/radenrishwan/samehadaku-api/pkg/episode"
	samehadakuhome "github.com/radenrishwan/samehadaku-api/pkg/home"
	samehadakujadwalrilis "github.com/radenrishwan/samehadaku-api/pkg/jadwal-rilis"
	samehadakusearch "github.com/radenrishwan/samehadaku-api/pkg/search"
	"github.com/radenrishwan/samehadaku-api/utility"
)

type Samehadaku struct {
	BaseUrl string
	samehadakuanimeterbaru.AnimeTerbaru
	samehadakudaftaranime.DaftarAnime
	samehadakudetail.Detail
	samehadakuepisode.Episode
	samehadakuhome.Home
	samehadakujadwalrilis.JadwalRilis
	samehadakusearch.Search
}

func NewSamehadaku(baseUrl string) *Samehadaku {
	if baseUrl == "" {
		baseUrl = utility.BASE_URL
	}

	return &Samehadaku{
		BaseUrl: baseUrl,
	}
}
