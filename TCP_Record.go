/*package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	//"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*func TCPRecord() {
	output, err := exec.Command("bash", "-c", "netstat -st | grep segments").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(output))
}*/
/*
var (
	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "Saeed",
			Name:      "TCP_retransmission_record",
			Help:      "Time series database of retransmitted TCP packets",
		})
)

func main() {

	prometheus.MustRegister(gauge)
	/*retransRecord := promauto.NewCounter(prometheus.CounterOpts{
		Name: "Retransmitted_packets_record",
		Help: "The total number of retransmitted packets",
	})*/

/*	registry := prometheus.NewRegistry()

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		retransRecord,
	)*/
/*	go func() {
		for {
			output, err := exec.Command("bash", "-c", "netstat -st | grep segments").Output()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(output))
		}
	}()

	/*http.Handle(
		"/metrics", promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			}),
	)*/
/*	log.Fatalln(http.ListenAndServe(":9100", nil))
}
*/

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	output, err := exec.Command("bash", "-c", "netstat -st | grep segments").Output()
	if err != nil {
		log.Fatal(err)
	}
	num_of_packets := strings.Split(string(output), " ")
	i := 0
	for i < len(num_of_packets) {
		fmt.Printf("The %d element is: %s\n", i, num_of_packets[i])
		i++
	}
	//fmt.Println(len(num_of_packets))
	//fmt.Println(num_of_packets[4])
}
