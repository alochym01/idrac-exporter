package chassis

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stmcginnis/gofish/redfish"
)

var (
	status = map[string]float64{"OK": 0.0}
)

func Powers(s *redfish.Chassis) []prometheus.GaugeFunc {
	var result []prometheus.GaugeFunc
	// var (
	// 	value float64
	// ok    bool
	// )
	// if value, ok = status[string(p.Status.Health)]; !ok {
	// 	value = float64(1)
	// }
	powers, ok := s.Power()
	if ok != nil {
		return result
	}
	if powers != nil {
		for _, p := range powers.PowerControl {
			// Metric control
			control := prometheus.NewGaugeFunc(
				prometheus.GaugeOpts{
					Name: "idrac_power_control",
					Help: "Power Control",
					ConstLabels: prometheus.Labels{
						"average_consumed": fmt.Sprintf("%v", p.PowerMetrics.AverageConsumedWatts),
						"capacity":         fmt.Sprintf("%v", p.PowerCapacityWatts),
						"id":               fmt.Sprintf("%v", p.MemberID),
						"interval_in_min":  fmt.Sprintf("%v", p.PowerMetrics.IntervalInMin),
						"max_consumed":     fmt.Sprintf("%v", p.PowerMetrics.MaxConsumedWatts),
						"min_consumed":     fmt.Sprintf("%v", p.PowerMetrics.MaxConsumedWatts),
					},
				},
				func() float64 { return float64(0) },
			)

			// Metric consumed_by_all
			consumed_by_all := prometheus.NewGaugeFunc(
				prometheus.GaugeOpts{
					Name: "idrac_power_consumed_by_all",
					Help: "Power Consumed By All",
					ConstLabels: prometheus.Labels{
						"capacity": fmt.Sprintf("%v", p.PowerCapacityWatts),
						"id":       fmt.Sprintf("%v", p.MemberID),
					},
				},
				func() float64 { return float64(p.PowerConsumedWatts) },
			)
			result = append(result, control, consumed_by_all)
		}

		for _, p := range powers.PowerSupplies {
			// Metric consumed_by_each
			consumed_by_each := prometheus.NewGaugeFunc(
				prometheus.GaugeOpts{
					Name: "idrac_power_consumed_by_each",
					Help: "Power Consumed By Each",
					ConstLabels: prometheus.Labels{
						"id":       fmt.Sprintf("%v", p.MemberID),
						"capacity": fmt.Sprintf("%v", p.PowerCapacityWatts),
						"model":    fmt.Sprintf("%v", p.Model),
						"location": fmt.Sprintf("%v", p.Manufacturer),
					},
				},
				func() float64 { return float64(p.LastPowerOutputWatts) },
			)
			result = append(result, consumed_by_each)
		}
	}
	return result
}

// supply_status := prometheus.NewGaugeFunc(
// 	prometheus.GaugeOpts{
// 		Name: "idrac_power_supply_status",
// 		Help: "idrac_power_supply_status {0: OK, 1: Warning, 2: Critical}",
// 		ConstLabels: prometheus.Labels{
// 			"capacity":"500W",
// 			"firmware_version":"1.03",
// 			"hot_plug_capable":"True",
// 			"id":"1",
// 			"location":"",
// 			"max_output_watts_10s_interval":"79",
// 			"model":"865408-B21",
// 			"power_supply_status":"Ok",
// 			"spare_part_number":"866729-001",
// 		},
// 	},
// 	func() float64 { return 1 },
// )

// supply_redundancy_status := prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "idrac_power_supply_redundancy_status",
// 	Help: "Power Supply Redundancy Status",
// })

// supplies_status := prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "idrac_power_supplies_status",
// 	Help: "Power Supplies Status",
// 	[]string{"hostname"}
// })

// line_input_voltage := prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "idrac_power_line_input_voltage",
// 	Help: "Power Line Input Voltage",
// })

// status := map[string]float64{"OK": 1.0}

// if v, ok := status[string(s.Status.Health)]; ok {

// 	health.Set(v)
// 	return health
// }

// 	return [control]
// }
