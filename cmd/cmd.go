package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/AmolKumarGupta/crona/internal"
	"github.com/AmolKumarGupta/crona/parser"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringP("config", "c", "", "Path to the config file")
	rootCmd.Flags().StringP("log-level", "l", "error", "Log level (debug, info, warn, error)")
	// rootCmd.Flags().StringP("log-file", "f", "", "Path to the log file")
	// rootCmd.Flags().BoolP("daemon", "d", false, "Run as a daemon")
}

var rootCmd = &cobra.Command{
	Use:   "crona",
	Short: "Cron Advanced",
	Long:  `Crona is a experimental job scheduler`,
	Run: func(cmd *cobra.Command, args []string) {
		logLevel, err := logLevel(cmd)
		if err != nil {
			slog.Error(fmt.Sprintf("error getting log level: %s", err))
			os.Exit(1)
		}

		prevLevel := slog.SetLogLoggerLevel(logLevel)
		defer slog.SetLogLoggerLevel(prevLevel)

		fileDriver := &parser.FileDriver{}

		if err := fileDriver.Init(cmd); err != nil {
			slog.Error(fmt.Sprintf("error initializing file driver: %s", err))
			os.Exit(1)
		}

		tasks, err := fileDriver.Parse()

		if err != nil {
			slog.Error(fmt.Sprintf("error parsing config: %s", err))
			os.Exit(1)
		}

		tm := parser.GetTaskManager()
		for _, task := range tasks {
			tm.AddTask(task)
		}

		internal.NewCron().Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}
