package handler

import "github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/core_service"

func StartWorkspaceProcessHandler() error {
	if err := core_service.WorkspaceService().RepairWorkspaceDirectoryStructure(); err != nil {
		return err
	}
	if _, err := core_service.WorkspaceService().ScanAndGetAllServiceManagementInfoFromWorkspace(); err != nil {
		return err
	}
	return nil
}
