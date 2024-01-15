package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func envDefault(envVar, defaultVal string) string {
	if v, ok := os.LookupEnv(envVar); ok {
		return v
	}
	return defaultVal
}

var (
	addr            = flag.String("listen-address", envDefault("SABNZBD_EXPORTER_ADDR", ":3008"), "The address to listen on for HTTP requests.")
	sabnzbdAddress  = flag.String("sabnzbd-address", envDefault("SABNZBD_EXPORTER_INSTANCE_ADDR", "http://localhost:8080"), "SABnzbd instance address")
	sabnzbdAPIKey   = flag.String("sabnzbd-api-key", envDefault("SABNZBD_EXPORTER_API_KEY", ""), "SABnzbd API Key")
	sabnzbdTimezone = flag.String("sabnzbd-timezone", envDefault("SABNZBD_EXPORTER_TZ", "Local"), "Sabnzbd server timezone (needed for today's per-server article availability counter)")
)

func main() {
	flag.Parse()

	reg := prometheus.NewRegistry()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		&sabnzbdHistoryCollector{},
		&sabnzbdQueueCollector{},
		&sabnzbdServerStatsCollector{},
		&sabnzbdStatusCollector{},
	)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
