package cmode

import (
	"net/http"

	"github.com/klosmo/atlas-app-toolkit/v2/cmode/logger"
	"github.com/sirupsen/logrus"
)

func ExampleCMode() {
	appLogger := logrus.StandardLogger()

	cmLogger := logger.New(appLogger)
	cm := New(appLogger, cmLogger)

	http.Handle("/", cm.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		appLogger.Fatalf("Server fatal error - %s", err)
	}
}
