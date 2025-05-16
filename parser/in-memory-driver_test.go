package parser_test

import (
	"testing"

	"github.com/AmolKumarGupta/crona/job"
	"github.com/AmolKumarGupta/crona/parser"
	"github.com/spf13/cobra"
)

func TestInit(t *testing.T) {
	tests := []struct {
		input []parser.Task
	}{
		{
			[]parser.Task{
				*parser.NewTask(
					parser.NewParseOptions("*", "*", "*", "*", "*", "*", []parser.Flag{}),
					job.NewJob("pwd", []string{}),
				),
			},
		},
		{
			[]parser.Task{
				*parser.NewTask(
					parser.NewParseOptions("*", "*", "*", "*", "*", "*", []parser.Flag{}),
					job.NewJob("pwd", []string{}),
				),
				*parser.NewTask(
					parser.NewParseOptions("*", "*", "*/2", "*", "*", "*", []parser.Flag{}),
					job.NewJob("pwd", []string{}),
				),
			},
		},
	}

	cmd := &cobra.Command{}

	for _, test := range tests {
		t.Run("Init", func(t *testing.T) {
			parser.DefaultInMemoryTasks = test.input

			driver := &parser.InMemoryDriver{}
			err := driver.Init(cmd)
			if err != nil {
				t.Errorf("InMemoryDriver.Init() throw error %v", err)
			}

			tasks, _ := driver.Parse()

			if len(tasks) != len(test.input) {
				t.Errorf("InMemoryDriver.Task length is %d, want %d", len(tasks), len(test.input))
			}
		})
	}
}

func TestParse(t *testing.T) {
	parser.DefaultInMemoryTasks = []parser.Task{}

	cmd := &cobra.Command{}

	driver := &parser.InMemoryDriver{}
	err := driver.Init(cmd)

	if err == nil {
		t.Errorf("empty tasks should throw error but it didnt")
	}
}
