package infrastructure

import (
	"github.com/Sirupsen/logrus"
	"github.com/weekface/mgorus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	hooker, err := mgorus.NewHooker("localhost:27017", "db", "collection")
	if err == nil {
		logger.Hooks.Add(hooker)
	}
}

func GetLogger() *logrus.Logger {
	return logger
}

