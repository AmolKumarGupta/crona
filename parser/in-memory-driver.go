package parser

import (
	"errors"

	"github.com/AmolKumarGupta/crona/job"
	"github.com/spf13/cobra"
)

var DefaultInMemoryTasks = []Task{
	*NewTask(
		NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{}),
		job.NewJob("pwd", []string{}),
	),
}

func init() {
	Resolver.Register("mem", func() Driver {
		return &InMemoryDriver{}
	})
}

type InMemoryDriver struct {
	Tasks []Task
}

func (d *InMemoryDriver) Init(_ *cobra.Command) error {
	d.Tasks = DefaultInMemoryTasks

	if len(d.Tasks) == 0 {
		return errors.New("no task in memroy driver")
	}

	return nil
}

func (d *InMemoryDriver) Parse() ([]Task, error) {
	return d.Tasks, nil
}
