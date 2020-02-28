package define

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
)

const SYSTEM_SETTINGS_GROUP_KEY_SYSTEM = "system"

var SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_SESSION_DURATION = lccc_model.SystemSettingItemKeyDefine{
	GroupKey: SYSTEM_SETTINGS_GROUP_KEY_SYSTEM,
	ItemKey:  "session_duration",
}
var SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_JWT_SIGN_KEY = lccc_model.SystemSettingItemKeyDefine{
	GroupKey: SYSTEM_SETTINGS_GROUP_KEY_SYSTEM,
	ItemKey:  "jwt_sign_key",
}

func GetSystemSettings() *lccc_model.SystemSettingsDefine {
	return &lccc_model.SystemSettingsDefine{
		Introduce: "对柠檬云Dashboard服务的相关设置",
		SettingGroupList: []*lccc_model.SystemSettingGroupDefine{
			{
				Key:       SYSTEM_SETTINGS_GROUP_KEY_SYSTEM,
				Name:      "系统设置",
				Introduce: "这里包含系统运行所需要的重要相关配置",
				SettingList: []*lccc_model.SystemSettingItemDefine{
					{
						Key:          SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_SESSION_DURATION.ItemKey,
						Name:         "登录会话时长",
						UiType:       "input",
						UiParams:     `^[0-9]{3,6}$`,
						DefaultValue: "86400",
						Introduce:    "签发的JWT Token过期时长，单位秒(s)，请根据需要设置，建议不要设置过小的时间。3-6位纯数字",
						NeedRestart:  false,
					},
					{
						Key:          SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_SESSION_DURATION.ItemKey,
						Name:         "JWT签名密钥",
						UiType:       "input",
						UiParams:     `^[A-Za-z0-9]{8,128}$`,
						DefaultValue: "86400",
						Introduce:    "在签发JWT Token的时候所使用的的签名密钥。8-128位字符串，支持大小写字母、数字",
						NeedRestart:  false,
					},
				},
			},
		},
	}
}
