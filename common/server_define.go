package common

// 服务器类型名
const (
	// 账号登录服
	ServerType_auth = "auth"
	// api服务
	ServerType_api = "api"
	// draw服务
	ServerType_draw = "draw"
)

// redis key name

const (
	// 服务注册与发现
	Server_Type   = "Server_%s"
	Server_IP     = "Server_%s_IP"
	Server_status = "Server_%s_status"
)
