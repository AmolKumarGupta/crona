package parser

type Driver interface {
	Init() error
	Parse() ([]Task, error)
}
