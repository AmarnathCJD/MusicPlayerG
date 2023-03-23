package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GetJSON gets a json response from a url and decodes it into the target
func GetJSON(url string, target interface{}, headers ...http.Header) error {
	var rreq, _ = http.NewRequest("GET", url, nil)
	if len(headers) > 0 {
		rreq.Header = headers[0]
	}
	r, err := http.DefaultClient.Do(rreq)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// PostJSON posts a json body to a url and decodes the response into the target
func PostJSON(url string, body interface{}, target interface{}, headers ...http.Header) error {
	var b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	var rreq, _ = http.NewRequest("POST", url, b)
	if len(headers) > 0 {
		rreq.Header = headers[0]
	}
	r, err := http.DefaultClient.Do(rreq)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// WriteJson writes a json response to the response writer with the specified status code and intented json
func WriteJson(w http.ResponseWriter, data interface{}, status int, intent ...bool) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	}
	if len(intent) > 0 {
		json.NewEncoder(w).Encode(data)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// MapHeaders converts a map to a http.Header type
func MapHeaders(headers map[string]string) http.Header {
	var h = make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return h
}
