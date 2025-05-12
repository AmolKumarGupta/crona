package executor_test

import (
	"io"
	"testing"

	"github.com/AmolKumarGupta/crona/executor"
	"github.com/AmolKumarGupta/crona/global"
)

func TestNew(t *testing.T) {
	cmd := executor.New(
		executor.Name("go"),
	)

	switch cmd.(type) {
	case *executor.RealExecutor:
	default:
		t.Errorf("executor.New should return *executor.RealExecutor")
	}
}

func TestNewMock(t *testing.T) {
	global.TestMode = true

	cmd := executor.New(
		executor.Name("go"),
		executor.Args([]string{}),
		executor.Stdout(&io.OffsetWriter{}),
		executor.Stderr(&io.OffsetWriter{}),
	)

	switch cmd.(type) {
	case *executor.MockExecutor:
	default:
		t.Errorf("executor.New should return *executor.MockExecutor")
	}
}
