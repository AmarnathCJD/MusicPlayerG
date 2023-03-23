package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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

func MapHeaders(headers map[string]string) http.Header {
	var h = make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return h
}
