package handler

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_core"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/define"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/manager"
	"os"
)

func SystemStart() {
	var err error
	// 打印系统信息
	define.PrintSystemInfo()

	// 从磁盘中读取配置文件
	lccu_log.Info("Start reading configuration files...")
	err = manager.ConfigManagerInstance().Init()
	if err != nil {
		lccu_log.Error("System start failed. Error reading configuration file: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Configuration file read completed")

	// 工作区相关初始化
	if err = StartWorkspaceProcessHandler(); err != nil {
		lccu_log.Error("An error occurred during the initialization of the workspace，reason: ", err.Error())
		os.Exit(1)
	}

	// 调用核心服务启动函数
	lccu_log.Info("Start configuring the registry...")
	err = lccc_core.CoreService().Start(&lccc_core.CoreStartParams{
		RunGrpcService:                 true,
		ServiceGeneralConfig:           manager.ConfigManagerInstance().GeneralConfig(),
		ServiceBaseInfo:                define.GetServiceBaseInfo(),
		ServiceApplicationInfo:         define.GetServiceApplicationInfo(),
		GrpcServiceImplRegisterHandler: registerAllGrpcImpl,
		SystemSettingsDefine:           define.GetSystemSettings(),
	})
	if err != nil {
		lccu_log.Error("System start failed. Reason: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Registry config completed")

	// 启动完毕
	lccu_log.Infoln("Service started, waiting for access")
}
