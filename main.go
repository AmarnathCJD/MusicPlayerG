package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url_prefix := "https://rabbitstream.net/"
		total_url := url_prefix + r.URL.Path[1:]
		log.Println(total_url)
		req, err := http.NewRequest(r.Method, total_url, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		req.Header = r.Header
		req.Header.Set("Host", "rabbitstream.net")
		req.Header.Set("Origin", "https://rabbitstream.net")
		req.Header.Set("Referer", "https://rabbitstream.net/")
		req.Header.Set("Authority", "rabbitstream.net")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)
		w.Write(ReadBytes(resp.Body))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ReadBytes(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}

// Copyright (c) 2023 by RoseLoverX. All rights reserved.
