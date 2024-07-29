package config

import (
	"fmt"
	"go-web-server/errors_handling"
	"path/filepath"
)

var APP_ROOT string
var MAP_PATH string
var HTDOCS_PATH string

var LOG_FILES_DIRPATH string

var PROTOCOL = "https"
var DOMAIN = "localhost"
var PORT = "443"
var REDIRECT_BASEURL string

var MODE = "HTTP"

var CRT_FILEPATH string
var KEY_FILEPATH string

func InitConfig() {
	var err error
	APP_ROOT, err = filepath.Abs("../")
	errors_handling.Handle(err)

	MAP_PATH = APP_ROOT + "/config/proxymap.json"

	LOG_FILES_DIRPATH = APP_ROOT + "/_log_files/"

	REDIRECT_BASEURL = fmt.Sprintf("%s://%s:%s", PROTOCOL, DOMAIN, PORT)
}
