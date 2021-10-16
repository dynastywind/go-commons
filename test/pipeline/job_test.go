package pipeline_test

import (
	"context"
	"fmt"

	"github.com/dynastywind/go-commons/pipeline"
)

type SumJob struct {
	value int
}

func (job SumJob) Do(ctx context.Context) pipeline.JobResult {
	return pipeline.SuccessResultWithData(job.value)
}

var sumAggregator = func(ctx context.Context, prior, current interface{}) (interface{}, error) {
	if p, ok := prior.(int); ok {
		if c, ok := current.(int); ok {
			return p + c, nil
		} else {
			return nil, fmt.Errorf("Current value type error: not an int")
		}
	}
	return nil, fmt.Errorf("Prior value type error: not an int")
}

type ErrorJob struct{}

func (job ErrorJob) Do(ctx context.Context) pipeline.JobResult {
	return pipeline.FailureResult(fmt.Errorf("error"), "error")
}

type PanicJob struct{}

func (job PanicJob) Do(ctx context.Context) pipeline.JobResult {
	panic("panic job")
}
