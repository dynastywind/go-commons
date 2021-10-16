package pipeline

import "github.com/sirupsen/logrus"

type ErrorHandler interface {
	name() string
	handleError(config JobConfig, name, id, msg string, e error) *JobResult
}

type DefaultErrorHandler struct{}

func (handler DefaultErrorHandler) name() string {
	return "DefaultErrorHandler"
}

func (handler DefaultErrorHandler) handleError(config JobConfig, name, id, msg string, e error) *JobResult {
	if config.logError {
		logrus.WithField("name", name).WithField("id", id).Error(msg)
	}
	if !config.allowError {
		r := FailureResult(e, msg)
		return &r
	}
	return nil
}

func NewDefaultErrorHandler() DefaultErrorHandler {
	return DefaultErrorHandler{}
}
