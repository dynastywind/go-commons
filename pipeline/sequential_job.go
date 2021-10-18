package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SequentialJob struct {
	id           string
	name         string
	jobs         []Job
	config       JobConfig
	defaultValue interface{}
	aggregator   Aggregator
	errorHandler ErrorHandler
	digester     Digester
}

func (sequential SequentialJob) Do(ctx context.Context) JobResult {
	jobCount := len(sequential.jobs)
	sequential.digester.whenJobStarts(sequential.id, sequential.name, jobCount, sequential.config)
	start := time.Now()
	r := sequential.do(ctx)
	elapsed := time.Since(start).Milliseconds()
	sequential.digester.whenJobEnds(sequential.id, sequential.name, jobCount, sequential.config, elapsed)
	return r
}

func (sequential SequentialJob) do(ctx context.Context) JobResult {
	ch := make(chan JobResult)
	go func(c context.Context, jobs []Job) {
		var data = sequential.defaultValue
		defer func() {
			if re := recover(); re != nil {
				if r := sequential.errorHandler.handleError(sequential.config, sequential.name, sequential.id,
					"Unexpected failure", fmt.Errorf("sequential job %v unexpected error: %v", sequential.id, re)); r != nil {
					ch <- *r
				} else {
					ch <- SuccessResultWithData(data)
				}
			}
		}()
		for _, job := range sequential.jobs {
			r := job.Do(ctx)
			if r.Success {
				d, e := sequential.aggregator(ctx, data, r.Data)
				if e != nil {
					if terminate := sequential.errorHandler.handleError(sequential.config, sequential.name, sequential.id,
						fmt.Sprintf("Sequential job %v aggregation error: %v", sequential.id, e.Error()), e); terminate != nil {
						ch <- *terminate
						return
					}
				} else {
					data = d
				}
			} else {
				if terminate := sequential.errorHandler.handleError(sequential.config, sequential.name, sequential.id,
					fmt.Sprintf("Sequential job %v error: %v", sequential.id, r.Message), r.Error); terminate != nil {
					ch <- *terminate
					return
				}
			}
		}
		ch <- SuccessResultWithData(data)
	}(ctx, sequential.jobs)
	return <-ch
}

func NewDefaultSequentialJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator) SequentialJob {
	return NewSequentialJob(name, defaultValue, jobs, aggregator, NewDefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultDigester())
}

func NewSequentialJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator, config JobConfig,
	errorHandler ErrorHandler, digester Digester) SequentialJob {
	return SequentialJob{
		id:           uuid.New().String(),
		name:         name,
		jobs:         jobs,
		config:       config,
		defaultValue: defaultValue,
		aggregator:   aggregator,
		errorHandler: errorHandler,
		digester:     digester,
	}
}

func NewDefaultSequentialJobWithDoable(name string, defaultValue interface{}, doables []Doable, aggregator Aggregator) SequentialJob {
	return NewSequentialJobWithDoable(name, defaultValue, doables, aggregator, NewDefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultDigester())
}

func NewSequentialJobWithDoable(name string, defaultValue interface{}, doables []Doable, aggregator Aggregator, config JobConfig,
	errorHandler ErrorHandler, digester Digester) SequentialJob {
	var jobs []Job
	for _, doable := range doables {
		jobs = append(jobs, DoableJob{
			doable: doable,
		})
	}
	return NewSequentialJob(name, defaultValue, jobs, aggregator, config, errorHandler, digester)
}
