package executor

import "io"

type Options struct {
	name   string
	args   []string
	stdout io.Writer
	stderr io.Writer
}

type Option func(*Options)

func Name(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

func Args(args []string) Option {
	return func(o *Options) {
		o.args = args
	}
}

func Stdout(stdout io.Writer) Option {
	return func(o *Options) {
		o.stdout = stdout
	}
}

func Stderr(stderr io.Writer) Option {
	return func(o *Options) {
		o.stderr = stderr
	}
}
