# GO-COMMONS

![go-commons workflow](https://github.com/dynastywind/go-commons/actions/workflows/go.yml/badge.svg)

This is a repository of commonly used utility functions of GOLANG, aiming to facilitate gopher's daily work.

## Import

```bash
go get github.com/dynastywind/go-commons
```

## Components

This repository is composed of different utility function modules. You can pick any one of them to use at your own service.

### Pipeline

This module contains a powerful tool to execute a series of jobs either sequentially or concurrently.

#### Usage

```go
pipeline.NewSequentialJob("sum", 0, jobs, sumAggregator, pipeline.DefaultJobConfig().WithAllowError(false),pipeline.NewDefaultErrorHandler(), pipeline.NewDefaultDigester()).Do(context.Background())

pipeline.NewConcurrentJob("sum", 0, jobs, sumAggregator, pipeline.DefaultJobConfig().WithAllowError(false),pipeline.NewDefaultErrorHandler(), pipeline.NewDefaultDigester()).Do(context.Background())
```

If you don't want to create an extra job struct, you can also opt to pass a Doable function alternatively. See below example.

```go
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
pipeline.NewDefaultSequentialJobWithDoable("sum", 0, doables, sumAggregator).Do(context.Background())
pipeline.NewDefaultConcurrentJobWithDoable("sum", 0, doables, sumAggregator).Do(context.Background())
```

By default, a job execution digest will be given during its execution process.

### Either

This module provides an Either type containing either one type or another.

#### Usage

```go
either.OfLeft(1).HashLeft()
either.OfRight(1).HasRight()
```

### Optional

This module provides an Optional utility function to avoid nil return type in GOLANG.

#### Usage

```go
optional.Of(1).IsPresent()
optional.OfEmpty().IsPresent()
```

### Structs

This module helps one to convert a struct to a **string-interface map** or vice verser.

#### Usage

Converting a struct to map

```go
type A struct {
    A int
}
m := structs.Map(A{A: 1})
```

Converting a map to struct

```go
type A struct {
    A int
}
s := structs.Struct(map[string]interface{}{"A": 1}, reflect.TypeOf(A{}))
```


## License

MIT
