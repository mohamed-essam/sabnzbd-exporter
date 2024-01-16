# SABnzbd exporter

A simple [SABnzbd](https://sabnzbd.org/) prometheus exporter written in Go.

## Usage

Metrics are served by default on `http://127.0.0.1:3008/metrics`

### Local Running 

```shell
go run . --sabzbd-address 'http://sabnzbd:8080' --sabnzbd-api-key ...
```

### Docker

```shell
docker run -d --restart=always -p 3008:3008 -e 'SABNZBD_EXPORTER_INSTANCE_ADDR=http://sabnzbd:8080' -e 'SABNZBD_EXPORTER_API_KEY=....' -e 'SABNZBD_EXPORTER_TZ=UTC' ghcr.io/mohamed-essam/sabnzbd-exporter:v0.0.1
```

### Docker compose

```yaml
version: '3.8'

services:
  sabnzbd-exporter:
    image: ghcr.io/mohamed-essam/sabnzbd-exporter:v0.0.1
    environment:
      SABNZBD_EXPORTER_INSTANCE_ADDR: http://sabnzbd:8080
      SABNZBD_EXPORTER_API_KEY: ....
      SABNZBD_EXPORTER_TZ: UTC
    ports:
      - 3008:3008
    restart: always
```

## Metrics

name|description|labels
---|---|---
queue_speed_limit|Queue speed limit in b/s|N/A|
queue_speed|Queue current speed in b/s|N/A|
queue_mb|Queue total MBs|N/A|
queue_mb_remaining|Queue total remaining MBs|N/A|
queue_count|Queue number of items|N/A|
history_count|Number of items in history|N/A|
history_pp_count|Number of items in post-processing|N/A|
total_download|Total amount downloaded in bytes|N/A|
server_total_download|Total amount downloaded from server in bytes| server_host
server_articles_tried|Number of articles tried on server, day=today returns current day counter| server_host, day
server_articles_success|Number of articles found on server, day=today returns current day counter| server_host, day
server_active|Server enabled| server_host
server_active_conn|Server active connection count| server_host
server_max_conn|Server maximum connection count, 0 is unset| server_host
server_ssl_enabled|Server SSL enabled| server_host
server_priority|Server priority| server_host
server_speed|Server current speed in b/s| server_host
