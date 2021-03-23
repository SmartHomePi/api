package logger

import (
	"io"
	"os"
	"strings"

	"github.com/SmartHomePi/api/pkg/config"
	"github.com/op/go-logging"
)

const logModule = `SmartHomePi`

var logInstance = logging.MustGetLogger(logModule)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfile} [%{level}] ▶%{color:reset} %{message}`,
)

func InitLogger() {

	if !config.LogsEnabled.GetBool() {
		return
	}

	logInstance.ExtraCalldepth = 1

	// if config.LogsEnabled.GetBool() {
	// 	err := os.Mkdir(config.LogsFilePath.GetString(), 0744)
	// 	if err != nil && !os.IsExist(err) {
	// 		Fatalf("Could not create log folder: %s", err.Error())
	// 	}
	// }

	// fileWriter := GetLogWriter()

	level, err := logging.LogLevel(strings.ToUpper(config.LogsLevel.GetString()))
	if err != nil {
		Fatalf("Error setting database log level: %s", err.Error())
	}

	stdoutBackend := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), format)
	// fileBackend := logging.NewBackendFormatter(logging.NewLogBackend(fileWriter, "", 0), format)

	stdoutBackendLeveled := logging.AddModuleLevel(stdoutBackend)
	// fileBackendLeveled := logging.AddModuleLevel(fileBackend)

	stdoutBackendLeveled.SetLevel(level, logModule)
	// fileBackendLeveled.SetLevel(level, logModule)

	logInstance.SetBackend(stdoutBackendLeveled)
	// logInstance.SetBackend(fileBackendLeveled)

}

// GetLogWriter returns the writer to where the normal log goes, depending on the config
func GetLogWriter() (writer io.Writer) {

	fullLogFilePath := config.LogsFilePath.GetString() + "/" + config.LogsFileName.GetString()
	f, err := os.OpenFile(fullLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		Fatalf("Could not create logfile %s: %s", fullLogFilePath, err.Error())
	}
	writer = f

	return
}

func GetLogger() *logging.Logger {
	return logInstance
}

// Debug is for debug messages
func Debug(args ...interface{}) {
	logInstance.Debug(args...)
}

// Debugf is for debug messages
func Debugf(format string, args ...interface{}) {
	logInstance.Debugf(format, args...)
}

// Info is for info messages
func Info(args ...interface{}) {
	logInstance.Info(args...)
}

// Infof is for info messages
func Infof(format string, args ...interface{}) {
	logInstance.Infof(format, args...)
}

// Error is for error messages
func Error(args ...interface{}) {
	logInstance.Error(args...)
}

// Errorf is for error messages
func Errorf(format string, args ...interface{}) {
	logInstance.Errorf(format, args...)
}

// Warning is for warning messages
func Warning(args ...interface{}) {
	logInstance.Warning(args...)
}

// Warningf is for warning messages
func Warningf(format string, args ...interface{}) {
	logInstance.Warningf(format, args...)
}

// Critical is for critical messages
func Critical(args ...interface{}) {
	logInstance.Critical(args...)
}

// Criticalf is for critical messages
func Criticalf(format string, args ...interface{}) {
	logInstance.Criticalf(format, args...)
}

// Fatal is for fatal messages
func Fatal(args ...interface{}) {
	logInstance.Fatal(args...)
}

// Fatalf is for fatal messages
func Fatalf(format string, args ...interface{}) {
	logInstance.Fatalf(format, args...)
}
