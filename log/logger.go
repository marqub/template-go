package log

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// type Fields struct {
// 	time string
// 	corr string
// 	appn string
// }

// type Logger interface {
// }

// type logger struct {
// 	logger Logger
// }
// hh
// var Logger = logrus.New()

// func init() {
// 	// Log as JSON instead of the default ASCII formatter
// 	Logger.SetFormatter(&log.JSONFormatter{})

// 	// Output to stdout instead of the default stderr
// 	// Can be any io.Writer, see below for File example
// 	*Logger.SetOutput(os.Stdout)

func Logger() *log.Entry {
	return log.WithFields(log.Fields{"appn": "template-go", "date": time.Now().UTC().Format(time.RFC3339)})
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
}
