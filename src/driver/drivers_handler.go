package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func driversHandlerFunc(d []driver) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		requested, err := strconv.Atoi(path[len(path)-1])
		if err != nil {
			w.WriteHeader(400)
		}

		for _, d := range d {
			if d.ID == requested {
				json, err := json.Marshal(d)
				if err != nil {
					w.WriteHeader(500)
				}
				io.WriteString(w, string(json))
				return
			}
		}

		w.WriteHeader(404)
	})
}
