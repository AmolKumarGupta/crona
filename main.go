package main

import (
	"github.com/AmolKumarGupta/crona/cmd"

	_ "embed"
)

//go:embed VERSION
var VersionFile []byte

func main() {
	cmd.SetVersion(string(VersionFile))
	cmd.Execute()
}
