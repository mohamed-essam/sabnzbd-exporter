package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Ensure interface
var _ prometheus.Collector = &sabnzbdServerStatsCollector{}

type sabnzbdServerStatsCollector struct{}

func (s *sabnzbdServerStatsCollector) Collect(ch chan<- prometheus.Metric) {
	out := serverStatsOutput{}
	err := call(&out, map[string]string{"mode": "server_stats"})
	loc, err := time.LoadLocation(*sabnzbdTimezone)
	if err != nil {
		log.Println("error loading timezone", *sabnzbdTimezone, ":", err)
	}
	today := time.Now().In(loc)
	if err != nil {
		log.Println("error collecting history metrics: ", err)
	} else {
		ch <- prometheus.MustNewConstMetric(serverStatsTotalDownload, prometheus.GaugeValue, float64(out.Total))
		for server, data := range out.Servers {
			ch <- prometheus.MustNewConstMetric(serverStatsServerTotal, prometheus.CounterValue, float64(data.Total), server)
			for day, articleCount := range data.ArticlesTried {
				ch <- prometheus.MustNewConstMetric(serverStatsArticlesTried, prometheus.CounterValue, float64(articleCount), server, day)
			}
			ch <- prometheus.MustNewConstMetric(serverStatsArticlesTried, prometheus.CounterValue, float64(data.ArticlesTried[today.Format(time.DateOnly)]), server, "today")
			for day, articleCount := range data.ArticlesSuccess {
				ch <- prometheus.MustNewConstMetric(serverStatsArticlesSuccess, prometheus.CounterValue, float64(articleCount), server, day)
			}
			ch <- prometheus.MustNewConstMetric(serverStatsArticlesSuccess, prometheus.CounterValue, float64(data.ArticlesSuccess[today.Format(time.DateOnly)]), server, "today")
		}
	}
}

func (s *sabnzbdServerStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- serverStatsTotalDownload
	ch <- serverStatsServerTotal
	ch <- serverStatsArticlesTried
	ch <- serverStatsArticlesSuccess
}
