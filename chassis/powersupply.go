package chassis

import (
	"fmt"

	"github.com/alochym01/go-exporter-sample/config"
	"github.com/prometheus/client_golang/prometheus"
)

// PowerSupplyConSumByEach is a Chassis Power Control metric
type PowerSupplyConSumByEach struct{}

// Describe return a description of metrics
func (s PowerSupplyConSumByEach) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_powerconsumedbyeach
}

// Collect return a metric with all desc value and metric value
func (s PowerSupplyConSumByEach) Collect(ch chan<- prometheus.Metric) {
	var (
		value float64
		ok    bool
	)
	metric := config.GOFISH.Service

	chass, _ := metric.Chassis()

	for _, v := range chass {
		powers, _ := v.Power()
		if powers != nil {
			for _, p := range powers.PowerSupplies {
				ch <- prometheus.MustNewConstMetric(config.C_powerconsumedbyeach, prometheus.GaugeValue, float64(p.PowerOutputWatts),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.MemberID),
					fmt.Sprintf("%v", p.Model),
					fmt.Sprintf("%v", ""),
				)

				if value, ok = config.Status[string(p.Status.Health)]; !ok {
					value = float64(1)
				}

				ch <- prometheus.MustNewConstMetric(config.C_powersupplystatus, prometheus.GaugeValue, float64(value),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.MemberID),
					fmt.Sprintf("%v", p.Model),
					fmt.Sprintf("%v", ""),
					fmt.Sprintf("%v", p.HotPluggable),
					fmt.Sprintf("%v", p.Status.Health),
					fmt.Sprintf("%v", ""),
					fmt.Sprintf("%v", p.SparePartNumber),
					fmt.Sprintf("%v", p.FirmwareVersion),
				)
			}
		}
	}
}
