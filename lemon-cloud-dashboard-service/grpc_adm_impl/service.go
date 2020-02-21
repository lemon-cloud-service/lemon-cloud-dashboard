package grpc_adm_impl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/grpc_adm"
)

type GrpcAdmServiceImpl struct{}

func (GrpcAdmServiceImpl) GetMyServiceBaseInfoList(context.Context, *empty.Empty, *grpc_adm.ServiceBaseInfoListDto) error {
	panic("implement me")
}

func (GrpcAdmServiceImpl) GetMyServiceInstanceInfoList(context.Context, *empty.Empty, *grpc_adm.ServiceInstanceInfoListDto) error {
	panic("implement me")
}
