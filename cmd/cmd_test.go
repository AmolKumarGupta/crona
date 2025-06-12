package cmd

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/AmolKumarGupta/crona/job"
	"github.com/AmolKumarGupta/crona/parser"
)

func TestRootCmd(t *testing.T) {
	output := &bytes.Buffer{}
	job.SetJobStdOut(output)

	cmd := exec.Command("hostname")
	buf, berr := cmd.CombinedOutput()
	if berr != nil {
		t.Fatalf("failed to execute command: %v", berr)
	}

	parser.DefaultInMemoryTasks = []parser.Task{
		*parser.NewTask(
			parser.NewParseOptions("*", "*", "*", "*", "*", "*", []parser.Flag{}),
			job.NewJob("hostname", []string{}),
		),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rootCmd.SetArgs([]string{"-r", "mem"})
	rootCmd.SetContext(ctx)

	err := rootCmd.Execute()
	if err != nil && ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected no error or context deadline exceeded, got %v", err)
	}

	outputStr := output.String()

	expectedHostname := string(buf)
	if strings.HasPrefix(outputStr, expectedHostname) == false {
		t.Errorf("expected output to start with %s, got %s", expectedHostname, outputStr)
	}
}
