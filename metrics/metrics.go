package metrics

import (
	"errors"
	"fmt"

	"github.com/cloudfoundry/sonde-go/events"
)

type Point struct {
	Timestamp int64
	Value     float64
}

type MetricKey struct {
	EventType events.Envelope_EventType
	Name      string
	TagsHash  string
}

type MetricValue struct {
	Tags   []string
	Points []Point
	Host   string
}

type MetricPackage struct {
	MetricKey   *MetricKey
	MetricValue *MetricValue
}

type Series struct {
	Metric string   `json:"metric"`
	Points []Point  `json:"points"`
	Type   string   `json:"type"`
	Host   string   `json:"host,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

type Metrics Series

func (p Point) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`[%d, %f]`, p.Timestamp, p.Value)), nil
}

func (p *Point) UnmarshalJSON(in []byte) error {
	var timestamp int64
	var value float64

	parsed, err := fmt.Sscanf(string(in), `[%d,%f]`, &timestamp, &value)
	if err != nil {
		return err
	}
	if parsed != 2 {
		return errors.New("expected two parsed values")
	}

	p.Timestamp = timestamp
	p.Value = value

	return nil
}
