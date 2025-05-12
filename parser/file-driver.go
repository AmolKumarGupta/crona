package parser

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/AmolKumarGupta/crona/job"
	"github.com/spf13/cobra"
)

type FileDriver struct {
	FilePath string
}

func (f *FileDriver) Init(cmd *cobra.Command) error {
	if cmd.Flag("config").Changed {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		f.FilePath = configPath
	}

	if f.FilePath != "" {
		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return errors.New("unable to get current working directory")
	}

	f.FilePath = fmt.Sprintf("%s/%s", dir, ".crona")

	if _, err := os.Stat(f.FilePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("file does not exist")
	}

	return nil
}

func (f *FileDriver) Parse() ([]Task, error) {
	slog.Info("current config", "file", f.FilePath)

	bytes, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, errors.New("file does not exist")
	}

	var tasks []Task

	list := f.split(string(bytes))

	for _, str := range list {
		task, err := f.asTask(str)
		if err != nil {
			slog.Warn("invalid string", "str", err)
			continue
		}

		tasks = append(tasks, *task)
	}

	return tasks, nil
}

// support single line comment
func (f *FileDriver) isComment(str string) bool {
	return strings.HasPrefix(str, "//")
}

// takes the text file as string and
// return only list of string that look like command
func (f *FileDriver) split(text string) []string {
	listOfStrings := strings.Split(text, "\n")

	list := []string{}

	for _, l := range listOfStrings {
		l = strings.Trim(l, " ")
		if f.isComment(l) {
			continue
		}

		if strings.TrimSpace(l) == "" {
			continue
		}

		list = append(list, l)
	}

	return list
}

func (f *FileDriver) asTask(str string) (*Task, error) {
	chunk := strings.Split(str, " ")
	if len(chunk) < 7 {
		return nil, errors.New("invalid task")
	}

	opts := &ParseOptions{}

	for i, v := range chunk {
		switch i {
		case 0:
			_, err := SecondBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Second = v

		case 1:
			_, err := MinuteBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Minute = v
		case 2:
			_, err := HourBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Hour = v
		case 3:
			_, err := DomBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Dom = v
		case 4:
			_, err := MonthBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Month = v
		case 5:
			_, err := DowBound.Validate(v)
			if err != nil {
				return nil, err
			}
			opts.Dow = v
		}
	}

	cmdString := strings.TrimSpace(strings.Join(chunk[6:], " "))
	if cmdString == "" {
		return nil, errors.New("command is empty")
	}

	if strings.HasPrefix(cmdString, "*") {
		return nil, errors.New("command is invalid")
	}

	job := job.NewJob(chunk[6], chunk[7:])

	task := &Task{
		ParseOptions: *opts,
		Job:          *job,
	}

	return task, nil
}
