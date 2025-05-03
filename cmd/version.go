package cmd

import (
	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(Version)
		cmd.Println(Version)
	},
}

func SetVersion(version string) {
	Version = version
}
