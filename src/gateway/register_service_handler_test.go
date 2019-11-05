package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRegisterServiceHandler(t *testing.T) {
	t.Run("PUT /service/test", func(c *testing.T) {
		c.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "success!")
		}))
		defer ts.Close()
		router := http.NewServeMux()

		data := url.Values{}
		data.Add("address", ts.URL)

		w1 := httptest.NewRecorder()
		req1 := httptest.NewRequest("PUT", "/service/test", strings.NewReader(data.Encode()))
		req1.PostForm = data
		req1.Header.Add("content-type", "application/x-www-form-urlencoded")

		registerServiceHandler(router).ServeHTTP(w1, req1)

		res1 := w1.Result()
		if res1.StatusCode != 200 {
			fmt.Println("Expected HTTP success response!")
			fmt.Printf("Actual %d", res1.StatusCode)
			fmt.Println("")
			c.Fail()
		}

		w2 := httptest.NewRecorder()
		req2 := httptest.NewRequest("GET", "/test/", nil)
		router.ServeHTTP(w2, req2)

		res2 := w2.Result()
		body, _ := ioutil.ReadAll(res2.Body)
		if res2.StatusCode != 200 || string(body) != "success!" {
			fmt.Println("Expected 200, success!")
			fmt.Printf("Actual %d, %s", res2.StatusCode, string(body))
			fmt.Println("")
			c.Fail()
		}
		return
	})
}
