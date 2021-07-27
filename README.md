# grafana-poc

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

## Check Grafana Metrics API 
- `http://localhost:8080/metrics`

## Grafana Dashboard view
- `http://localhost:3000`