package modules

import (
	"main/util"
	"strings"
)

const globalTopURL = "https://chartex.com/api/charts/spotify/top_200_daily/?pageSize=200&ordering=&country="

// GlobalTopSong is a struct that contains all the information about a song that is returned by the chartex API
type globalTopSong struct {
	ID                int    `json:"id"`
	SongTitle         string `json:"song_title"`
	RowOrder          int    `json:"row_order"`
	SpotifyPlatformID string `json:"spotify_platform_id"`
	TotalViews        int    `json:"total_views"`
	YoutubePlatformID string `json:"youtube_platform_id"`
	Artist            string `json:"artist"`
	Thumbnail         string `json:"thumbnail"`
}

// GlobalTop is a struct that contains all the information about Top 200 songs that is returned by the chartex API
type globalTop struct {
	Results []globalTopSong `json:"results"`
}

// GetSongName returns the song name of the song, it is the first part of the song title
func (g globalTopSong) GetSongName() string {
	if strings.Contains(g.SongTitle, " - ") {
		splita := strings.Split(g.SongTitle, " - ")
		return splita[0]
	} else if strings.Contains(g.SongTitle, ", ") {
		splita := strings.Split(g.SongTitle, ", ")
		return splita[0]
	} else if strings.Contains(g.SongTitle, " by ") {
		splita := strings.Split(g.SongTitle, " by ")
		return splita[0]
	}
	return g.SongTitle
}

// GetArtistName returns the artist name of the song, it is the last part of the song title
func (g globalTopSong) GetArtistName() string {
	if strings.Contains(g.SongTitle, " - ") {
		splita := strings.Split(g.SongTitle, " - ")
		return splita[len(splita)-1]
	} else if strings.Contains(g.SongTitle, ", ") {
		splita := strings.Split(g.SongTitle, ", ")
		return splita[len(splita)-1]
	} else if strings.Contains(g.SongTitle, " by ") {
		splita := strings.Split(g.SongTitle, " by ")
		return splita[len(splita)-1]
	}
	return ""
}

// GetThumbnail returns the thumbnail of the song from youtube
func (g globalTopSong) GetThumbnail() string {
	return "https://i1.ytimg.com/vi/" + g.YoutubePlatformID + "/hqdefault.jpg"
}

// GetGlobalTop gets the global top 200 songs from the Kworb API
func getGlobalTop() (*globalTop, error) {
	var globalTop globalTop
	headers := map[string]string{
		"Authorization": "Token 3fc2a8c4624b8f6ff94ee3ca5b8ba9fd335024d2f3ee76e3a812aed3a0c55690",
	}
	if err := util.GetJSON(globalTopURL, &globalTop, util.MapHeaders(headers)); err != nil {
		return nil, err
	}

	return &globalTop, nil
}

// ParseGlobalTop parses the global top 200 songs from the Kworb API
func GetTop(html ...bool) ([]globalTopSong, string, error) {
	var songs []globalTopSong
	globalT, err := getGlobalTop()
	if err != nil {
		return nil, "", err
	}
	for _, song := range globalT.Results {
		if len(songs) == 20 {
			break
		}
		song.Artist = song.GetArtistName()
		song.Thumbnail = song.GetThumbnail()
		songs = append(songs, song)
	}
	if len(html) > 0 && html[0] {
		return nil, parseGlobalAsHTML(globalT), nil
	}
	return songs, "", nil
}

func parseGlobalAsHTML(globalT *globalTop) string {
	var html string
	for i, song := range globalT.Results {
		if i == 20 {
			break
		}
		html += `
		<a href="/song?id=` + song.SpotifyPlatformID + `" class="flex items-center p-2 hover:bg-gray-400 rounded-xs shadow-lg hover:shadow-xl transition duration-300 ease-in-out">
			<img src="` + song.GetThumbnail() + `" alt="` + song.GetSongName() + `" class="w-16 h-10">
			<div class="ml-4">
				<h1 class="text-lg">#` + util.Itoa(song.RowOrder) + " " + song.GetSongName() + `</h1>
				<h2 class="text-sm">` + song.GetArtistName() + `</h2>
			</div>
			<div class="ml-auto">
				<div class="flex items-center">
					<div class="text-gray-500 text-xs font-bold">` + util.ItoaF(song.TotalViews) + ` views</div>
				</div>
			</div>
		</a>
		`
	}
	return html
}
