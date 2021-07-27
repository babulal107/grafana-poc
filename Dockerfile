FROM golang:alpine
COPY . /go/src/grafana-poc
WORKDIR /go/src/grafana-poc
RUN go mod vendor
RUN cd cmd/service && go build -o grafana-app .
CMD ["./cmd/service/grafana-app"]