package parser

import "github.com/AmolKumarGupta/crona/internal"

type Parser struct {
}

type Task struct {
	ParseOptions
	Job internal.Job
}
