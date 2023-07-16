package utils

import "context"

func Go(ctx context.Context, run func(ctx context.Context), rec func(ctx context.Context, r interface{})) {
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil && rec != nil {
				rec(ctx, r)
			}
		}()
		run(ctx)
	}(ctx)
}
func GoWithoutCtx(run func(), rec func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil && rec != nil {
				rec(r)
			}
		}()
		run()
	}()
}
