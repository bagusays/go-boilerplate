package log

import (
	"flag"
	"fmt"
	"go-boilerplate/shared/config"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type TDRModel struct {
	AppName        string      `json:"app"`
	AppVersion     string      `json:"ver"`
	IP             string      `json:"ip"`
	LogTime        string      `json:"logTime"`
	Port           int         `json:"port"`
	SrcIP          string      `json:"srcIP"`
	RespTime       int64       `json:"rt"`
	Path           string      `json:"path"`
	Header         interface{} `json:"header"` // better to pass data here as is, don't cast it to string. use map or array
	Request        interface{} `json:"req"`
	Response       interface{} `json:"resp"`
	ResponseCode   string      `json:"rc"`
	Error          string      `json:"error"`
	ThreadID       string      `json:"threadID"`
	AdditionalData interface{} `json:"addData"`
}

type logger struct {
	log *logrus.Logger
	tdr *logrus.Logger
}

func initializeLogger(category string, cfg config.Log) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	if cfg.IsWriteToFile {
		if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
			return log
		}
		today := time.Now().Local().Format("2006-01-02")
		logFile, err := os.OpenFile(fmt.Sprintf("logs/%s-%s.log", category, today), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file for output: %s", err)
		}
		log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	}

	return log
}

func NewLogger(cfg config.Log) *logger {
	syslog := initializeLogger("syslog", cfg)
	tdr := initializeLogger("tdr", cfg)

	return &logger{
		log: syslog,
		tdr: tdr,
	}
}

func (l *logger) Error(message string, err error, additionalField ...logrus.Fields) {
	log := l.log.WithFields(getSource())
	log = log.WithField("error", err.Error())
	for _, i := range additionalField {
		log = log.WithFields(i)
	}
	log.Error(message)
}

func (l *logger) Info(message string, additionalField ...logrus.Fields) {
	log := l.log.WithFields(getSource())
	for _, i := range additionalField {
		log = log.WithFields(i)
	}
	log.Error(message)
}

func (l *logger) Fatal(message string, err error, additionalField ...logrus.Fields) {
	log := l.log.WithFields(getSource())
	log = log.WithField("error", err.Error())
	for _, i := range additionalField {
		log = log.WithFields(i)
	}
	log.Error(message)
}

func (l *logger) TDR(tdr TDRModel) {
	l.tdr.WithField("data", tdr).Info()
}

func getSource() logrus.Fields {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return logrus.Fields{"source": fmt.Sprintf("[%s:%d]", file, line)}
	}
	return logrus.Fields{"source": nil}
}
