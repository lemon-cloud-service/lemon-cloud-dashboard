package core_service

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_io"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/model"
	"sync"
)

type WorkspaceServiceStruct struct {
	PathWorkspace                        string `json:"path_workspace"`
	PathWorkspaceStatic                  string `json:"path_static"`
	PathWorkspaceStaticServiceManagement string `json:"path_workspace_static_service_management"`
}

var workspaceServiceInstance *WorkspaceServiceStruct
var workspaceServiceOnce sync.Once

const WORKSPACE_KEY_WORKSPACE string = "workspace"
const WORKSPACE_KEY_SERVICE_MANAGEMENT string = "service_management"
const WORKSPACE_KEY_STATIC string = "static"

// 单例函数
func WorkspaceService() *WorkspaceServiceStruct {
	workspaceServiceOnce.Do(func() {
		workspaceServiceInstance = &WorkspaceServiceStruct{
			PathWorkspace:                        WORKSPACE_KEY_WORKSPACE,
			PathWorkspaceStatic:                  fmt.Sprintf("%v/%v", WORKSPACE_KEY_WORKSPACE, WORKSPACE_KEY_STATIC),
			PathWorkspaceStaticServiceManagement: fmt.Sprintf("%v/%v/%v", WORKSPACE_KEY_WORKSPACE, WORKSPACE_KEY_STATIC, WORKSPACE_KEY_SERVICE_MANAGEMENT),
		}
	})
	return workspaceServiceInstance
}

func (ws *WorkspaceServiceStruct) RepairWorkspaceDirectoryStructure() error {
	lccu_log.Info("Start self-test repair workspace directory structure...")
	if err := lccu_io.PathDirRepair(ws.PathWorkspace); err != nil {
		return err
	} else if err = lccu_io.PathDirRepair(ws.PathWorkspaceStatic); err != nil {
		return err
	} else if err = lccu_io.PathDirRepair(ws.PathWorkspaceStaticServiceManagement); err != nil {
		return err
	}
	lccu_log.Info("Workspace self-test repair completed")
	return nil
}

// 从工作区扫描在重写配置文件中的服务信息
func (ws *WorkspaceServiceStruct) ScanAllOverrideServiceManagementInfoFromWorkspace() (map[string]model.ServiceManagementDefine, error) {
	return nil, nil
}

// 先从工作区中按照约定的格式扫描所有服务信息，在与重写信息进行拼装得到最终的服务管理信息定义
func (ws *WorkspaceServiceStruct) ScanAllServiceManagementInfoFromWorkspace() (map[string]model.ServiceManagementDefine, error) {
	return nil, nil
}
