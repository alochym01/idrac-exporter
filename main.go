package main

import (
	"fmt"
	"net/http"

	"github.com/alochym01/go-exporter-sample/chassis"
	"github.com/alochym01/go-exporter-sample/config"
	"github.com/alochym01/go-exporter-sample/system"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stmcginnis/gofish"
)

func metrichandler(w http.ResponseWriter, r *http.Request) {
	var err error
	conf := gofish.ClientConfig{
		Endpoint: r.URL.Query().Get("ilo_host"),
		Username: config.Idracuser,
		Password: config.Idracpassword,
		Insecure: true,
	}
	fmt.Println(r.URL.Query().Get("ilo_host"))
	config.GOFISH, err = gofish.Connect(conf)
	if err != nil {
		panic(err)
	}
	defer config.GOFISH.Logout()

	fmt.Println(" Connect successfull")

	mhandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
	})
	mhandler.ServeHTTP(w, r)

}
func main() {

	// Prometheus Metrics Registration
	health := system.Health{}
	prometheus.MustRegister(health)
	powercontrol := chassis.PowerControl{}
	prometheus.MustRegister(powercontrol)
	powerconsumbyeach := chassis.PowerSupplyConSumByEach{}
	prometheus.MustRegister(powerconsumbyeach)

	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe("192.168.2.143:9519", nil)
}
