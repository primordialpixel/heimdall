package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
)

type observation struct {
	IP string
	Language string
	UserAgent string
	Timestamp string
}

func main() {
	_observations := []observation{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_obs := observation{
			IP: r.RemoteAddr,
			Language: r.Header.Get("Accept-Language"),
			UserAgent: r.Header.Get("User-Agent"),
			Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		}

		_observations = append(_observations, _obs)

		b, err := json.MarshalIndent(_obs, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(b)
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(_observations, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
