package handler

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_micro_service"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/define"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/manager"
	"os"
)

func SystemStart() {
	var err error

	// print system info
	define.PrintSystemInfo()

	// read config file
	lccu_log.Info("Start reading configuration files...")
	err = manager.ConfigManagerInstance().Init()
	if err != nil {
		lccu_log.Error("System start failed. Error reading configuration file: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Configuration file read completed")

	// init registry
	lccu_log.Info("Start configuring the registry...")
	err = lccc_micro_service.SystemServiceInstance().RegisterNewService(&lccc_micro_service.ServiceRegisterConfig{
		ServiceGeneralConfig: manager.ConfigManagerInstance().GeneralConfig(),
		ServiceInfo:          define.GetServiceInfo(),
	}, define.GetSystemSettings())
	if err != nil {
		lccu_log.Error("System start failed. Error configuring registry: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Registry config completed")

	// start grpc server
	StartGrpcWebServer()

	lccu_log.Infoln("Service started, waiting for access")
}
