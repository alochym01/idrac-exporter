package system

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/stmcginnis/gofish/redfish"
)

func Health(s *redfish.ComputerSystem) prometheus.Gauge {
	health := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_health",
		Help: "Current System Health.",
	})

	status := map[string]float64{"OK": 1.0}

	if v, ok := status[string(s.Status.Health)]; ok {

		health.Set(v)
		return health
	}

	health.Set(0)
	return health
}
