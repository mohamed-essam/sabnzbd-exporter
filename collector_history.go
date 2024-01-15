package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

// Ensure interface
var _ prometheus.Collector = &sabnzbdHistoryCollector{}

type sabnzbdHistoryCollector struct{}

func (s *sabnzbdHistoryCollector) Collect(ch chan<- prometheus.Metric) {
	out := historyOutput{}
	err := call(&out, map[string]string{"mode": "history", "limit": "1"})
	if err != nil {
		log.Println("error collecting history metrics: ", err)
	} else {
		ch <- prometheus.MustNewConstMetric(historyCount, prometheus.GaugeValue, float64(out.History.NoOfSlots))
		ch <- prometheus.MustNewConstMetric(historyPPCount, prometheus.GaugeValue, float64(out.History.PPSlots))
	}
}

func (s *sabnzbdHistoryCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- historyCount
	ch <- historyPPCount
}
