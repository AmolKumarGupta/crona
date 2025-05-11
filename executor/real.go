package executor

import "os/exec"

type RealExecutor struct {
	options *Options
}

func (r *RealExecutor) Run() error {
	cmd := exec.Command(r.options.name, r.options.args...)
	cmd.Stdout = r.options.stdout
	cmd.Stderr = r.options.stderr

	return cmd.Run()
}
