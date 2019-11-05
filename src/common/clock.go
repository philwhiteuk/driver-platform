package common

import (
	"fmt"
	"time"
)

// Clock definition
type Clock interface {
	Now() time.Time
}

// NewSystemClock returns a system clock which implements Clock
func NewSystemClock() Clock {
	return systemClock{}
}

type systemClock struct{}

func (s systemClock) Now() time.Time {
	return time.Now()
}

type testClock struct {
	now string
}

func (t *testClock) Now() time.Time {
	fmt.Println(t.now)
	fixedTime, _ := time.Parse("2006-01-02T15:04:05Z", t.now)
	return fixedTime
}

func (t *testClock) advance(secs int) {
	fixedTime, _ := time.Parse("2006-01-02T15:04:05Z", t.now)
	t.now = fixedTime.Add(time.Second * time.Duration(secs)).Format("2006-01-02T15:04:05Z")
	fmt.Println(t.now)
}
