package core_service

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_config"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_io"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/model"
	"io/ioutil"
	"sync"
)

type WorkspaceServiceStruct struct {
	PathWorkspace                        string `json:"path_workspace"`
	PathWorkspaceStatic                  string `json:"path_static"`
	PathWorkspaceStaticServiceManagement string `json:"path_workspace_static_service_management"`
	PathWorkspaceOverrideFile            string `json:"path_workspace_override_file"`
	// 缓存
	AllServiceManagementDefineCache map[string]*model.ServiceManagementDefine `json:"all_service_management_define_cache"`
}

var workspaceServiceInstance *WorkspaceServiceStruct
var workspaceServiceOnce sync.Once

const WORKSPACE_KEY_WORKSPACE string = "workspace"
const WORKSPACE_KEY_SERVICE_MANAGEMENT string = "service_management"
const WORKSPACE_KEY_STATIC string = "static"
const WORKSPACE_KEY_LC_OVERRIDE_FILE string = "lc_override.json"
const WORKSPACE_KEY_LC_SERVICE_FILE string = "lc_service.json"

// 单例函数
func WorkspaceService() *WorkspaceServiceStruct {
	workspaceServiceOnce.Do(func() {
		workspaceServiceInstance = &WorkspaceServiceStruct{
			PathWorkspace:                        WORKSPACE_KEY_WORKSPACE,
			PathWorkspaceStatic:                  fmt.Sprintf("%v/%v", WORKSPACE_KEY_WORKSPACE, WORKSPACE_KEY_STATIC),
			PathWorkspaceStaticServiceManagement: fmt.Sprintf("%v/%v/%v", WORKSPACE_KEY_WORKSPACE, WORKSPACE_KEY_STATIC, WORKSPACE_KEY_SERVICE_MANAGEMENT),
			PathWorkspaceOverrideFile:            fmt.Sprintf("%v/%v", WORKSPACE_KEY_WORKSPACE, WORKSPACE_KEY_LC_OVERRIDE_FILE),
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
func (ws *WorkspaceServiceStruct) ScanAllOverrideServiceManagementInfoFromWorkspace() (map[string]*model.ServiceManagementDefine, error) {
	result := make(map[string]*model.ServiceManagementDefine)
	overrideServices := &struct {
		Services []*model.ServiceManagementDefine `json:"services"`
	}{}
	if exists, err := lccu_io.PathIsExists(ws.PathWorkspaceOverrideFile); err != nil {
		return result, err
	} else if !exists {
		// 没有提供override文件，认为不需要任何覆盖，返回空map
		return result, nil
	}
	// 提供了override文件，开始读取
	if err := lccu_config.LoadJsonConfigFile(ws.PathWorkspaceOverrideFile, overrideServices); err != nil {
		return result, err
	}
	for _, serviceManagementDefine := range overrideServices.Services {
		result[serviceManagementDefine.ServiceKey] = serviceManagementDefine
	}
	return result, nil
}

// 先从工作区中按照约定的格式扫描所有服务信息，在与重写信息进行拼装得到最终的服务管理信息定义
func (ws *WorkspaceServiceStruct) ScanAndGetAllServiceManagementInfoFromWorkspace() (map[string]*model.ServiceManagementDefine, error) {
	lccu_log.Info("Start reading service management definition data from the workspace...")
	result := make(map[string]*model.ServiceManagementDefine)
	if fileInfoList, err := ioutil.ReadDir(ws.PathWorkspaceStaticServiceManagement); err != nil {
		return result, err
	} else {
		for _, fileInfo := range fileInfoList {
			serviceConfigFilePath := fmt.Sprintf("%v/%v/%v", ws.PathWorkspaceStaticServiceManagement, fileInfo.Name(), WORKSPACE_KEY_LC_SERVICE_FILE)
			if fileInfo.IsDir() {
				// 是目录，开始判断里面是否有服务信息配置文件
				if exists, err := lccu_io.PathIsExists(serviceConfigFilePath); err != nil {
					continue
				} else if exists {
					if lccu_io.PathIsDir(serviceConfigFilePath) {
						continue
					}
				}
			} else {
				// 不是目录，跳过
				continue
			}
			// 能到这里说明这个路径确实是一个配置文件，尝试读取
			serviceManagementDefine := &struct {
				service *model.ServiceManagementDefine
			}{}
			if err := lccu_config.LoadJsonConfigFile(serviceConfigFilePath, serviceManagementDefine); err != nil {
				return result, err
			}
			result[serviceManagementDefine.service.ServiceKey] = serviceManagementDefine.service
		}
	}
	// 从本地文件夹中都已经扫描完成，开始于override中的数据覆盖合并
	if overrideResult, err := ws.ScanAllOverrideServiceManagementInfoFromWorkspace(); err != nil {
		return result, nil
	} else {
		// 成功拿到了override结果，开始合并
		for serviceKey, overrideServiceManagementDefine := range overrideResult {
			if oldServiceManageDefine, ok := result[serviceKey]; !ok {
				// 直接扫描的结果中不包含这个override中定义的服务，那么全部添加进去
				result[serviceKey] = overrideServiceManagementDefine
			} else {
				// 直接扫描的结果中也有这个服务定义，那么遍历具体的每一个override管理模块，逐个判断覆盖，不存在的则追加
				oldServiceManageDefine.ServiceIconUrl = overrideServiceManagementDefine.ServiceIconUrl
				tempModuleMapping := make(map[string]*model.ServiceManagementModuleDefine)
				for _, moduleDefine := range oldServiceManageDefine.ManagementModuleList {
					// 遍历所有的约定扫描出来的管理模块，并存到临时的map中
					tempModuleMapping[moduleDefine.ModuleKey] = moduleDefine
				}
				for _, overrideModuleDefine := range overrideServiceManagementDefine.ManagementModuleList {
					tempModuleMapping[overrideModuleDefine.ModuleKey] = overrideModuleDefine
				}
				// 此时tempModuleMapping中的所有module为合并后的模块，将其转换为数组并重新放置回去
				var newModuleList []*model.ServiceManagementModuleDefine
				for _, moduleDefine := range tempModuleMapping {
					newModuleList = append(newModuleList, moduleDefine)
				}
				oldServiceManageDefine.ManagementModuleList = newModuleList
			}
		}
	}
	// 存储到缓存
	ws.AllServiceManagementDefineCache = result
	lccu_log.Infof("Successfully read %d services management definition data from the workspace", len(result))
	return result, nil
}

// 快速直接从缓存中读取服务管理数据定义数据
func (ws *WorkspaceServiceStruct) FastGetAllServiceManagementInfoFromWorkspace() map[string]*model.ServiceManagementDefine {
	return ws.AllServiceManagementDefineCache
}
