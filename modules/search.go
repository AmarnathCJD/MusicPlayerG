package modules

import (
	"main/util"
	"net/url"
)

// Song is a struct that contains all the information about a song that is returned by the Spotify API
type song struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Image string `json:"image"`
	URI   string `json:"uri"`
	Album struct {
		Name string `json:"name"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"artists"`
}

// PrivateSearchEndpoint hosted on vercel
const searchEndpoint = "https://a.ztorr.me/api/spotify?q="

// SearchSong searches for a song on Spotify and returns the results
func SearchSong(song string, is_H bool) (songs []song, html string, err error) {
	err = util.GetJSON(searchEndpoint+url.QueryEscape(song), &songs)
	if is_H {
		html = parseAsHTML(songs)
	}
	return
}

// parseAsHTML parses the songs into HTML for the frontend
func parseAsHTML(songs []song) string {
	var html = ""
	for _, song := range songs {
		html += `
		<a href="/song?id=` + song.ID + `" class="flex items-center p-2 hover:bg-gray-400 rounded-xs shadow-lg hover:shadow-xl transition duration-300 ease-in-out">
			<img src="` + song.Image + `" alt="" class="w-16 h-16 rounded-lg">
			<div class="ml-4">
				<h1 class="text-lg">` + song.Name + `</h1>
				<h2 class="text-sm">` + song.Artists[0].Name + `</h2>
			</div>
		</a>
		`
	}
	return html
}
