package server

import (
	"log"
	"main/modules"
	"main/util"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		search_query := r.URL.Query().Get("query")
		should_html_str := r.URL.Query().Get("html")
		var should_html bool
		if should_html_str != "" {
			should_html, _ = strconv.ParseBool(should_html_str)
		}
		if search_query == "" {
			util.WriteJson(w, map[string]string{"error": "No query specified"}, http.StatusBadRequest)
		} else {
			songs, html, err := modules.SearchSong(search_query, should_html)
			if err != nil {
				util.WriteJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			} else if len(songs) == 0 {
				util.WriteJson(w, map[string]string{"error": "No results found"}, http.StatusNotFound)
			} else {
				if should_html {
					w.WriteHeader(http.StatusOK) // 200
					w.Write([]byte(html))
				} else {
					util.WriteJson(w, songs, http.StatusOK)
				}
			}
		}
	})

	http.HandleFunc("/get_song", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			util.WriteJson(w, map[string]string{"error": "No id specified"}, http.StatusBadRequest)
			return
		}
		song, err := modules.DownloadLink(id)
		if err != nil {
			util.WriteJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			return
		}
		util.WriteJson(w, song, http.StatusOK)
	})

	http.HandleFunc("/get_top", func(w http.ResponseWriter, r *http.Request) {
		var As_html bool
		As_html_str := r.URL.Query().Get("html")
		if As_html_str != "" {
			As_html, _ = strconv.ParseBool(As_html_str)
		}
		songs, html, err := modules.GetTop(As_html)
		if err != nil {
			util.WriteJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		} else {
			if As_html {
				w.WriteHeader(http.StatusOK) // 200
				w.Write([]byte(html))
			} else {
				util.WriteJson(w, songs, http.StatusOK)
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/song", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/song.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Music Player -G - Server started on port", util.GetSpecifiedPort())
	go util.InterruptHandler()
}
