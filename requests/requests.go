package requests

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/logs"
	"go-web-server/proxymap"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getType(path string) string {
	parts := strings.Split(path, ".")
	if len(parts) <= 1 {
		return "text/html"
	}

	extension := parts[len(parts)-1]

	if extension == "js" {
		return "application/javascript"
	} else if extension == "css" {
		return "text/css"
	} else if extension == "html" {
		return "text/html"
	}

	return "text/plain"
}

var basePath = config.HTDOCS_PATH

func setHeader(w http.ResponseWriter, relpath string) {
	var typeStr = getType(relpath)
	w.Header().Set("Content-Type", typeStr+"; charset=utf-8")
}

func openFile(w http.ResponseWriter, r *http.Request,
	staticPath string) *os.File {

	var file, err = os.Open(staticPath)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		logs.LogError(err)
		http.NotFound(w, r)
		file.Close()
		return nil
	}
	return file
}

func sendFile(w http.ResponseWriter, file *os.File, staticPath string) {
	var data = make([]byte, 1024)
	for {
		var l, err = file.Read(data)
		fmt.Fprint(w, string(data[:l]))
		if err != nil {
			fmt.Println(err)
			logs.LogFileRead(err, staticPath)
			break
		}
	}
	file.Close()
}

func processFile(w http.ResponseWriter, r *http.Request, staticPath string) {
	file := openFile(w, r, staticPath)
	if file == nil {
		w.Header().Del("Content-Type")
		return
	}
	sendFile(w, file, staticPath)
}

func handleFileRequest(w http.ResponseWriter, r *http.Request, relpath string) {
	var path = basePath + relpath
	var staticPath, _ = filepath.Abs(path)
	logs.LogFileRequest(r, staticPath)
	setHeader(w, relpath)
	processFile(w, r, staticPath)
}

func reduceRelpath(r *http.Request) string {
	var relpath = proxymap.Map.Self[r.URL.String()]
	if relpath == "" {
		relpath = r.URL.String()
	}
	return relpath
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	relpath := reduceRelpath(r)
	handleFileRequest(w, r, relpath)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	redirectBaseurl := config.REDIRECT_BASEURL
	var redirectPath = redirectBaseurl + r.RequestURI
	http.Redirect(w, r, redirectPath, http.StatusMovedPermanently)
}
