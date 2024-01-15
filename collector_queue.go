package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

// Ensure interface
var _ prometheus.Collector = &sabnzbdQueueCollector{}

type sabnzbdQueueCollector struct{}

func (s *sabnzbdQueueCollector) Collect(ch chan<- prometheus.Metric) {
	out := queueOutput{}
	err := call(&out, map[string]string{"mode": "queue", "limit": "1"})
	if err != nil {
		log.Println("error collecting queue metrics: ", err)
	} else {
		ch <- prometheus.MustNewConstMetric(queueSpeedLimit, prometheus.GaugeValue, out.Queue.SpeedLimitAbs)
		ch <- prometheus.MustNewConstMetric(queueSpeed, prometheus.GaugeValue, out.Queue.KBPerSec*1024)
		ch <- prometheus.MustNewConstMetric(queueCount, prometheus.GaugeValue, float64(out.Queue.NoOfSlots))
		ch <- prometheus.MustNewConstMetric(queueMB, prometheus.GaugeValue, float64(out.Queue.MB))
		ch <- prometheus.MustNewConstMetric(queueMBRemaining, prometheus.GaugeValue, float64(out.Queue.MBLeft))
	}
}

func (s *sabnzbdQueueCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- queueSpeedLimit
	ch <- queueSpeed
	ch <- queueCount
	ch <- queueMB
	ch <- queueMBRemaining
}
