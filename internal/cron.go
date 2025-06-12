package internal

import (
	"context"
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

// Accepts a context which allows cancellation to stop the loop.
func (c *Cron) Start(ctx context.Context) {
	if c.running {
		return
	}
	c.running = true
	tm := parser.GetTaskManager()

	for {
		select {
		case <-ctx.Done():
			slog.Info("Cron stopped")
			return
		default:
			slog.Debug("tick")

			timer := time.NewTimer(time.Second)
			now := <-timer.C

			tasks := tm.Next(now)

			if len(tasks) == 0 {
				continue
			}

			slog.Info("running")
			for _, task := range tasks {
				task := task
				go func() {
					if err := task.Job.Run(); err != nil {
						slog.Error(fmt.Sprintf("running job: %s", err))
					}
				}()
			}
		}
	}
}
