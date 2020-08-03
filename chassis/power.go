package chassis

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/stmcginnis/gofish/redfish"
)

func Power(s *redfish.Chassis) prometheus.Gauge {
	control := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_control",
		Help: "Power Control",
	})

	supply_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_supply_status",
		Help: "Power Supply Status",
	})

	supply_redundancy_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_supply_redundancy_status",
		Help: "Power Supply Redundancy Status",
	})

	supplies_status := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_supplies_status",
		Help: "Power Supplies Status",
	})

	consumed_by_all := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_consumed_by_all",
		Help: "Power Consumed By All",
	})

	consumed_by_each := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_consumed_by_each",
		Help: "Power Consumed By Each",
	})

	line_input_voltage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "idrac_power_line_input_voltage",
		Help: "Power Line Input Voltage",
	})

	status := map[string]float64{"OK": 1.0}

	if v, ok := status[string(s.Status.Health)]; ok {

		health.Set(v)
		return health
	}

	health.Set(0)
	return health
}
