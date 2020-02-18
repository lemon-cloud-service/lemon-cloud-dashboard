package define

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_define"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
)

const SYSTEM_INFO_NAME string = "Lemon Cloud Dashboard"
const SYSTEM_INFO_SERVICE_TAG string = "lemon_cloud_dashboard"
const SYSTEM_INFO_SERVICE_INTRODUCE string = "Lemon Cloud Dashboard Service"
const SYSTEM_INFO_VERSION string = "1.0.0"
const SYSTEM_INFO_VERSION_NUM uint16 = 1
const SYSTEM_INFO_SPLIT_LINE string = "====================================================================="

func GetServiceInfo() *lccc_model.ServiceInfo {
	return &lccc_model.ServiceInfo{
		ServiceTag:            SYSTEM_INFO_SERVICE_TAG,
		ServiceName:           SYSTEM_INFO_NAME,
		ServiceIntroduce:      SYSTEM_INFO_SERVICE_INTRODUCE,
		ApplicationVersion:    SYSTEM_INFO_VERSION,
		ApplicationVersionNum: SYSTEM_INFO_VERSION_NUM,
	}
}

// 打印系统的基础信息，包含LemonCloud字符画和系统名称及版本
func PrintSystemInfo() {
	fmt.Print(lccc_define.LEMON_CLOUD_ASCII_IMAGE)
	fmt.Print("\n")
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
	fmt.Printf("Welcome to %v [ver: %v(%d)]\n", SYSTEM_INFO_NAME, SYSTEM_INFO_VERSION, SYSTEM_INFO_VERSION_NUM)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
}
