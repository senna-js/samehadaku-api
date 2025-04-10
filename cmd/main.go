package main

import (
	"encoding/json"
	"net/http"

	samehadakudetail "github.com/radenrishwan/mcp-server-samehadaku/external/detail"
	samehadakuhome "github.com/radenrishwan/mcp-server-samehadaku/external/home"
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

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
