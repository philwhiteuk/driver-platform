package main

import (
	"fmt"
	"testing"
)

func TestParseDriverData(t *testing.T) {

	actual, _ := parseDriverData()

	t.Run("Darryl Burke", func(c *testing.T) {
		c.Parallel()

		if actual[0].ID != 1 || actual[0].Name != "Darryl Burke" {
			fmt.Println("Expected 1, Darryl Burke")
			fmt.Printf("Actual %d, %s", actual[0].ID, actual[0].Name)
			fmt.Println("")
			c.Fail()
		}

		return
	})

	t.Run("Marianne Colon", func(c *testing.T) {
		c.Parallel()

		if actual[5].ID != 6 || actual[5].Name != "Marianne Colon" {
			fmt.Println("Expected 6, Marianne Colon")
			fmt.Printf("Actual %d, %s", actual[5].ID, actual[5].Name)
			fmt.Println("")
			c.Fail()
		}

		return
	})
}
