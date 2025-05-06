package job

import "testing"

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
