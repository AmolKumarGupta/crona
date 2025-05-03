package parser

import (
	"sync"
	"time"

	"github.com/AmolKumarGupta/crona/job"
)

type Parser struct {
}

type Task struct {
	ParseOptions
	Job job.Job
}

var (
	once       sync.Once
	tmInstance *TaskManager
)

type TaskManager struct {
	Tasks []Task
}

func GetTaskManager() *TaskManager {
	once.Do(func() {
		tmInstance = &TaskManager{
			Tasks: []Task{},
		}
	})

	return tmInstance
}

func ResetTaskManager() {
	tmInstance = &TaskManager{
		Tasks: []Task{},
	}
}

func (tm *TaskManager) AddTask(task Task) {
	tm.Tasks = append(tm.Tasks, task)
}

// func (tm *TaskManager) RemoveTask(task Task) {
// 	for i, t := range tm.Tasks {
// 		if t == task {
// 			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
// 			break
// 		}
// 	}
// }

func (tm *TaskManager) Next(now time.Time) []Task {
	var curTasks []Task

	for _, task := range tm.Tasks {
		if task.MatchTime(now) {
			curTasks = append(curTasks, task)
		}
	}

	return curTasks
}
