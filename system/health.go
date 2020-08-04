package system

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stmcginnis/gofish/redfish"
)

var (
	status = map[string]float64{"OK": 0.0}
	value  float64
	ok     bool
)

func Health(s *redfish.ComputerSystem) prometheus.GaugeFunc {
	if value, ok = status[string(s.Status.Health)]; !ok {
		value = float64(1)
	}

	health := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "system_health",
			Help: fmt.Sprintf("idrac_system_health. {0: OK, 1: Warning, 2: Critical}"),
			ConstLabels: prometheus.Labels{
				"bios_backup":                      "",
				"bios_version":                     s.BIOSVersion,
				"intelligent_provisioning_version": "",
				"os_description":                   "",
				"os_name":                          "",
				"pca_part_number":                  s.PartNumber,
				"post_state":                       "",
				"power_state":                      string(s.PowerState),
				"serial_number":                    s.SerialNumber,
				"server_model":                     fmt.Sprintf("%v %v", s.Manufacturer, s.Model),
				"sku":                              s.SKU,
			},
		},
		func() float64 { return value },
	)

	return health
}
