package modules

import (
	"errors"
	"main/util"
)

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
