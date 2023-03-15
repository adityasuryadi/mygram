package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func WriteLog(errData interface{}) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.Error(errData)
}