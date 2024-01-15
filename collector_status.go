package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Ensure interface
var _ prometheus.Collector = &sabnzbdStatusCollector{}

type sabnzbdStatusCollector struct{}

func (s *sabnzbdStatusCollector) Collect(ch chan<- prometheus.Metric) {
	out := statusOutput{}
	err := call(&out, map[string]string{"mode": "status"})
	if err != nil {
		log.Println("error collecting history metrics: ", err)
	} else {
		configOut := configServersOutput{}
		err = call(&configOut, map[string]string{"mode": "get_config", "section": "servers"})
		if err != nil {
			log.Println("error getting servers config:", err)
		}
		for _, server := range out.Status.Servers {
			serverHost := server.Name
			for _, s := range configOut.Config.Servers {
				if server.Name == s.DisplayName {
					serverHost = s.Host
				}
			}
			active := 0
			if server.Active {
				active = 1
			}
			ch <- prometheus.MustNewConstMetric(statusServerActive, prometheus.GaugeValue, float64(active), serverHost)
			ch <- prometheus.MustNewConstMetric(statusServerActiveConn, prometheus.GaugeValue, float64(server.ActiveConn), serverHost)
			ch <- prometheus.MustNewConstMetric(statusServerTotalConn, prometheus.GaugeValue, float64(server.TotalConn), serverHost)
			ch <- prometheus.MustNewConstMetric(statusServerSSL, prometheus.GaugeValue, float64(server.SSLEnabled), serverHost)
			ch <- prometheus.MustNewConstMetric(statusServerPriority, prometheus.GaugeValue, float64(server.Priority), serverHost)
			splitValUnit := strings.Split(server.BPS, " ")
			valStr, unit := splitValUnit[0], splitValUnit[1]
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				log.Println("error parsing float", valStr, ":", err)
			}
			switch unit {
			case "":
				break
			case "K":
				val = val * 1024
				break
			case "M":
				val = val * 1024 * 1024
				break
			case "G":
				val = val * 1024 * 1024 * 1024
				break
			case "T":
				val = val * 1024 * 1024 * 1024 * 1024
				break
			case "P":
				val = val * 1024 * 1024 * 1024 * 1024 * 1024
				break
			}
			ch <- prometheus.MustNewConstMetric(statusServerBPS, prometheus.GaugeValue, val, serverHost)
		}
	}
}

func (s *sabnzbdStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- statusServerActive
	ch <- statusServerActiveConn
	ch <- statusServerTotalConn
	ch <- statusServerSSL
	ch <- statusServerPriority
	ch <- statusServerBPS
}
