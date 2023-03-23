package modules

import (
	"errors"
	"main/util"
)

// Track is a struct that contains the metadata of a track and the download link
type Track struct {
	Success  bool `json:"success"`
	Metadata struct {
		ID      string `json:"id"`
		Artists string `json:"artists"`
		Title   string `json:"title"`
		Cover   string `json:"cover"`
		Album   string `json:"album"`
	} `json:"metadata"`
	Link string `json:"link"`
}

// DownloadLink returns the download link of a track given the track URI from Spotify
func DownloadLink(trackURI string) (*Track, error) {
	headers := map[string]string{
		"origin":  "https://spotifydown.com",
		"referer": "https://spotifydown.com/",
	}
	var track Track
	err := util.GetJSON("https://api.spotifydown.com/download/"+trackURI, &track, util.MapHeaders(headers))
	if err != nil {
		return nil, errors.New("failed to get track")
	}
	if !track.Success {
		return nil, errors.New("failed to get track")
	}
	return &track, nil
}
