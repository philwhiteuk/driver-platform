package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func registerServiceHandler(router *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(404)
		}

		u, err := url.Parse(r.PostFormValue("address"))
		if err != nil {
			w.WriteHeader(400)
		}

		p := strings.Split(r.URL.Path, "/")
		service := p[len(p)-1]

		router.Handle(fmt.Sprintf("/%s/", service), httputil.NewSingleHostReverseProxy(u))

		w.WriteHeader(200)
	})
}
