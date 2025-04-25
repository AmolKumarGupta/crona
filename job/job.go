package job

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

type Job struct {
	command string
	args    []string
}

func NewJob(command string, args []string) *Job {
	return &Job{
		command: command,
		args:    args,
	}
}

func (j *Job) Run() {
	cmd := exec.Command(j.command, j.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		slog.Error(fmt.Sprintf("running job: %s:", err))
		return
	}
}
