module github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service

go 1.13

require (
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.7.0 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.27.1
)

replace github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common => ../lemon-cloud-dashboard-common
