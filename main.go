package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dynastywind/go-commons/pool"
)

type simpleTask struct {
	id int
}

func (t *simpleTask) Callback(ctx context.Context, result *pool.TaskResult) {
	fmt.Println(result.Data)
}

func (t *simpleTask) Cancel(ctx context.Context) {}

func (t *simpleTask) Run(ctx context.Context) *pool.TaskResult {
	time.Sleep(2 * time.Second)
	return &pool.TaskResult{
		Success: true,
		Data:    fmt.Sprintf("task %d done", t.id),
	}
}

func main() {
	p := pool.NewDefaultRoutinePool(&pool.Config{
		Concurrency:  2,
		Timeout:      10,
		MaxQueueSize: 100,
	})
	for i := 0; i < 10; i++ {
		id, e := p.SubmitTask(context.Background(), &simpleTask{
			id: i,
		})
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Printf("%s is submitted\n", id)
		}
	}
	time.Sleep(20 * time.Second)
}
