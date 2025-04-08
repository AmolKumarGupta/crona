package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/AmolKumarGupta/crona/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crona",
	Short: "Cron Advanced",
	Long:  `Crona is a experimental scheduler manager`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.NewCron().Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}
