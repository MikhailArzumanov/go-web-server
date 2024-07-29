package main

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/logs"
	"go-web-server/proxymap"
	"go-web-server/requests"
	"net/http"
)

const (
	HTTP_PORT  = ":80"
	HTTPS_PORT = ":443"
)

func initialize() {
	config.InitConfig()
	logs.InitLogging()
	proxymap.InitProxymap()
}

func runRedirectionGoroutine() {
	go func() {
		handler := http.HandlerFunc(requests.HandleRedirect)
		http.ListenAndServe(HTTP_PORT, handler)
	}()
}

func startHttpsMode() {
	runRedirectionGoroutine()

	http.HandleFunc("/", requests.Handle)
	http.ListenAndServeTLS(
		HTTPS_PORT,
		config.CRT_FILEPATH,
		config.KEY_FILEPATH,
		nil,
	)
}

func startHttpMode() {
	http.HandleFunc("/", requests.Handle)
	http.ListenAndServe(HTTP_PORT, nil)
}

func main() {
	initialize()

	if config.MODE == "HTTP" {
		startHttpMode()
	} else if config.MODE == "HTTPS" {
		startHttpsMode()
	} else {
		fmt.Println("Error: check config files.")
	}
}
