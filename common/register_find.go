package common

func init() {
	ServerIPMap = make(map[string]string)
}

var ServerIPMap map[string]string

type ServerStatue struct {
	IP   string
	Name string
	Type string
}

func Register(s ServerStatue) error {
	return nil
}
