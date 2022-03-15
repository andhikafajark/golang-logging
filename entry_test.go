package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()

	entry := logrus.NewEntry(logger)

	entry.Info("Hello")
}
