package grpc_adm_service_impl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/adm_dto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type ServiceServiceImpl struct{}

func (ServiceServiceImpl) GetMyServiceList(ctx context.Context, empty *empty.Empty) (*adm_dto.ServiceBaseInfoDto, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		logrus.Info("token: %s", md.Get("authorization")[0])
	}
	return nil, nil
}
