package main

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	sent_seg    prometheus.Gauge
	retrans_seg prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		sent_seg: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "all_transmitted_packets",
				Help: "Number of all transmitted packets.",
			}),
		retrans_seg: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "retransmitted_packets",
				Help: "number of all retransmitted packets.",
			}),
	}
	reg.MustRegister(m.sent_seg)
	reg.MustRegister(m.retrans_seg)
	return m
}

var (
	sent_segments          int
	retransmitted_segments int
)

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)
	go func() {
		for {
			output, err := exec.Command("bash", "-c", "netstat -st | grep segments").Output()

			if err != nil {
				log.Fatal(err)
			}
			num_of_packets := strings.Split(string(output), " ")
			sent_segments, err = strconv.Atoi(num_of_packets[10])
			retransmitted_segments, err = strconv.Atoi(num_of_packets[17])
			m.sent_seg.Set(float64(sent_segments))
			m.retrans_seg.Set(float64(retransmitted_segments))
		}
	}()
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.ListenAndServe(":9090", nil)
}
