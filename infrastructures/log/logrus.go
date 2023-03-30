package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Logger{
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	return logger
}

func WriteLog(errData interface{}) {
		logger := NewLog()
		logger.Error(errData)
}