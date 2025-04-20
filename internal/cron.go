package internal

import (
	"fmt"
	"time"

	"github.com/AmolKumarGupta/crona/parser"
)

type Cron struct {
	running bool
}

func NewCron() *Cron {
	return &Cron{
		running: false,
	}
}

func (c *Cron) Start() {
	if c.running {
		return
	}

	tm := parser.GetTaskManager()

	for {
		var timer = *time.NewTimer(time.Second)
		<-timer.C

		tasks := tm.Next()

		if len(tasks) == 0 {
			continue
		}

		fmt.Printf("running at %s\n", time.Now())
		for _, task := range tasks {
			go func() {
				task.Job.Run()
			}()
		}
	}
}
