package parser

import (
	"errors"

	"github.com/AmolKumarGupta/crona/job"
)

var DefaultInMemoryTasks = []Task{
	*NewTask(
		NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{}),
		job.NewJob("ls", []string{"-l"}),
	),
}

type InMemoryDriver struct {
	Tasks []Task
}

func (d *InMemoryDriver) Init() error {
	d.Tasks = DefaultInMemoryTasks

	if len(d.Tasks) == 0 {
		return errors.New("no task in memroy driver")
	}

	return nil
}

func (d *InMemoryDriver) Parse() ([]Task, error) {
	return d.Tasks, nil
}
