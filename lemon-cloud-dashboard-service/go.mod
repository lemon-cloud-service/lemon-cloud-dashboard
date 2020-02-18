module github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service

go 1.13

require (
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/protobuf v1.3.3
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.7.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0
	google.golang.org/grpc v1.27.1
)

replace github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common => ../lemon-cloud-dashboard-common

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils => ../../lemon-cloud-common/lemon-cloud-common-utils

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components => ../../lemon-cloud-common/lemon-cloud-common-components

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
