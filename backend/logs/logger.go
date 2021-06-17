package logs

import (
	"homefill/backend/config"

	"github.com/sirupsen/logrus"
)

func LogIt(logCase int, funcName, errMsg string, err error) {
	switch logCase {
	case LogFatal:
		config.Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Fatal(errMsg)

	case LogWarn:
		config.Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Warn(errMsg)

	case LogInfo:
		config.Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Info(errMsg)
	}
}
