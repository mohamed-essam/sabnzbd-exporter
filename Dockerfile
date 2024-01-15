FROM golang:1.21.6 AS builder
WORKDIR /go/src/github.com/mohamed-essam/sabnzbd-exporter
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o exporter .

FROM alpine:latest
COPY --from=builder /go/src/github.com/mohamed-essam/sabnzbd-exporter/exporter .
CMD ["./exporter"]