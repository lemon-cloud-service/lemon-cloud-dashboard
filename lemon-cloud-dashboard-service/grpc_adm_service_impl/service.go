package grpc_adm_service_impl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/adm_dto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type ServiceServiceImpl struct{}

func (ServiceServiceImpl) GetMyServiceList(ctx context.Context, empty *empty.Empty) (*adm_dto.AllServiceStatusDto, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		logrus.Info("token: %s", md.Get("authorization")[0])
	}
	return &adm_dto.AllServiceStatusDto{
		AllServiceStatusList: []*adm_dto.ServiceStatusDto{
			&adm_dto.ServiceStatusDto{
				Info: &adm_dto.ServiceInfoDto{
					Name:        "Lemon Cloud User Service",
					Tag:         "lemon_cloud_user",
					Introduce:   "这是一个用户服务",
					ServiceIcon: "hello",
				},
				OnlineInfoList: []*adm_dto.OnlineServiceInfoDto{
					&adm_dto.OnlineServiceInfoDto{
						EndpointHost:          "www.lemonit.cn",
						EndpointPort:          33385,
						ApplicationVersion:    "1.0.0",
						ApplicationVersionNum: 1,
					},
				},
			},
		},
	}, nil
}
