package server

import (
	"encoding/json"
	"log"
	"main/modules"
	"main/util"
	"net/http"
)

func init() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		html := r.URL.Query().Get("html")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if html == "true" {
			songs, html, err := modules.SearchSong(query, true)
			if err != nil {
				util.WriteJson(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if len(songs) == 0 {
				util.WriteJson(w, "No results found", http.StatusNotFound)
				return
			}
			w.Write([]byte(html))
			return
		} else {
			songs, _, err := modules.SearchSong(query, false)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if len(songs) == 0 {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(songs)
			return
		}
	})

	http.HandleFunc("/get_song", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			util.WriteJson(w, "No id provided", http.StatusBadRequest)
			return
		}
		song, err := modules.DownloadLink(id)
		if err != nil {
			util.WriteJson(w, err.Error(), http.StatusInternalServerError)
			return
		}
		util.WriteJson(w, song, http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/song", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/song.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Music Player -G started on port", util.GetSpecifiedPort())
	go util.InterruptHandler()
}
