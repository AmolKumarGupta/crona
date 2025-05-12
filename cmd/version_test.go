package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	version := "1.0.0"
	SetVersion(version)

	out := new(bytes.Buffer)

	versionCmd.SetOut(out)
	rootCmd.SetArgs([]string{"version"})

	err := rootCmd.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if out.String() != version+"\n" {
		t.Errorf("expected %s, got %s", version, out.String())
	}
}
