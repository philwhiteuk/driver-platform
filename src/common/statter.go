package common

import (
	"fmt"
	"strings"
)

// Statter interface
type Statter interface {
	Count(metric string, value float64, tags []string) error
	Gauge(metric string, value float64, tags []string) error
}

// StatsDConnection
type statsDConnection interface {
	Write(b []byte) (n int, err error)
}

// NewDogStatsDStatter returns a statter conforming with DogStatsD protocol
func NewDogStatsDStatter(n string, c statsDConnection) (Statter, error) {
	return dogStatsD{name: n, conn: c}, nil
}

type dogStatsD struct {
	name string
	conn statsDConnection
}

type statsDMetricType string

const (
	count statsDMetricType = "c"
	gauge statsDMetricType = "g"
)

func (d dogStatsD) send(t statsDMetricType, m string, v float64, tags []string) error {
	dgram := fmt.Sprintf("%s.%s:%.1f|%s", d.name, m, v, t)

	var dgramWTags string
	if len(tags) > 0 {
		dgramWTags = fmt.Sprintf("%s|#%s", dgram, strings.Join(tags, ","))
	} else {
		dgramWTags = dgram
	}

	_, err := d.conn.Write([]byte(dgramWTags))
	if err != nil {
		return fmt.Errorf("Failed to write to connection %s", err)
	}

	return nil
}

func (d dogStatsD) Count(metric string, value float64, tags []string) error {
	return d.send(count, metric, value, tags)
}

func (d dogStatsD) Gauge(metric string, value float64, tags []string) error {
	return d.send(gauge, metric, value, tags)
}

type testStatsDConnection struct {
	metrics []string
}

func (c *testStatsDConnection) Write(b []byte) (n int, err error) {
	c.metrics = append(c.metrics, string(b))
	return 0, nil
}
