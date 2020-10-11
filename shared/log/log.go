package log

import (
	"fmt"
	"runtime"

	logger "github.com/sirupsen/logrus"
)

func Error(str string) {
	source := getSource()
	logger.Error(source, str)
}

func Errorf(str string, args ...interface{}) {
	source := getSource()
	logger.Errorf(fmt.Sprintf("%s%s", source, str), args...)
}

func Info(str string) {
	logger.Info(str)
}

func Infof(str string, args ...interface{}) {
	logger.Infof(str, args...)
}

func getSource() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("[%s:%d] ", file, line)
	}
	return ""
}
