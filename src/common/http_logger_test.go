package common

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpLogger(t *testing.T) {
	t.Run("common log format", func(c *testing.T) {
		c.Parallel()
		tl := testLogger{}
		tc := testStatsDConnection{}
		s, _ := NewDogStatsDStatter("testApp", &tc)
		fc := testClock{now: "2019-11-03T00:00:00Z"}

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/test/", nil)
		HTTPLogger(&tl, s, &fc)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fc.advance(3)
			w.WriteHeader(418)
			io.WriteString(w, "I'm a teapot")
		})).ServeHTTP(w, req)

		expectedLog := "Info: 192.0.2.1:1234 - - [3/Nov/2019:00:00:03 +0000] \"GET /test/ HTTP/1.1\" 418 12"
		if tl.Log != expectedLog {
			fmt.Printf("Expected %s", expectedLog)
			fmt.Println("")
			fmt.Printf("Actual %s", tl.Log)
			fmt.Println("")
			c.Fail()
		}

		expectedMetric := "testApp.http.request:1.0|c"
		if tc.metrics[0] != expectedMetric {
			fmt.Printf("Expected %s", expectedMetric)
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[0])
			fmt.Println("")
			c.Fail()
		}

		expectedMetric = "testApp.http.response:1.0|c|#status_code:4xx"
		if tc.metrics[1] != expectedMetric {
			fmt.Printf("Expected %s", expectedMetric)
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[1])
			fmt.Println("")
			c.Fail()
		}

		expectedMetric = "testApp.http.response_time_ms:3000.0|g"
		if tc.metrics[2] != expectedMetric {
			fmt.Printf("Expected %s", expectedMetric)
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[2])
			fmt.Println("")
			c.Fail()
		}

		return
	})
}
