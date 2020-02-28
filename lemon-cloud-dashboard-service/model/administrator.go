package model

type AdministratorInfo struct {
	DataKey             string `json:"data_key"`               // 管理员唯一数据Key，系统自动生成
	Number              string `json:"number"`                 // 账号，登录用，唯一登录方式，邮件等不可以作为登录方式
	Email               string `json:"email"`                  // 电子邮件地址，只供后期邮件通知时或验证码发送等使用
	Name                string `json:"name"`                   // 此管理员用户的姓名
	PasswordCipherText  string `json:"password_cipher_text"`   // 密码密文
	IsSuperAdmin        bool   `json:"is_super_admin"`         // 是否是超级管理员，超级管理员将忽略下面的AllowServiceKeyList字段，允许访问任何服务
	AllowServiceKeyList string `json:"allow_service_key_list"` // 允许使用的服务service_key列表，使用逗号分隔拼接成的字符串
	LatestLoginAt       int64  `json:"latest_login_at"`        // 最近登录时间
	CreateAt            int64  `json:"create_at"`              // 账号创建时间
}
