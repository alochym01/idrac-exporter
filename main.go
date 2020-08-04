package main

// idrac
// bios_or_hardware_health
// system_health
// power_supply
// fan
// temperature_reading_status
// storages
// array_controller
// physical_drive
// processors
// cpuutil
// cpu0power
// cpuicutil
// memory
// memorybusutil
// avgcpu1freq
// network_adapters
// network_port

import (
	"net/http"
	"os"

	"github.com/alochym01/idrac-exporter/chassis"
	"github.com/alochym01/idrac-exporter/system"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	"github.com/stmcginnis/gofish"
)

var log = logrus.New()

func main() {
	// request sample: GET http://192.168.2.143:9416/metrics?ilo_host=https://192.168.2.157
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe("192.168.2.143:9519", nil)
}

func metrichandler(w http.ResponseWriter, rsp *http.Request) {
	switch rsp.Method {
	case http.MethodGet:
		midrachost, _ := rsp.URL.Query()["ilo_host"]
		//redfish object create to connect with desire host
		idracuser := "alochym"
		idracpassword := "alochym"

		mredfishconfig := gofish.ClientConfig{
			Endpoint: midrachost[0],
			Username: idracuser,
			Password: idracpassword,
			Insecure: true,
		}
		mredfishconnect, err := gofish.Connect(mredfishconfig)
		if err != nil {
			panic(err)
		}
		defer mredfishconnect.Logout()

		service := mredfishconnect.Service

		sys, _ := service.Systems()
		chass, _ := service.Chassis()

		// Register System Health
		health := system.Health(sys[0])
		prometheus.Register(health)

		// Register Chassis
		for _, c := range chass {
			powers := chassis.Powers(c)
			for _, power := range powers {
				prometheus.Register(power)
			}
		}
		log.Out = os.Stdout

		// Set log level
		l, err := logrus.ParseLevel("debug")
		if err != nil {
			log.Fatal(err)
		}
		log.SetLevel(l)

		log.Infoln("Starting dellhw_exporter", version.Info())
		log.Infoln("Build context", version.BuildContext())

		mhandler := promhttp.HandlerFor(prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				ErrorLog:      log,
				ErrorHandling: promhttp.ContinueOnError,
			})
		mhandler.ServeHTTP(w, rsp)
	}
}
