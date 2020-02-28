package core_service

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_core"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_strings"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/define"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"strconv"
	"strings"
	"sync"
	"time"
)

type JwtServiceStruct struct {
}

var jwtServiceInstance *JwtServiceStruct
var jwtServiceOnce sync.Once

// 单例函数
func JwtService() *JwtServiceStruct {
	jwtServiceOnce.Do(func() {
		jwtServiceInstance = &JwtServiceStruct{}
	})
	return jwtServiceInstance
}

func (jss *JwtServiceStruct) GetSessionDuration() int64 {
	if val, err := strconv.ParseInt(lccc_core.SystemSettingsService().KGet(define.SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_SESSION_DURATION), 10, 64); err != nil {
		return 120
	} else {
		return val
	}
}

// 生成一个Token
// administratorDataKey 管理员用户的唯一标识
// allowServiceKeyList 用户拥有的业务使用权限列表，将service_key以逗号分隔，拼接一个字符串, 超级管理员的本字段为一个*号
func (jss *JwtServiceStruct) GenerateToken(administratorDataKey, allowServiceKeyList string) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		Audience:  administratorDataKey,
		ExpiresAt: now.Unix() + jss.GetSessionDuration(),
		Id:        lccu_strings.RandomUUIDString(),
		IssuedAt:  now.Unix(),
		Issuer:    administratorDataKey,
		NotBefore: now.Unix(),
		Subject:   allowServiceKeyList,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(lccc_core.SystemSettingsService().KGet(define.SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_JWT_SIGN_KEY))
}

// 检查某一个token是否允许使用指定的业务（根据service_key）
func (jss *JwtServiceStruct) CheckToken(token, serviceKey string) bool {
	if jwtObj, err := jwt.Parse(token, func(jwtTokenObj *jwt.Token) (interface{}, error) {
		return lccc_core.SystemSettingsService().KGet(define.SYSTEM_SETTINGS_ITEM_KEY_SYSTEM_JWT_SIGN_KEY), nil
	}); err != nil {
		return false
	} else if !jwtObj.Valid {
		return false
	} else if claims, ok := jwtObj.Claims.(jwt.StandardClaims); !ok {
		return false
	} else {
		if strings.EqualFold(claims.Subject, "*") {
			// 超级管理员，可以访问任何功能
			return true
		}
		return strings.Contains(claims.Subject, serviceKey)
	}
}
