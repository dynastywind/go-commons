package pipeline

type JobResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   error       `json:"-"`
	Data    interface{} `json:"data"`
}

func SuccessResultWithData(data interface{}) JobResult {
	return JobResult{
		Success: true,
		Data:    data,
	}
}

func SuccessResult() JobResult {
	return SuccessResultWithData(nil)
}

func FailureResult(e error, msg string) JobResult {
	return JobResult{
		Success: false,
		Message: msg,
		Error:   e,
	}
}
