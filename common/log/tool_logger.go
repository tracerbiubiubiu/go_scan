package log

import (
	"fmt"
	"log"
	"os"
)

// 根据日志等级创建日志对象
var (
	traceLogger *log.Logger
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
)

const (
	LogLevelTrace int = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelUnknown
	TraceLogFileName = "Trace.txt"
	DebugLogFileName = "Debug.txt"
	InfoLogFileName  = "Info.txt"
	WarnLogFileName  = "Warn.txt"
	ErrorLogFileName = "Error.txt"
	FatalLogFileName = "Fatal.txt"
)

func init() {

}

//createLogFile 默认每次生成新的日志文件，删除原日志文件
func createLogFile(filePath string) error {

	return nil
}
func isValidPath(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		}
		return false, err
	}
	return true, nil
}
func createNewLogFile(filePath string) error {
	removeErr := os.Remove(filePath)
	if removeErr != nil {
		fmt.Println("delete old log file failed")
	}
	_, createErr := os.Create(filePath)
	if createErr != nil {
		fmt.Println("create new log file failed")
	}
}
