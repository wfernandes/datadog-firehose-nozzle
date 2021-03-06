package datadogclient

import (
	"fmt"

	"github.com/DataDog/datadog-firehose-nozzle/metrics"
	"github.com/cloudfoundry/sonde-go/events"
)

func (c *Client) ParseAppMetric(envelope *events.Envelope) ([]metrics.MetricPackage, error) {
	metricsPackages := []metrics.MetricPackage{}
	var err error

	if c.appMetrics == nil {
		return metricsPackages, fmt.Errorf("app metrics are not configured")
	}

	if envelope.GetEventType() != events.Envelope_ContainerMetric {
		return metricsPackages, fmt.Errorf("not an app metric")
	}

	metricsPackages, err = c.appMetrics.ParseAppMetric(envelope)

	return metricsPackages, err
}
