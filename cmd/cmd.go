package cmd

import (
	"fmt"
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
		fmt.Println(err)
		os.Exit(1)
	}
}
