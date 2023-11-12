package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	logger := logrus.New()

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.SetOutput(file)

	return logger
}
