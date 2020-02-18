package grpc_adm_service_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/adm_dto"
)

type AdminUsrServiceImpl struct{}

func (AdminUsrServiceImpl) Login(context context.Context, loginRequest *adm_dto.AdministratorLoginRequestDto) (*adm_dto.AdministratorLoginResponseDto, error) {
	return &adm_dto.AdministratorLoginResponseDto{
		Token: "token1122334455",
	}, nil
}
