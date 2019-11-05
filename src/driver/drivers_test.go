package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDrivers(t *testing.T) {
	t.Run("42=Lynda Carr", func(c *testing.T) {
		c.Parallel()

		expected := "{\"id\":42,\"name\":\"Joy Hardy\"}"
		resp, err := http.Get("http://127.0.0.1:8080/drivers/42")
		if err != nil {
			fmt.Println(err)
			c.Fail()
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			c.Fail()
		}

		actual := string(body)

		if string(actual) != expected {
			fmt.Println(expected)
			fmt.Println("!=")
			fmt.Println(actual)
			c.Fail()
		}

		return
	})
}
