package internal

import (
	"fmt"
	"log/slog"
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
		slog.Debug("tick")

		var timer = *time.NewTimer(time.Second)
		now := <-timer.C

		tasks := tm.Next(now)

		if len(tasks) == 0 {
			continue
		}

		slog.Info("running")
		for _, task := range tasks {
			go func() {
				err := task.Job.Run()
				if err != nil {
					slog.Error(fmt.Sprintf("running job: %s:", err))
				}
			}()
		}
	}
}
