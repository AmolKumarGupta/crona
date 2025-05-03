package cmd

import (
	"log/slog"
	"testing"
)

func TestLogLevel(t *testing.T) {
	rootCmd.Flags().Set("log-level", "info")

	level, err := logLevel(rootCmd)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	if level != slog.LevelInfo {
		t.Errorf("expected log level info, got %v", level)
	}
}
