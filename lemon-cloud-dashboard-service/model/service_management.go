package model

// workspace中，各个业务系统对自己功能和模块定义时所需要的结构体

// 业务管理模块定义
type ServiceManagementModuleDefine struct {
	ModuleKey       string `json:"module_key" yaml:"module_key"`             // 管理模块唯一标识，如果出现重复后定义的会将先定义的模块覆盖
	ModuleName      string `json:"module_name" yaml:"module_name"`           // 管理模块名称
	ModuleIntroduce string `json:"module_introduce" yaml:"module_introduce"` // 管理模块介绍
	ModuleIconUrl   string `json:"module_icon_url" yaml:"module_icon_url"`   // 管理模块的图标URL，指定文件夹内的文件时需以/开头
	IndexUrl        string `json:"index_url" yaml:"index_url"`               // 管理模块的入口URL，指定文件夹内的文件时需以/开头，支持URL中包含#锚点以及?参数
}

// 业务系统管理定义
type ServiceManagementDefine struct {
	ServiceKey           string                           `json:"service_key"  yaml:"service_icon_url"`                 // 服务的唯一标识
	ServiceIconUrl       string                           `json:"service_icon_url" yaml:"service_icon_url"`             // 业务管理的图标URL，指定文件夹内的文件时需以/开头
	ManagementModuleList []*ServiceManagementModuleDefine `json:"management_module_list" yaml:"management_module_list"` // 管理模块列表
}
