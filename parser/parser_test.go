package parser

import (
	"testing"
	"time"

	"github.com/AmolKumarGupta/crona/job"
)

func TestGetTaskManager(t *testing.T) {
	tmInstance := GetTaskManager()
	defer ResetTaskManager()

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
	defer ResetTaskManager()

	tm.AddTask(*task)

	if len(tm.Tasks) != 1 {
		t.Errorf("unable to add task in task manager")
	}
}

func TestNext(t *testing.T) {
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
	defer ResetTaskManager()

	tm.AddTask(*task)

	now := time.Now()
	matchingTasks := tm.Next(now)

	if len(matchingTasks) != 1 {
		t.Errorf("expected 1 matching task, got %d", len(matchingTasks))
	}

	if !task.Compare(matchingTasks[0].ParseOptions) || !task.Job.Compare(matchingTasks[0].Job) {
		t.Errorf("expected matching task to be %v, got %v", *task, matchingTasks[0])
	}
}
