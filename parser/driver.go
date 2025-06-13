package parser

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Resolver = &DriverResolver{
	resolves: make(map[string]Resolve),
}

type Driver interface {
	Init(cmd *cobra.Command) error
	Parse() ([]Task, error)
}

type Resolve func() Driver

type DriverResolver struct {
	resolves map[string]Resolve
}

func (r *DriverResolver) Register(key string, cb Resolve) {
	r.resolves[key] = cb
}

func (r DriverResolver) Get(key string) (Resolve, bool) {
	cb, exists := r.resolves[key]
	return cb, exists
}

func Get(key string) (Driver, error) {
	cb, exists := Resolver.Get(key)
	if !exists {
		return nil, fmt.Errorf("driver with name '%s' does not exists", key)
	}

	return cb(), nil
}
