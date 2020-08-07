package system

import (
	"fmt"

	"github.com/alochym01/go-exporter-sample/config"
	"github.com/prometheus/client_golang/prometheus"
)

// Health is a system health metric
type Health struct{}

// Describe return a description of metrics
func (s Health) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_health
}

// Collect return a metric with all desc value and metric value
func (s Health) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service

	service, _ := metric.Systems()
	for _, v := range service {
		ch <- prometheus.MustNewConstMetric(config.S_health, prometheus.GaugeValue, float64(0),
			"",
			v.BIOSVersion,
			"",
			v.Description,
			"",
			v.PartNumber,
			"",
			fmt.Sprintf("%v", v.PowerState),
			v.SerialNumber,
			v.Model,
			v.SKU,
		)
	}
}
