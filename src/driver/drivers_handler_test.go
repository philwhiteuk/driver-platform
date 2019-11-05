package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestDriversHandler(t *testing.T) {
	var testDrivers []driver
	testDrivers = append(
		testDrivers,
		driver{ID: 34, Name: "Phil White"},
		driver{ID: 1, Name: "Colin McRae"},
		driver{ID: 88, Name: "Ayrton Senna"},
	)
	driversHandler := driversHandlerFunc(testDrivers)

	t.Run("GET /drivers/1", func(c *testing.T) {
		c.Parallel()

		req := httptest.NewRequest("GET", "/drivers/1", nil)
		w := httptest.NewRecorder()
		driversHandler.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != 200 {
			fmt.Println("Expected HTTP success response")
			fmt.Printf("Actual %d", resp.StatusCode)
			fmt.Println("")
			c.Fail()
		}

		body, _ := ioutil.ReadAll(resp.Body)
		actual := string(body)
		expected := "{\"id\":1,\"name\":\"Colin McRae\"}"
		if actual != expected {
			fmt.Printf("Expected %s\n", expected)
			fmt.Printf("Actual %s", actual)
			fmt.Println("")
			c.Fail()
		}

		return
	})
}
