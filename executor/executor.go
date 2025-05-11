package executor

import "github.com/AmolKumarGupta/crona/global"

type Executor interface {
	Run() error
}

func New(setters ...Option) Executor {
	options := &Options{}

	for _, setter := range setters {
		setter(options)
	}

	if global.TestMode {
		return &MockExecutor{options}
	}

	return &RealExecutor{options}
}
