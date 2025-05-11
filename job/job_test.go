package job

import (
	"testing"

	"github.com/AmolKumarGupta/crona/global"
)

func TestCompare(t *testing.T) {
	job1 := NewJob("go", []string{"main.go"})
	job2 := NewJob("go", []string{"main.go"})
	job3 := NewJob("go", []string{"sample.go"})

	if !job1.Compare(*job2) {
		t.Error("job1 should be same as job2")
	}

	if job1.Compare(*job3) {
		t.Error("job1 should be not same as job3")
	}

	job4 := NewJob("go", []string{"main.go", "--test"})
	job5 := NewJob("go", []string{"main.go", "--fake"})
	if job4.Compare(*job5) {
		t.Error("job4 should be not same as job5")
	}
}

func TestRun(t *testing.T) {
	global.TestMode = true

	job := NewJob("php", []string{"main.php"})

	err := job.Run()
	if err != nil {
		t.Error("job should be run successfully, but it didn't")
	}
}

func TestRunFailed(t *testing.T) {
	global.TestMode = true
	global.TestExecutorError = true

	job := NewJob("", []string{"main.php"})

	err := job.Run()
	if err == nil {
		t.Error("invalid job should throw error")
	}
}
