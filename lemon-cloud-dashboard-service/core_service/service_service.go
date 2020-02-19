package core_service

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_micro_service"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
	"sync"
)

type ServiceServiceStruct struct {
}

var serviceServiceInstance *ServiceServiceStruct
var serviceServiceOnce sync.Once

// 单例函数
func ServiceService() *ServiceServiceStruct {
	serviceServiceOnce.Do(func() {
		serviceServiceInstance = &ServiceServiceStruct{}
	})
	return serviceServiceInstance
}

// 获取当前已登录管理员所有有权限的服务基础信息
func (ss *ServiceServiceStruct) GetMyAllServiceBaseInfo() (map[string]*lccc_model.ServiceBaseInfo, error) {
	return lccc_micro_service.CoreServiceSingletonInstance().FastGetAllServiceBaseInfo(), nil
}
