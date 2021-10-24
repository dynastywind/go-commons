package pipeline

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Digester interface {

	// whenJobStarts execute when a job starts
	//
	// @param id Job ID
	// @param name Job name
	// @param jobs Job count
	// @param config Job config
	whenJobStarts(id, name string, jobs int, config JobConfig)

	// whenJobEnds execute when a job ends
	//
	// @param id Job ID
	// @param name Job name
	// @param jobs Job count
	// @param config Job config
	// @param elapsed Time elapsed during job execution
	whenJobEnds(id, name string, jobs int, config JobConfig, elapsed int64)

	// whenChildJobStarts execute when a job's child job starts
	//
	// @param id Job ID
	// @param name Job name
	// @param index Child job index
	// @param config Job config
	whenChildJobStarts(id, name string, index int, config JobConfig)

	// whenChildJobEnds execute when a job's child job ends
	//
	// @param id Job ID
	// @param name Job name
	// @param index Child job index
	// @param config Job config
	// @param elapsed Time elapsed during job execution
	whenChildJobEnds(id, name string, index int, config JobConfig, elapsed int64)

	// whenEarlyStopped execute when a job is early stopped
	//
	// @param id Job ID
	// @param name Job name
	// @param config Job config
	whenEarlyStopped(id, name string, config JobConfig)
}

type DefaultDigester struct {
}

func (digester DefaultDigester) whenJobStarts(id, name string, jobs int, config JobConfig) {
	if config.digest {
		logrus.WithField("id", id).WithField("name", name).WithField("jobs count", jobs).
			WithField("config", config.String()).
			Info(fmt.Sprintf("Job %v starts", id))
	}
}

func (digester DefaultDigester) whenJobEnds(id, name string, jobs int, config JobConfig, elapsed int64) {
	if config.digest {
		logrus.WithField("id", id).WithField("name", name).WithField("jobs count", jobs).
			WithField("config", config.String()).WithField("elapsed", elapsed).
			Info(fmt.Sprintf("Job %v ends, time elapsed: %vms", id, elapsed))
	}
}

func (digester DefaultDigester) whenChildJobStarts(id, name string, index int, config JobConfig) {
	if config.digest {
		logrus.WithField("id", id).WithField("name", name).WithField("index", index).
			WithField("config", config.String()).
			Info(fmt.Sprintf("Child job %v of %v starts", index, id))
	}
}

func (digester DefaultDigester) whenChildJobEnds(id, name string, index int, config JobConfig, elapsed int64) {
	if config.digest {
		logrus.WithField("id", id).WithField("name", name).WithField("index", index).
			WithField("config", config.String()).WithField("elapsed", elapsed).
			Info(fmt.Sprintf("Child job %v of %v ends, time elapsed: %vms", index, id, elapsed))
	}
}

func (digester DefaultDigester) whenEarlyStopped(id, name string, config JobConfig) {
	if config.digest {
		logrus.WithField("id", id).WithField("name", name).
			WithField("config", config).Info(fmt.Sprintf("Job %v is early stopped", id))
	}
}

func NewDefaultDigester() DefaultDigester {
	return DefaultDigester{}
}
