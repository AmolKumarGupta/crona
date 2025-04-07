package internal

import (
	"fmt"
	"time"
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

	for {
		var timer time.Timer = *time.NewTimer(time.Second)
		<-timer.C

		fmt.Printf("running at %s\n", time.Now())
		go func() {
			job := NewJob("php", []string{"/home/amol/workspace/go-space/crona/example/hello-world.php"})
			job.Run()
		}()
	}
}
