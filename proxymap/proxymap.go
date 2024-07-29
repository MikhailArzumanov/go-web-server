package proxymap

import (
	"encoding/json"
	"go-web-server/config"
	"go-web-server/errors_handling"
	"os"
)

type ProxyMap struct {
	Self map[string]string
}

var Map ProxyMap

func InitProxymap() {
	mapPath := config.MAP_PATH
	jsonBytes, err := os.ReadFile(mapPath)
	errors_handling.Handle(err)

	err = json.Unmarshal(jsonBytes, &Map)
	errors_handling.Handle(err)
}
