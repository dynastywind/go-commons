package pipeline

import (
	"github.com/sirupsen/logrus"
)

type Summary interface {
	summary(id, name string, jobs int, config JobConfig, elapsed int64)
}

type DefaultSummary struct{}

func (summary DefaultSummary) summary(id, name string, jobs int, config JobConfig, elapsed int64) {
	if config.summary {
		logrus.WithField("id", id).WithField("name", name).WithField("jobs count", jobs).
			WithField("config", config.String()).WithField("elapsed", elapsed).
			Info("Job summary details")
	}
}

func NewDefaultSummary() DefaultSummary {
	return DefaultSummary{}
}
