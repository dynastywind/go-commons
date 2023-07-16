package pool

import "context"

type TaskResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

type Task interface {
	Run(ctx context.Context) *TaskResult
	Cancel(ctx context.Context)
	Callback(ctx context.Context, result *TaskResult)
}

type taskWrapper struct {
	id   string
	ctx  context.Context
	task Task
}

type TaskRunFunc func(ctx context.Context) *TaskResult

type TaskCallback func(ctx context.Context, result *TaskResult)

type TaskCancelFunc func(ctx context.Context)

type Taskable struct {
	RunFunc      TaskRunFunc
	CancelFunc   TaskCancelFunc
	CallbackFunc TaskCallback
}

func (t *Taskable) Run(ctx context.Context) *TaskResult {
	return t.RunFunc(ctx)
}

func (t *Taskable) Cancel(ctx context.Context) {
	if t.CancelFunc == nil {
		return
	}
	t.CancelFunc(ctx)
}

func (t *Taskable) Callback(ctx context.Context, result *TaskResult) {
	if t.CallbackFunc == nil {
		return
	}
	t.CallbackFunc(ctx, result)
}
