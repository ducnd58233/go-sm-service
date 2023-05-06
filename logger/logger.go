package logger

import (
    "fmt"
    "os"

    "github.com/sirupsen/logrus"
    "github.com/rifflock/lfshook"
)

var Log *logrus.Logger

func init() {
    Log = logrus.New()
    Log.SetFormatter(&logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
        FullTimestamp:   true,
        ForceColors:     true, // enable colors for terminal output
    })

    // Output to stdout instead of the default stderr
    Log.SetOutput(os.Stdout)

    // Set log level to info by default
    Log.SetLevel(logrus.InfoLevel)

    // Add hooks for rotating file logs
    logDir := "logs"
    logFile := "app.log"
    logHook := lfshook.NewHook(lfshook.PathMap{
        logrus.InfoLevel:  fmt.Sprintf("%s/%s", logDir, logFile),
        logrus.ErrorLevel: fmt.Sprintf("%s/%s", logDir, logFile),
    }, &logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
        FullTimestamp:   true,
    })
    Log.AddHook(logHook)
}

// SetLogLevel sets the log level
func SetLogLevel(level string) error {
    logLevel, err := logrus.ParseLevel(level)
    if err != nil {
        return err
    }
    Log.SetLevel(logLevel)
    return nil
}

// Info logs a message with level info
func Info(args ...interface{}) {
    Log.Info(args...)
}

// Error logs a message with level error
func Error(args ...interface{}) {
    Log.Error(args...)
}