package common

import (
	"fmt"
	"net/http"
	"strconv"
)

// HTTPLogger logger middleware for http handlers
func HTTPLogger(l Logger, s Statter, c Clock) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t0 := c.Now()
			s.Count("http.request", 1, nil)
			resp := newHTTPResponseRecorder(w)
			defer func() {
				t1 := c.Now()
				rs := strconv.Quote(fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto))
				l.Info(
					fmt.Sprintf(
						"%s - - [%s] %s %d %d",
						r.RemoteAddr,
						c.Now().Format("2/Jan/2006:15:04:05 -0700"),
						rs,
						resp.statusCode,
						resp.contentLength,
					),
				)
				s.Count("http.response", 1, []string{fmt.Sprintf("status_code:%sxx", strconv.Itoa(resp.statusCode)[:1])})
				s.Gauge("http.response_time_ms", t1.Sub(t0).Seconds()*1000, nil)
			}()
			next.ServeHTTP(&resp, r)
		})
	}
}
