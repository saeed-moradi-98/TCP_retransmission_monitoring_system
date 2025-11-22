package main

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	sent_seg    prometheus.Gauge
	retrans_seg prometheus.Gauge
}

var (
	sent_segments          int
	retransmitted_segments int
)

var m = &metrics{
	sent_seg: promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "all_transmitted_packets",
			Help: "Number of all transmitted packets.",
		}),
	retrans_seg: promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "all_retransmitted_packets",
			Help: "Number of all retransmitted packets.",
		}),
}

// This function registers the prometheus metrics.
func CustomRegisterer() *metrics {
	reg := prometheus.NewRegistry()
	reg.MustRegister(m.sent_seg)
	reg.MustRegister(m.retrans_seg)
	return m
}

// This function counts the number of all and retransmitted TCP packets through running a Linux command.
func NumOfPackets() (int, int) {
	output, err := exec.Command("bash", "-c", "netstat -st | grep segments").Output()

	if err != nil {
		log.Fatal(err)
	}
	num_of_packets := strings.Split(string(output), " ")
	sent_segments, err = strconv.Atoi(num_of_packets[10]) // Convert strings to integers
	retransmitted_segments, err = strconv.Atoi(num_of_packets[17])
	return sent_segments, retransmitted_segments
}

// This function handles the http requests sent to the exporter by ex
func custom_handler(w http.ResponseWriter, req *http.Request) {
	sent_segments, retransmitted_segments := NumOfPackets()
	m.retrans_seg.Set(float64(retransmitted_segments))
	m.sent_seg.Set(float64(sent_segments))
	promhttp.Handler().ServeHTTP(w, req)
}

func main() {
	CustomRegisterer()
	http.HandleFunc("/metrics", custom_handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
