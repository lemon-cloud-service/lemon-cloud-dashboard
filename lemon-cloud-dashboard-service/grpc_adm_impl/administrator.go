package grpc_adm_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/grpc_adm"
)

type GrpcAdmAdministratorImpl struct{}

func (GrpcAdmAdministratorImpl) Login(context.Context, *grpc_adm.AdministratorLoginRequestDto, *grpc_adm.AdministratorLoginResponseDto) error {
	panic("implement me")
}
