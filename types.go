package main

type queueOutput struct {
	Queue struct {
		SpeedLimitAbs float64 `json:"speedlimit_abs,string"`
		KBPerSec      float64 `json:"kbpersec,string"`
		MB            float64 `json:"mb,string"`
		MBLeft        float64 `json:"mbleft,string"`
		NoOfSlots     int64   `json:"noofslots"`
	} `json:"queue"`
}

type historyOutput struct {
	History struct {
		NoOfSlots int64 `json:"noofslots"`
		PPSlots   int64 `json:"ppslots"`
	} `json:"history"`
}

type serverStatsOutput struct {
	Total   int64 `json:"total"`
	Servers map[string]struct {
		Total           int64            `json:"total"`
		ArticlesTried   map[string]int64 `json:"articles_tried"`
		ArticlesSuccess map[string]int64 `json:"articles_success"`
	} `json:"servers"`
}

type statusOutput struct {
	Status struct {
		Servers []struct {
			Active     bool   `json:"serveractive"`
			ActiveConn int64  `json:"serveractiveconn"`
			TotalConn  int64  `json:"servertotalconn"`
			Name       string `json:"servername"`
			SSLEnabled int64  `json:"serverssl"`
			Priority   int64  `json:"serverpriority"`
			BPS        string `json:"serverbps"`
		} `json:"servers"`
	} `json:"status"`
}

type configServersOutput struct {
	Config struct {
		Servers []struct {
			DisplayName string `json:"displayname"`
			Host        string `json:"host"`
		} `json:"servers"`
	} `json:"config"`
}
