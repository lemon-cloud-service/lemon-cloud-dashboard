package usr_service_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/usr_dto"
)

type AdminUsrServiceImpl struct{}

func (AdminUsrServiceImpl) Login(context.Context, *usr_dto.AdminLoginRequest) (*usr_dto.AdminLoginResponse, error) {
	return &usr_dto.AdminLoginResponse{
		Token:    "token1122334455",
		Username: "LemonIT.CN柠檬信息技术有限公司",
	}, nil
}
