package core_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_core"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_strings"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/define"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/model"
	"sync"
)

const SYSTEM_CREATE_ADMINISTRATOR = "administrator"

type AdministratorServiceStruct struct {
	AdministratorPool map[string]*model.AdministratorInfo `json:"administrator_pool"`   // 管理员DataKey和管理员对象的映射关系，map[dataKey]administratorInfo
	NumberDataKeyPool map[string]string                   `json:"number_data_key_pool"` // 账号和管理员DataKey的映射关系map[number]dataKey
}

var administratorServiceInstance *AdministratorServiceStruct
var administratorServiceOnce sync.Once

// 单例函数
func AdministratorService() *AdministratorServiceStruct {
	administratorServiceOnce.Do(func() {
		administratorServiceInstance = &AdministratorServiceStruct{}
		administratorServiceInstance.SyncAdministratorPool()
	})
	return administratorServiceInstance
}

func (ass *AdministratorServiceStruct) Login(number, password string) {

}

func (ass *AdministratorServiceStruct) encryptPassword(password string) string {

}

// 添加管理员账号
func (ass *AdministratorServiceStruct) Add(administratorInfo *model.AdministratorInfo) error {
	ass.SyncAdministratorPool()
	if _, ok := ass.NumberDataKeyPool[administratorInfo.Number]; ok {
		return errors.New("the newly created administrator number is already occupied")
	}
	newDataKey := lccu_strings.RandomUUIDStringNoLine()
	administratorInfo.DataKey = newDataKey
	ass.AdministratorPool[newDataKey] = administratorInfo
	ass.NumberDataKeyPool[administratorInfo.Number] = newDataKey
	ass.SaveAdministratorPool()
	return nil
}

// 从数据沙箱中同步管理员池数据
func (ass *AdministratorServiceStruct) SyncAdministratorPool() {
	data := lccc_core.DataSandboxService().Get(define.DATA_SANDBOX_KEY_ADMINISTRATOR_POOL)
	err := json.Unmarshal([]byte(data), ass.AdministratorPool)
	if err != nil {
		lccu_log.Errorf("Failed to sync admin user data pool from data sandbox，reason: %v", err.Error())
	}
	ass.NumberDataKeyPool = make(map[string]string)
	for _, administratorInfo := range ass.AdministratorPool {
		ass.NumberDataKeyPool[administratorInfo.Number] = administratorInfo.DataKey
	}
}

// 保存管理员池数据到数据沙箱中
func (ass *AdministratorServiceStruct) SaveAdministratorPool() {
	jsonData, err := json.Marshal(ass.AdministratorPool)
	if err != nil {
		lccu_log.Errorf("Failed to save admin user data pool to data sandbox，reason: %v", err.Error())
		return
	}
	lccc_core.DataSandboxService().Set(define.DATA_SANDBOX_KEY_ADMINISTRATOR_POOL, fmt.Sprintf("%s", jsonData))
}
