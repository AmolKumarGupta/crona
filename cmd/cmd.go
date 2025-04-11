package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/AmolKumarGupta/crona/internal"
	"github.com/AmolKumarGupta/crona/parser"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crona",
	Short: "Cron Advanced",
	Long:  `Crona is a experimental scheduler manager`,
	Run: func(cmd *cobra.Command, args []string) {

		fileDriver := &parser.FileDriver{}
		if err := fileDriver.Init(); err != nil {
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
			// slog.Info("task", "task", task)
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
