package pipeline

import "fmt"

const (
	DefaultMaxConcurrency = 40
)

type JobConfig struct {
	allowError     bool
	logError       bool
	digest         bool
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

func (config JobConfig) WithDigest(digest bool) JobConfig {
	config.digest = digest
	return config
}

func (config JobConfig) String() string {
	return fmt.Sprintf("AllowError: %v\nLogError: %v\nDigest: %v\nMaxConcurrency: %v\n",
		config.allowError, config.logError, config.digest, config.maxConcurrency)
}

func NewDefaultJobConfig() JobConfig {
	return JobConfig{
		allowError:     true,
		logError:       true,
		digest:         true,
		maxConcurrency: DefaultMaxConcurrency,
	}
}
