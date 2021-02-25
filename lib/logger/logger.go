package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

}

//SetLogLevel  set log level
func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

//SetReportCaller set reprot caller
func SetReportCaller(bool bool) {
	log.SetReportCaller(bool)
}
