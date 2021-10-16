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
	summary      Summary
}

func (sequential SequentialJob) Do(ctx context.Context) JobResult {
	start := time.Now()
	r := sequential.do(ctx)
	elapsed := time.Since(start).Milliseconds()
	sequential.summary.summary(sequential.id, sequential.name, len(sequential.jobs), sequential.config, elapsed)
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
	return NewSequentialJob(name, defaultValue, jobs, aggregator, DefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultSummary())
}

func NewSequentialJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator, config JobConfig,
	errorHandler ErrorHandler, summary Summary) SequentialJob {
	return SequentialJob{
		id:           uuid.New().String(),
		name:         name,
		jobs:         jobs,
		config:       config,
		defaultValue: defaultValue,
		aggregator:   aggregator,
		errorHandler: errorHandler,
		summary:      summary,
	}
}
