package chassis

import (
	"fmt"

	"github.com/alochym01/go-exporter-sample/config"
	"github.com/prometheus/client_golang/prometheus"
)

// PowerControl is a Chassis Power Control metric
type PowerControl struct{}

// Describe return a description of metrics
func (s PowerControl) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_powercontrol
}

// Collect return a metric with all desc value and metric value
func (s PowerControl) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service

	chass, _ := metric.Chassis()

	for _, v := range chass {
		powers, _ := v.Power()
		if powers != nil {
			for _, p := range powers.PowerControl {
				ch <- prometheus.MustNewConstMetric(config.C_powercontrol, prometheus.GaugeValue, float64(0),
					fmt.Sprintf("%v", p.PowerMetrics.AverageConsumedWatts),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.MemberID),
					fmt.Sprintf("%v", p.PowerMetrics.IntervalInMin),
					fmt.Sprintf("%v", p.PowerMetrics.MaxConsumedWatts),
					fmt.Sprintf("%v", p.PowerMetrics.MinConsumedWatts),
				)

				ch <- prometheus.MustNewConstMetric(config.C_powerconsumedbyall, prometheus.GaugeValue, float64(p.PowerConsumedWatts),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.MemberID),
				)
			}

		}
	}
}
