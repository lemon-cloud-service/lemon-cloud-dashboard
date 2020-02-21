package handler

import (
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/grpc_adm"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/grpc_adm_impl"
	"github.com/micro/go-micro/v2/server"
)

func registerAllGrpcImpl(server server.Server) {
	grpc_adm.RegisterAdministratorServiceHandler(server, new(grpc_adm_impl.GrpcAdmAdministratorImpl))
}
