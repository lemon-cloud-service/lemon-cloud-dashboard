package grpc_adm_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/grpc_adm"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type GrpcAdmAdministratorImpl struct{}

func (GrpcAdmAdministratorImpl) Login(ctx context.Context, req *grpc_adm.AdministratorLoginRequestDto, rsp *grpc_adm.AdministratorLoginResponseDto) error {
	claims := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: 0,
		Id:        "",
		IssuedAt:  0,
		Issuer:    "",
		NotBefore: 0,
		Subject:   "",
	}
	if token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(""); err != nil {
		return err
	} else {
		rsp.Token = token
	}
	return nil
}
