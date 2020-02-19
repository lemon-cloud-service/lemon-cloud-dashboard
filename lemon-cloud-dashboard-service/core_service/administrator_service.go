package core_service

import "sync"

type AdministratorServiceStruct struct {
}

var administratorServiceInstance *AdministratorServiceStruct
var administratorServiceOnce sync.Once

// 单例函数
func AdministratorService() *AdministratorServiceStruct {
	administratorServiceOnce.Do(func() {
		administratorServiceInstance = &AdministratorServiceStruct{}
	})
	return administratorServiceInstance
}
