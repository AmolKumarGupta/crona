package parser

import (
	"fmt"
	"testing"

	"github.com/AmolKumarGupta/crona/job"
)

func TestIsComment(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"// This is a comment", true},
		{"//This is a comment", true},
		{"//", true},
		{"# Another comment", false},
		{"This is not a comment", false},
		{"", false},
	}

	fd := &FileDriver{}

	for _, test := range tests {
		t.Run("isComment", func(t *testing.T) {
			got := fd.isComment(test.line)

			if got != test.want {
				t.Errorf("fd.isComment('%s') = %v; want %v", test.line, got, test.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{
			"* * * * * * php main.php\n*/5 * * * * * ./sample.sh",
			[]string{"* * * * * * php main.php", "*/5 * * * * * ./sample.sh"},
		},
		{
			"",
			[]string{},
		},
		{
			"* * * * * * php main.php\n\n//  \n",
			[]string{"* * * * * * php main.php"},
		},
		{
			"* * * * * * php main.php\n\n\n*/5 * * * * * ./sample.sh",
			[]string{"* * * * * * php main.php", "*/5 * * * * * ./sample.sh"},
		},
	}

	for _, test := range tests {
		t.Run("split", func(t *testing.T) {
			fd := &FileDriver{test.text}
			got := fd.split(test.text)

			if len(got) != len(test.want) {
				t.Errorf("fd.split('%s') = %v; want %v", test.text, got, test.want)
				return
			}

			for i, v := range got {
				if v != test.want[i] {
					t.Errorf("fd.split('%s')[%d] = '%s'; want '%s'", test.text, i, v, test.want[i])
				}
			}
		})
	}
}

func TestAsTask(t *testing.T) {
	tests := []struct {
		input string
		want  Task
	}{
		{
			"* * * * * * php main.php",
			*NewTask(
				NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{}),
				job.NewJob("php", []string{"main.php"}),
			),
		},
		{
			"0 0 */2 * * * ./sample.sh",
			*NewTask(
				NewParseOptions("0", "0", "*/2", "*", "*", "*", []Flag{}),
				job.NewJob("./sample.sh", []string{}),
			),
		},
	}

	for _, test := range tests {
		t.Run("AsTask", func(t *testing.T) {
			fd := &FileDriver{}
			got, err := fd.asTask(test.input)
			if err != nil {
				t.Error(err)
			}

			if !got.Compare(test.want.ParseOptions) {
				t.Errorf("fd.asTask('%s') = %v; want %v", test.input, got, test.want)
			}

			if !got.Job.Compare(test.want.Job) {
				t.Errorf("fd.asTask('%s') = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestAsTaskFailed(t *testing.T) {
	tests := []struct {
		input string
		err   bool
	}{
		{"* * * * * * php main.php", false},
		{"0 0 */2 * * ./sample.sh", true},
		{"a * * * * * php main.php", true},
		{"* a * * * * php main.php", true},
		{"* * a * * * php main.php", true},
		{"* * * a * * php main.php", true},
		{"* * * * a * php main.php", true},
		{"* * * * * a php main.php", true},
		{"* * * * * *", true},
		{"* * * * * * ", true},
		{"* * * * * * *", true},
		{"* * * * * * * php main.php", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("AsTask '%s'", test.input), func(t *testing.T) {
			fd := &FileDriver{}
			task, err := fd.asTask(test.input)

			if test.err && err == nil {
				t.Errorf("it should throw error, %v", task)

			} else if !test.err && err != nil {
				t.Errorf("it should not throw error, %v", task)
			}
		})
	}
}
