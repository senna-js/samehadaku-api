package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	samehadakuanimeterbaru "github.com/radenrishwan/samehadaku-api/external/pkg/anime-terbaru"
	samehadakudaftaranime "github.com/radenrishwan/samehadaku-api/external/pkg/daftar-anime"
	samehadakudetail "github.com/radenrishwan/samehadaku-api/external/pkg/detail"
	samehadakuepisode "github.com/radenrishwan/samehadaku-api/external/pkg/episode"
	samehadakuhome "github.com/radenrishwan/samehadaku-api/external/pkg/home"
	samehadakujadwalrilis "github.com/radenrishwan/samehadaku-api/external/pkg/jadwal-rilis"
	samehadakusearch "github.com/radenrishwan/samehadaku-api/external/pkg/search"
)

func main() {
	mux := http.NewServeMux()
	prefix := "/api/v1"

	mux.HandleFunc("GET "+prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		result, err := samehadakuhome.Fetch()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, result)
	})

	mux.HandleFunc("GET "+prefix+"/anime/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		if slug == "" {
			http.Error(w, "slug is required", http.StatusBadRequest)
			return
		}

		result, err := samehadakudetail.Fetch(slug)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, result)
	})

	mux.HandleFunc("GET "+prefix+"/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		if slug == "" {
			http.Error(w, "slug is required", http.StatusBadRequest)
			return
		}

		streamUrls := r.URL.Query().Get("stream_urls")
		var response samehadakuepisode.Episode
		if streamUrls == "1" {
			result, err := samehadakuepisode.Fetch(slug, true)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response = result
		} else {
			result, err := samehadakuepisode.Fetch(slug, false)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response = result
		}

		WriteJSON(w, response)
	})

	mux.HandleFunc("GET "+prefix+"/anime-terbaru/{page}", func(w http.ResponseWriter, r *http.Request) {
		page := r.PathValue("page")
		if page == "" {
			http.Error(w, "page is required", http.StatusBadRequest)
			return
		}

		p, err := strconv.Atoi(page)
		if err != nil {
			p = -1
		}

		response, err := samehadakuanimeterbaru.Fetch(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, response)
	})

	mux.HandleFunc("GET "+prefix+"/jadwal-rilis/{day}", func(w http.ResponseWriter, r *http.Request) {
		day := r.PathValue("day")
		if day == "" {
			http.Error(w, "page is required", http.StatusBadRequest)
			return
		}

		response, err := samehadakujadwalrilis.Fetch(day)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, response)
	})

	mux.HandleFunc("GET "+prefix+"/daftar-anime", func(w http.ResponseWriter, r *http.Request) {
		seperate := r.URL.Query().Get("seperate")

		p, err := strconv.Atoi(seperate)
		if err != nil {
			p = -1
		}

		isSeperate := true
		if p != 1 {
			isSeperate = false
		}

		response, err := samehadakudaftaranime.Fetch(isSeperate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, response)
	})

	mux.HandleFunc("GET "+prefix+"/search-anime", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "page is required", http.StatusBadRequest)
			return
		}

		response, err := samehadakusearch.Fetch(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, response)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
