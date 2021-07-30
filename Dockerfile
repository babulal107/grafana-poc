FROM golang:alpine as build
COPY . /go/src/grafana-poc
WORKDIR /go/src/grafana-poc
RUN go mod vendor
RUN cd cmd/service && go build -o grafana-app .

FROM alpine:3.12
COPY --from=build /go/src/grafana-poc/cmd/service/grafana-app .
EXPOSE 8080
CMD ["./grafana-app"]