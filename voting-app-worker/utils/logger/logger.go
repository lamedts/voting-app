package logger

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"voting-app/voting-app-worker/config"

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
	logPath := config.AppConfig.LogDir
	setPrimaryOutStream()

	// General hooks
	appLogHook = lfshook.NewHook(
		getPathMap(path.Join(logPath, "app.log")),
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

	warnLogHook = lfshook.NewHook(
		lfshook.PathMap{
			logrus.WarnLevel:  path.Join(logPath, "warn.log"),
			logrus.ErrorLevel: path.Join(logPath, "warn.log"),
			logrus.FatalLevel: path.Join(logPath, "warn.log"),
			logrus.PanicLevel: path.Join(logPath, "warn.log"),
		},
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

	errorLogHook = lfshook.NewHook(
		lfshook.PathMap{
			logrus.ErrorLevel: path.Join(logPath, "error.log"),
			logrus.FatalLevel: path.Join(logPath, "error.log"),
			logrus.PanicLevel: path.Join(logPath, "error.log"),
		},
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true},
	)

}

func setPrimaryOutStream() {
	env := os.Getenv("ENV")
	switch strings.ToUpper(env) {
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
