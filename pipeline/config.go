package pipeline

import "fmt"

const (
	DefaultMaxConcurrency = 40
)

type JobConfig struct {
	allowError     bool
	logError       bool
	summary        bool
	maxConcurrency int
}

func (config JobConfig) WithAllowError(allowError bool) JobConfig {
	config.allowError = allowError
	return config
}

func (config JobConfig) WithLogError(logError bool) JobConfig {
	config.logError = logError
	return config
}

func (config JobConfig) WithMaxConcurrency(maxConcurrency int) JobConfig {
	if maxConcurrency > 0 && maxConcurrency <= DefaultMaxConcurrency {
		config.maxConcurrency = maxConcurrency
	}
	return config
}

func (config JobConfig) WithSummary(summary bool) JobConfig {
	config.summary = summary
	return config
}

func (config JobConfig) String() string {
	return fmt.Sprintf("AllowError: %v\nLogError: %v\nSummary: %v\nMaxConcurrency: %v\n",
		config.allowError, config.logError, config.summary, config.maxConcurrency)
}

func DefaultJobConfig() JobConfig {
	return JobConfig{
		allowError:     true,
		logError:       true,
		summary:        true,
		maxConcurrency: DefaultMaxConcurrency,
	}
}
