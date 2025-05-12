package executor

import (
	"errors"

	"github.com/AmolKumarGupta/crona/global"
)

type MockExecutor struct {
	options *Options
}

func (m *MockExecutor) Run() error {
	_ = m.options

	if global.TestExecutorError {
		return errors.New("error")
	}

	return nil
}
