package logs

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/errors_handling"
	"net/http"
	"os"
	"time"
)

var loggingFile *os.File

const loggingDateFormat = "2006.01.2 Monday 15:04:05.0000 MST"

func LogRequest(r *http.Request) {
	fmt.Fprintf(loggingFile, "================REQUEST================\n")
	fmt.Fprintf(loggingFile, "Host: %s\n", r.Host)
	fmt.Fprintf(loggingFile, "Method: %s\n", r.Method)
	fmt.Fprintf(loggingFile, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(loggingFile, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(loggingFile, "URL: %s\n", r.URL)
	fmt.Fprintf(loggingFile, "Sender address: %s\n", r.RemoteAddr)
	fmt.Fprintf(loggingFile, "Time: %s\n", time.Now().Format(loggingDateFormat))
	fmt.Fprintf(loggingFile, "=======================================\n\n")
}

func LogFileRequest(r *http.Request, filepath string) {
	fmt.Fprintf(loggingFile, "==============FILE REQUEST==============\n")
	fmt.Fprintf(loggingFile, "Host: %s\n", r.Host)
	fmt.Fprintf(loggingFile, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(loggingFile, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(loggingFile, "URL: %s\n", r.URL)
	fmt.Fprintf(loggingFile, "Filepath: %s\n", filepath)
	fmt.Fprintf(loggingFile, "Sender address: %s\n", r.RemoteAddr)
	fmt.Fprintf(loggingFile, "Time: %s\n", time.Now().Format(loggingDateFormat))
	fmt.Fprintf(loggingFile, "========================================\n\n")
}

func LogError(err error) {
	fmt.Fprintf(loggingFile, "=================ERROR=================\n")
	fmt.Fprintf(loggingFile, "Error: %s\n", err.Error())
	fmt.Fprintf(loggingFile, "Time: %s\n", time.Now().Format(loggingDateFormat))
	fmt.Fprintf(loggingFile, "=======================================\n\n")
}

func LogFileRead(err error, filepath string) {
	fmt.Fprintf(loggingFile, "===============FILE READ===============\n")
	fmt.Fprintf(loggingFile, "Filepath: %s\n", filepath)
	fmt.Fprintf(loggingFile, "End cause: %s\n", err.Error())
	fmt.Fprintf(loggingFile, "Time: %s\n", time.Now().Format(loggingDateFormat))
	fmt.Fprintf(loggingFile, "=======================================\n\n")
}

func InitLogging() {
	const loggingFileDateFormat = "2006_01_02___15_04_05_MST"

	var logsDirPath = config.LOG_FILES_DIRPATH
	var timestamp = time.Now().Format(loggingFileDateFormat)
	var loggingPath = fmt.Sprintf("%slogs_%s", logsDirPath, timestamp)

	var loggingFileError error

	loggingFile, loggingFileError = os.Create(loggingPath)
	errors_handling.Handle(loggingFileError)
}
