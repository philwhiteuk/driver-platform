package common

import "net/http"

func newHTTPResponseRecorder(w http.ResponseWriter) httpResponseRecorder {
	return httpResponseRecorder{w: w, contentLength: 0, statusCode: 200}
}

type httpResponseRecorder struct {
	w             http.ResponseWriter
	contentLength int
	statusCode    int
}

func (r *httpResponseRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *httpResponseRecorder) Write(b []byte) (int, error) {
	r.contentLength = len(b)
	return r.w.Write(b)
}

func (r *httpResponseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.w.WriteHeader(statusCode)
}
