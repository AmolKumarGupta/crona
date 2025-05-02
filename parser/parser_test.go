package parser

import (
	"testing"

	"github.com/AmolKumarGupta/crona/job"
)

func TestGetTaskManager(t *testing.T) {
	tmInstance := GetTaskManager()
	tmInstance2 := GetTaskManager()

	if tmInstance != tmInstance2 {
		t.Errorf("GetTaskManager does not return singleton instance")
	}
}

func TestAddTask(t *testing.T) {
	po := &ParseOptions{
		Second: "*",
		Minute: "*",
		Hour:   "*",
		Dom:    "*",
		Month:  "*",
		Dow:    "*",
	}

	job := job.NewJob("go", []string{"main.go"})

	task := &Task{
		*po,
		*job,
	}

	tm := GetTaskManager()
	tm.AddTask(*task)

	if len(tm.Tasks) != 1 {
		t.Errorf("unable to add task in task manager")
	}
}
