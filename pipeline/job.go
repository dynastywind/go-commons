package pipeline

import "context"

type Job interface {
	// Do returns job execution result
	//
	// @params ctx Execution context
	// @return Job execution result
	Do(ctx context.Context) JobResult
}

// Aggregator aggregates a job's result with an accumulated result, and returns this final result
//
// @param ctx Execution context
// @param prior Accumulated result
// @param current Current job's result, which can be diffrent from prior
// @return The first value is the job result to be returned, which should be the same type as prior.
// @return The second value is a potential error
type Aggregator func(ctx context.Context, prior, current interface{}) (interface{}, error)

// Doable is a sysnonym to Job interface's Do function, which helps to construc jobs in a funcational way
type Doable func(ctx context.Context) JobResult

type DoableJob struct {
	doable Doable
}

func (job DoableJob) Do(ctx context.Context) JobResult {
	return job.doable(ctx)
}
