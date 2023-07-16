package pool

import "time"

type Config struct {
	Concurrency  int           `json:"concurrency"`
	Timeout      time.Duration `json:"timeout"`
	MaxQueueSize int           `json:"waitQueueSize"`
}

func (c *Config) selfValidate() {
	if c.Concurrency <= 0 {
		c.Concurrency = 1
	}
	if c.Timeout < 0 {
		c.Timeout = 0
	}
	if c.MaxQueueSize < c.Concurrency {
		c.MaxQueueSize = c.Concurrency
	}
}
