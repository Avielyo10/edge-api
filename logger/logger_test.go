// FIXME: golangci-lint
// nolint:revive,typecheck
package logger_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/redhatinsights/edge-api/logger"

	"github.com/redhatinsights/edge-api/config"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestLogGenericErrorAndPanic(t *testing.T) {
	err := errors.New("generic error for testing")
	msg := "Generic error for testing"

	assert.PanicsWithError(
		t, err.Error(), func() {
			logger.LogErrorAndPanic(msg, err)
		})
}

func TestLogNilErrorDoesntPanic(t *testing.T) {
	msg := "Nil error for testing"

	assert.NotPanics(
		t, func() {
			logger.LogErrorAndPanic(msg, nil)
		})
}

func TestFlushLogger(t *testing.T) {
	var buffer bytes.Buffer

	logger.InitLogger(&buffer)

	want := "Test flushing log messages"
	myLog := log.WithFields(log.Fields{"app": "edge", "service": "images"})
	myLog.Info(want)

	logger.FlushLogger()
	got := buffer.String()
	assert.Contains(t, got, want)
}

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name     string
		logLevel string
		level    log.Level
	}{
		{
			name:     "Use DEBUG log level",
			logLevel: "DEBUG",
			level:    log.DebugLevel,
		},
		{
			name:     "Use ERROR log level",
			logLevel: "ERROR",
			level:    log.ErrorLevel,
		},
		{
			name:     "Use Info log level",
			logLevel: "INFO",
			level:    log.InfoLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalLogLevel := config.Get().LogLevel
			defer func(level string) {
				config.Get().LogLevel = level
			}(originalLogLevel)

			config.Get().LogLevel = tt.logLevel

			writer := &bytes.Buffer{}

			logger.InitLogger(writer)
			assert.Equal(t, tt.level, log.GetLevel())
			logger.FlushLogger()
		})
	}
}
