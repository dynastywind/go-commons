package pool

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/dynastywind/go-commons/utils"
	"github.com/google/uuid"
)

type RoutinePool interface {
	SubmitTask(ctx context.Context, task Task) (string, error)
	SubmitTaskRun(ctx context.Context, run TaskRunFunc) (string, error)
}

type defaultRoutinePool struct {
	config          *Config
	execChan        chan *taskWrapper
	quitChan        chan bool
	cancelChan      chan string
	runningRoutines *int32
	tasks           map[string]*taskWrapper
}

func NewDefaultRoutinePool(config *Config) RoutinePool {
	var initRoutines int32 = 0
	config.selfValidate()
	pool := &defaultRoutinePool{
		config:          config,
		execChan:        make(chan *taskWrapper, config.MaxQueueSize),
		quitChan:        make(chan bool),
		cancelChan:      make(chan string),
		runningRoutines: &initRoutines,
		tasks:           make(map[string]*taskWrapper),
	}
	utils.GoWithoutCtx(pool.exec, func(r interface{}) {
		fmt.Printf("recovered %s", r)
	})
	return pool
}

func (pool *defaultRoutinePool) SubmitTask(ctx context.Context, task Task) (string, error) {
	if len(pool.execChan) >= pool.config.MaxQueueSize {
		return "", fmt.Errorf("task denied")
	}
	id, _ := uuid.NewUUID()
	idStr := id.String()
	pool.execChan <- &taskWrapper{
		id:   idStr,
		ctx:  ctx,
		task: task,
	}
	return idStr, nil
}

func (pool *defaultRoutinePool) SubmitTaskRun(ctx context.Context, run TaskRunFunc) (string, error) {
	task := &struct{ Taskable }{}
	task.RunFunc = run
	return pool.SubmitTask(ctx, task)
}

func (pool *defaultRoutinePool) exec() {
	for {
		select {
		case <-pool.quitChan:
			close(pool.execChan)
			return
		case wrapper := <-pool.execChan:
			for size := atomic.LoadInt32(pool.runningRoutines); size >= int32(pool.config.Concurrency); size = atomic.LoadInt32(pool.runningRoutines) {
			}
			atomic.AddInt32(pool.runningRoutines, 1)
			pool.tasks[wrapper.id] = wrapper
			resultChan := make(chan *TaskResult, 1)
			utils.Go(wrapper.ctx, func(ctx context.Context) {
				defer atomic.AddInt32(pool.runningRoutines, -1)
				utils.GoWithoutCtx(func() {
					resultChan <- wrapper.task.Run(wrapper.ctx)
				}, nil)
				select {
				case result := <-resultChan:
					wrapper.task.Callback(wrapper.ctx, result)
					close(resultChan)
				case <-time.After(pool.config.Timeout * time.Second):
					fmt.Println("timeout")
				}
			}, nil)
		case id := <-pool.cancelChan:
			if wrapper, found := pool.tasks[id]; found {
				wrapper.task.Cancel(wrapper.ctx)
				delete(pool.tasks, wrapper.id)
			}
		}
	}
}
