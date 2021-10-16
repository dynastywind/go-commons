package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ConcurrentJob struct {
	id           string
	name         string
	defaultValue interface{}
	jobs         []Job
	config       JobConfig
	aggregator   Aggregator
	errorHandler ErrorHandler
	summary      Summary
}

func (concurrent ConcurrentJob) Do(ctx context.Context) JobResult {
	start := time.Now()
	r := concurrent.do(ctx)
	elapsed := time.Since(start).Milliseconds()
	concurrent.summary.summary(concurrent.id, concurrent.name, len(concurrent.jobs), concurrent.config, elapsed)
	return r
}

func (concurrent ConcurrentJob) do(ctx context.Context) JobResult {
	data := concurrent.defaultValue
	ch := make(chan JobResult, concurrent.config.maxConcurrency)
	length := len(concurrent.jobs)
	for index, job := range concurrent.jobs {
		go func(c context.Context, j Job) {
			defer func() {
				if re := recover(); re != nil {
					ch <- FailureResult(fmt.Errorf("concurrent job %v unexpected error: %v", concurrent.id, re), "Unexpected failure")
				}
			}()
			ch <- j.Do(c)
		}(ctx, job)
		if (index+1)%concurrent.config.maxConcurrency == 0 || index+1 == length {
			for i := 0; i <= index%concurrent.config.maxConcurrency; i++ {
				r := <-ch
				if r.Success {
					d, e := concurrent.aggregator(ctx, data, r.Data)
					if e != nil {
						if terminate := concurrent.errorHandler.handleError(concurrent.config, concurrent.name, concurrent.id,
							fmt.Sprintf("Concurrent job %v aggregation error: %v", concurrent.id, e.Error()), e); terminate != nil {
							return *terminate
						}
					} else {
						data = d
					}
				} else {
					if terminate := concurrent.errorHandler.handleError(concurrent.config, concurrent.name, concurrent.id,
						fmt.Sprintf("Concurrent job %v error: %v", concurrent.id, r.Message), r.Error); terminate != nil {
						return *terminate
					}
				}
			}
		}
	}
	return SuccessResultWithData(data)
}

func NewDefaultConcurrentJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator) ConcurrentJob {
	return NewConcurrentJob(name, defaultValue, jobs, aggregator, DefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultSummary())
}

func NewConcurrentJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator, config JobConfig,
	errorHandler ErrorHandler, summary Summary) ConcurrentJob {
	return ConcurrentJob{
		id:           uuid.New().String(),
		name:         name,
		defaultValue: defaultValue,
		jobs:         jobs,
		config:       config,
		aggregator:   aggregator,
		errorHandler: errorHandler,
		summary:      summary,
	}
}
