package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, info := range bi.Settings {
			if info.Key == "vcs.revision" {
				version = info.Value[0:7]
			}
		}
	}

	rootCmd.AddCommand(versionCmd)
}
