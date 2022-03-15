package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "master").Info("Hello")
	logger.WithField("username", "master").WithField("name", "Master").Info("Hello")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"username": "master",
		"name":     "Master",
	}).Info("Hello")
}