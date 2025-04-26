package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

func logLevel(cmd *cobra.Command) (slog.Level, error) {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return slog.LevelError, err
	}

	slogLevel := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	return slogLevel[logLevel], nil
}
