module github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service

go 1.13

require (
	github.com/golang/protobuf v1.3.3
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.1.0
)

replace github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common => ../lemon-cloud-dashboard-common

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils => ../../lemon-cloud-common/lemon-cloud-common-utils

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components => ../../lemon-cloud-common/lemon-cloud-common-components

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
