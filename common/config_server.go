package common

// 服务器网络配置
type configServerNet struct {
	// 内部服务器的监听
	Rpc string
	// 命令模式的监听
	Cmd string
	// 外部监听配置
	Ext string
	// http监听配置
	Http string
	// 从一个http地址获取配置ip
	// 用于防DDOS攻击
	ExtIpUrl string
}

type ConfigMysql struct {
	Alias    string
	Url      string
	Dbs      []string
	Idleconn int
	Openconn int
}

func (cfg *ConfigMysql) Clone() *ConfigMysql {
	newCfg := &ConfigMysql{
		Alias:    cfg.Alias,
		Url:      cfg.Url,
		Idleconn: cfg.Idleconn,
		Openconn: cfg.Openconn,
		Dbs:      make([]string, len(cfg.Dbs), len(cfg.Dbs)),
	}
	copy(newCfg.Dbs, cfg.Dbs)
	return newCfg
}

// 服务器配置
type configServer struct {
	Type string
	Id   string
	Name string
	Net  *configServerNet

	Redis string
	Mysql ConfigMysql
}
