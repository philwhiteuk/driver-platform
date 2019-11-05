package common

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewUnifiedLogger(t *testing.T) {
	t.Run("remote address", func(c *testing.T) {
		c.Parallel()

		tc := testConnection{}
		u, _ := url.Parse("udp://test:5140")
		NewUnifiedLogger(
			testEstablishConnectionFn(&tc),
			u,
			"testLog",
		)

		if tc.Name != "testLog" || tc.RemoteAddr != "test:5140" {
			fmt.Printf("Expected testLog, test:5140")
			fmt.Println("")
			fmt.Printf("Actual %s, %s", tc.Name, tc.RemoteAddr)
			fmt.Println("")
			c.Fail()
		}

		return
	})
}
