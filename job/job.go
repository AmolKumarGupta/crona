package job

import (
	"io"
	"os"
	"strings"

	"github.com/AmolKumarGupta/crona/executor"
)

var (
	stdOut io.Writer = os.Stdout
)

func SetJobStdOut(w io.Writer) {
	stdOut = w
}

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

func (j *Job) Run() error {
	cmd := executor.New(
		executor.Name(j.command),
		executor.Args(j.args),
		executor.Stdout(stdOut),
		executor.Stderr(os.Stderr),
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (j Job) Compare(other Job) bool {
	return j.command == other.command &&
		len(j.args) == len(other.args) &&
		(len(j.args) == 0 || strings.Join(j.args, " ") == strings.Join(other.args, " "))
}
