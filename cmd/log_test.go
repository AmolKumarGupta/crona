package cmd

import (
	"log/slog"
	"testing"
)

func TestLogLevel(t *testing.T) {
	err := rootCmd.Flags().Set("log-level", "info")
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	level, err := logLevel(rootCmd)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	if level != slog.LevelInfo {
		t.Errorf("expected log level info, got %v", level)
	}
}
