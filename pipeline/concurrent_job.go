package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ConcurrentJob executes child jobs concurrently
// It early stopps when a certain criteria is met and returns the aggregated result up to that time
type ConcurrentJob struct {
	id           string
	name         string
	defaultValue interface{}
	jobs         []Job
	config       JobConfig
	aggregator   Aggregator
	earlyStopper EarlyStopper
	errorHandler ErrorHandler
	digester     Digester
}

func (concurrent ConcurrentJob) Do(ctx context.Context) JobResult {
	jobCount := len(concurrent.jobs)
	concurrent.digester.whenJobStarts(concurrent.id, concurrent.name, jobCount, concurrent.config)
	start := time.Now()
	r := concurrent.do(ctx)
	elapsed := time.Since(start).Milliseconds()
	concurrent.digester.whenJobEnds(concurrent.id, concurrent.name, jobCount, concurrent.config, elapsed)
	return r
}

func (concurrent ConcurrentJob) do(ctx context.Context) JobResult {
	data := concurrent.defaultValue
	ch := make(chan JobResult, concurrent.config.maxConcurrency)
	length := len(concurrent.jobs)
	for index, job := range concurrent.jobs {
		go func(c context.Context, i int, j Job) {
			begin := time.Now()
			defer func() {
				passed := time.Since(begin).Microseconds()
				concurrent.digester.whenChildJobEnds(concurrent.id, concurrent.name, i, concurrent.config, passed)
				if re := recover(); re != nil {
					ch <- FailureResult(fmt.Errorf("concurrent job %v unexpected error: %v", concurrent.id, re), "Unexpected failure")
				}
			}()
			concurrent.digester.whenChildJobStarts(concurrent.id, concurrent.name, i, concurrent.config)
			ch <- j.Do(c)
		}(ctx, index, job)
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
						if concurrent.earlyStopper(ctx, r.Data, data) {
							concurrent.digester.whenEarlyStopped(concurrent.id, concurrent.name, concurrent.config)
							return SuccessResultWithData(data)
						}
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
	return NewConcurrentJob(name, defaultValue, jobs, aggregator, DefaultEarlyStopper, NewDefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultDigester())
}

func NewConcurrentJob(name string, defaultValue interface{}, jobs []Job, aggregator Aggregator, earlyStopper EarlyStopper, config JobConfig,
	errorHandler ErrorHandler, digester Digester) ConcurrentJob {
	return ConcurrentJob{
		id:           uuid.New().String(),
		name:         name,
		defaultValue: defaultValue,
		jobs:         jobs,
		config:       config,
		aggregator:   aggregator,
		earlyStopper: earlyStopper,
		errorHandler: errorHandler,
		digester:     digester,
	}
}

func NewDefaultConcurrentJobWithDoable(name string, defaultValue interface{}, doables []Doable, aggregator Aggregator) ConcurrentJob {
	return NewConcurrentJobWithDoable(name, defaultValue, doables, aggregator, DefaultEarlyStopper, NewDefaultJobConfig(), NewDefaultErrorHandler(), NewDefaultDigester())
}

func NewConcurrentJobWithDoable(name string, defaultValue interface{}, doables []Doable, aggregator Aggregator, earlyStopper EarlyStopper, config JobConfig,
	errorHandler ErrorHandler, digester Digester) ConcurrentJob {
	var jobs []Job
	for _, doable := range doables {
		jobs = append(jobs, DoableJob{
			doable: doable,
		})
	}
	return NewConcurrentJob(name, defaultValue, jobs, aggregator, earlyStopper, config, errorHandler, digester)
}
