package modules

import (
	"main/util"
	"net/url"
)

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

const searchEndpoint = "https://a.ztorr.me/api/spotify?q="

func SearchSong(song string, is_H bool) (songs []song, html string, err error) {
	err = util.GetJSON(searchEndpoint+url.QueryEscape(song), &songs)
	if is_H {
		html = parseAsHTML(songs)
	}
	return
}

func parseAsHTML(songs []song) string {
	var html = ""
	for _, song := range songs {
		html += `
		<a href="/song?id=` + song.ID + `" class="flex items-center p-2 hover:bg-gray-900 rounded-lg shadow-lg hover:shadow-xl transition duration-300 ease-in-out">
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
