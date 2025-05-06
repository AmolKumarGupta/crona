package job

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
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

func (j Job) Compare(other Job) bool {
	return j.command == other.command &&
		len(j.args) == len(other.args) &&
		(len(j.args) == 0 || strings.Join(j.args, " ") == strings.Join(other.args, " "))
}
