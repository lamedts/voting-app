package logger

import (
	"io"
	"os"
	"path"
	"strings"
	"zakkaya/config"

	"io/ioutil"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

var logLevel logrus.Level
var primaryOutStream io.Writer

var (
	appLogHook   logrus.Hook
	warnLogHook  logrus.Hook
	errorLogHook logrus.Hook

	coreLogHook logrus.Hook
)

func init() {
	setPrimaryOutStream()

	// General hooks
	appLogHook = lfshook.NewHook(
		getPathMap(path.Join(config.LOG_PATH, path.Base(os.Args[0])+".log")),
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

	warnLogHook = lfshook.NewHook(
		lfshook.PathMap{
			logrus.WarnLevel:  path.Join(config.LOG_PATH, "warn.log"),
			logrus.ErrorLevel: path.Join(config.LOG_PATH, "warn.log"),
			logrus.FatalLevel: path.Join(config.LOG_PATH, "warn.log"),
			logrus.PanicLevel: path.Join(config.LOG_PATH, "warn.log"),
		},
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

	errorLogHook = lfshook.NewHook(
		lfshook.PathMap{
			logrus.ErrorLevel: path.Join(config.LOG_PATH, "error.log"),
			logrus.FatalLevel: path.Join(config.LOG_PATH, "error.log"),
			logrus.PanicLevel: path.Join(config.LOG_PATH, "error.log"),
		},
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

	// Module hooks
	coreLogHook = lfshook.NewHook(
		getPathMap(path.Join(config.LOG_PATH, "core.log")),
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

}

func setPrimaryOutStream() {
	zakkayaEnv := os.Getenv("ENV")
	switch strings.ToUpper(zakkayaEnv) {
	case "DEBUG", "TEST":
		logLevel = logrus.DebugLevel
		primaryOutStream = os.Stdout
	default:
		logLevel = logrus.InfoLevel
		primaryOutStream = ioutil.Discard // abandone output
	}
}

func getPathMap(logPath string) lfshook.PathMap {
	return lfshook.PathMap{
		logrus.DebugLevel: logPath,
		logrus.InfoLevel:  logPath,
		logrus.WarnLevel:  logPath,
		logrus.ErrorLevel: logPath,
		logrus.FatalLevel: logPath,
		logrus.PanicLevel: logPath,
	}
}

func GetLogger(module string) *logrus.Entry {
	logger := logrus.New()
	logger.Formatter = &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true}
	logger.SetLevel(logLevel)
	setPrimaryOutStream()

	logger.AddHook(appLogHook)
	logger.AddHook(warnLogHook)
	logger.AddHook(errorLogHook)

	return logger.WithField("prefix", module)
}
