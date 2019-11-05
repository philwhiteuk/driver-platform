package common

import (
	"fmt"
	"testing"
)

func TestNewDogStatsDStatter(t *testing.T) {
	t.Run("count", func(c *testing.T) {
		c.Parallel()

		tc := testStatsDConnection{}
		s, _ := NewDogStatsDStatter("testApp", &tc)

		var tags []string
		s.Count("apples", 2, tags)

		if tc.metrics[0] != "testApp.apples:2.0|c" {
			fmt.Printf("Expected testApp.apples:2|c")
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[0])
			fmt.Println("")
			c.Fail()
		}

		return
	})

	t.Run("gauge", func(c *testing.T) {
		c.Parallel()

		tc := testStatsDConnection{}
		s, _ := NewDogStatsDStatter("testApp", &tc)

		var tags []string
		s.Gauge("pears", 1.9, tags)

		if tc.metrics[0] != "testApp.pears:1.9|g" {
			fmt.Printf("Expected testApp.pears:1.9|g")
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[0])
			fmt.Println("")
			c.Fail()
		}

		return
	})

	t.Run("tags", func(c *testing.T) {
		c.Parallel()

		tc := testStatsDConnection{}
		s, _ := NewDogStatsDStatter("testApp", &tc)

		var tags []string
		tags = append(tags, "tag1", "tag2:orange")
		s.Count("pears", 13, tags)

		if tc.metrics[0] != "testApp.pears:13.0|c|#tag1,tag2:orange" {
			fmt.Printf("Expected testApp.pears:13.0|c|#tag1,tag2:orange")
			fmt.Println("")
			fmt.Printf("Actual %s", tc.metrics[0])
			fmt.Println("")
			c.Fail()
		}

		return
	})
}
