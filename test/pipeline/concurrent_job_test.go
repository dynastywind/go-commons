package pipeline_test

import (
	"context"

	"github.com/dynastywind/go-commons/pipeline"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Concurrent job tests", func() {
	ginkgo.Context("Concurrent job test", func() {
		ginkgo.When("Executing a concurrent sum job", func() {
			ginkgo.It("Should return correct sum value", func() {
				jobs := []pipeline.Job{
					SumJob{
						value: 1,
					},
					SumJob{
						value: 2,
					},
					SumJob{
						value: 3,
					},
					SumJob{
						value: 4,
					},
				}
				r := pipeline.NewDefaultConcurrentJob("sum", 0, jobs, sumAggregator).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(r.Data).To(gomega.Equal(10))
			})
			ginkgo.It("Should ignore error and return remaining sum value", func() {
				jobs := []pipeline.Job{
					SumJob{
						value: 1,
					},
					SumJob{
						value: 2,
					},
					ErrorJob{},
					SumJob{
						value: 4,
					},
				}
				r := pipeline.NewDefaultConcurrentJob("sum", 0, jobs, sumAggregator).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(r.Data).To(gomega.Equal(7))
			})
			ginkgo.It("Should ignore error and return initial value", func() {
				jobs := []pipeline.Job{
					ErrorJob{},
				}
				r := pipeline.NewDefaultConcurrentJob("sum", 0, jobs, sumAggregator).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(r.Data).To(gomega.Equal(0))
			})
			ginkgo.It("Should recover from panic and ignore it, return remaining sum value", func() {
				jobs := []pipeline.Job{
					SumJob{
						value: 1,
					},
					SumJob{
						value: 2,
					},
					PanicJob{},
					SumJob{
						value: 4,
					},
				}
				r := pipeline.NewDefaultConcurrentJob("sum", 0, jobs, sumAggregator).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(r.Data).To(gomega.Equal(7))
			})
			ginkgo.It("Should return error result", func() {
				jobs := []pipeline.Job{
					SumJob{
						value: 1,
					},
					SumJob{
						value: 2,
					},
					ErrorJob{},
					SumJob{
						value: 4,
					},
				}
				r := pipeline.NewConcurrentJob("sum", 0, jobs, sumAggregator, pipeline.DefaultEarlyStopper, pipeline.NewDefaultJobConfig().WithAllowError(false),
					pipeline.NewDefaultErrorHandler(), pipeline.NewDefaultDigester()).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeFalse())
			})
			ginkgo.It("Should be early stopped", func() {
				jobs := []pipeline.Job{
					SumJob{
						value: 1,
					},
					SumJob{
						value: 2,
					},
					SumJob{
						value: 3,
					},
					SumJob{
						value: 4,
					},
				}
				var earlyStopped bool
				r := pipeline.NewConcurrentJob("sum", 0, jobs, sumAggregator, func(ctx context.Context, cur, aggr interface{}) bool {
					if d, ok := cur.(int); ok {
						if d == 3 {
							earlyStopped = true
							return true
						}
					}
					return false
				}, pipeline.NewDefaultJobConfig(), pipeline.NewDefaultErrorHandler(), pipeline.NewDefaultDigester()).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(earlyStopped).To(gomega.BeTrue())
			})
		})
		ginkgo.When("Executing a concurrent doable job", func() {
			ginkgo.It("Should return correct sum value", func() {
				doables := []pipeline.Doable{
					func(ctx context.Context) pipeline.JobResult {
						return pipeline.SuccessResultWithData(1)
					},
					func(ctx context.Context) pipeline.JobResult {
						return pipeline.SuccessResultWithData(2)
					},
					func(ctx context.Context) pipeline.JobResult {
						return pipeline.SuccessResultWithData(3)
					},
					func(ctx context.Context) pipeline.JobResult {
						return pipeline.SuccessResultWithData(4)
					},
				}
				r := pipeline.NewDefaultConcurrentJobWithDoable("sum", 0, doables, sumAggregator).Do(context.Background())
				gomega.Expect(r.Success).To(gomega.BeTrue())
				gomega.Expect(r.Data).To(gomega.Equal(10))
			})
		})
	})
})
