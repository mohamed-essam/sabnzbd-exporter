package main

import "github.com/prometheus/client_golang/prometheus"

var (
	queueSpeedLimit  = prometheus.NewDesc("queue_speed_limit", "Queue speed limit in B/s", nil, nil)
	queueSpeed       = prometheus.NewDesc("queue_speed", "Queue current speed in B/s", nil, nil)
	queueMB          = prometheus.NewDesc("queue_mb", "Queue total MBs", nil, nil)
	queueMBRemaining = prometheus.NewDesc("queue_mb_remaining", "Queue total remaining MBs", nil, nil)
	queueCount       = prometheus.NewDesc("queue_count", "Queue number of items", nil, nil)
)

var (
	historyCount   = prometheus.NewDesc("history_count", "Number of items in history", nil, nil)
	historyPPCount = prometheus.NewDesc("history_pp_count", "Number of items in post-processing", nil, nil)
)

var (
	serverStatsTotalDownload   = prometheus.NewDesc("total_download", "Total amount downloaded in bytes", nil, nil)
	serverStatsServerTotal     = prometheus.NewDesc("server_total_download", "Total amount downloaded from server in bytes", []string{"server_host"}, nil)
	serverStatsArticlesTried   = prometheus.NewDesc("server_articles_tried", "Number of articles tried on server, day=today returns current day counter", []string{"server_host", "day"}, nil)
	serverStatsArticlesSuccess = prometheus.NewDesc("server_articles_success", "Number of articles found on server, day=today returns current day counter", []string{"server_host", "day"}, nil)
)

var (
	statusServerActive     = prometheus.NewDesc("server_active", "Server enabled", []string{"server_host"}, nil)
	statusServerActiveConn = prometheus.NewDesc("server_active_conn", "Server active connection count", []string{"server_host"}, nil)
	statusServerTotalConn  = prometheus.NewDesc("server_max_conn", "Server maximum connection count, 0 is unset", []string{"server_host"}, nil)
	statusServerSSL        = prometheus.NewDesc("server_ssl_enabled", "Server SSL enabled", []string{"server_host"}, nil)
	statusServerPriority   = prometheus.NewDesc("server_priority", "Server priority", []string{"server_host"}, nil)
	statusServerBPS        = prometheus.NewDesc("server_speed", "Server current speed in b/s", []string{"server_host"}, nil)
)
