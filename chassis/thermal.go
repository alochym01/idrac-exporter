package chassis

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

func Thermal(s *redfish.Chassis) prometheus.Gauge {
	fan_speed := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_fan_speed",
		Help: "Fan Speed",
	})

	fan_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_fan_status",
		Help: "Fan Status",
	})
	fan_redundancy_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_fan_redundancy_status",
		Help: "Fan Redundancy Status",
	})

	fans_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_fans_status",
		Help: "Metrics Fans Status",
	})

	temperature_reading := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_temperature_reading",
		Help: "Temperature Reading",
	})

	temperature_reading_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_temperature_reading_status",
		Help: "Temperature Reading Status",
	})

	temperatures_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_temperatures_status",
		Help: "Temperatures Status",
	})

}
