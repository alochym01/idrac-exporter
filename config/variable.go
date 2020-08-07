package config

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish"
)

var (
	// GOFISH is global variable
	GOFISH *gofish.APIClient

	// Status map
	Status = map[string]float64{"OK": 0.0}

	// Idracuser info
	Idracuser = "root"
	// Idracpassword info
	Idracpassword = "calvin"

	// IDRAC Metric

	// S_health => System Health Metric
	S_health = prometheus.NewDesc(
		"idrac_system_health",
		"idrac_system_health {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"bios_backup",
			"bios_version",
			"intelligent_provisioning_version",
			"os_description",
			"os_name",
			"pca_part_number",
			"post_state",
			"power_state",
			"serial_number",
			"server_model",
			"sku",
		},
		nil,
	)

	// C_powercontrol => Chassis Power Control Metric
	C_powercontrol = prometheus.NewDesc(
		"idrac_power_control",
		"Power Control",
		[]string{
			"average_consumed",
			"capacity",
			"id",
			"interval_in_min",
			"max_consumed",
			"min_consumed",
		},
		nil,
	)

	// C_powerconsumedbyall => Chassis Power Control Metric
	C_powerconsumedbyall = prometheus.NewDesc(
		"idrac_power_consumed_by_all",
		"Power Consumed By All",
		[]string{
			"capacity",
			"id",
		},
		nil,
	)

	// C_powerconsumedbyeach => Chassis Power Control Metric
	C_powerconsumedbyeach = prometheus.NewDesc(
		"idrac_power_consumed_by_each",
		"Power Consumed By each",
		[]string{
			"capacity",
			"id",
			"model",
			"location",
		},
		nil,
	)

	// C_powersupplystatus => Chassis Power Control Metric
	C_powersupplystatus = prometheus.NewDesc(
		"idrac_power_supply_status",
		"Power Supply status {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"capacity",
			"id",
			"model",
			"location",
			"hot_plug_capable",
			"power_supply_status",
			"max_output_watts_10s_interval",
			"spare_part_number",
			"firmware_version",
		},
		nil,
	)
)
