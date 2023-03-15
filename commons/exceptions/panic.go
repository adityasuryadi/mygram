package exceptions

import (
	"os"

	"github.com/sirupsen/logrus"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		logger := logrus.New()

		file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		logger.SetOutput(file)
		logger.Error(err)
		panic(err)
	}
}