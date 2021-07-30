# grafana-poc
Golang examples with Grafana integration
- Go Gin Web Framework link [Go-Gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- Documentation link [Grafana](https://grafana.com/docs/grafana/latest/getting-started)
- Documentation link [Prometheus](https://prometheus.io/docs/introduction/overview)

## Used Packages
- Prometheus gin-metrics [genglongi/gin-metrics](https://github.com/penglongli/gin-metrics)
- Prometheus Client [prometheus/client_golang](https://github.com/prometheus/client_golang)

## Makefile
- `make build` to build the golang binary of application
- `make run` to run application on local
- `make clean` to clean generated binary of application

## Run application with Grafana on local
- `make docker-up`

## Stop application with Grafana on local
- `make docker-down`

## Health Check API
- `http://localhost:8080/health`

## Get Feeds Data API
- `http://localhost:8080/feds`

## Check Grafana Metrics API 
- `http://localhost:8080/metrics`

## Grafana Dashboard view
- `http://localhost:3000`