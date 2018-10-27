package cmd

import (
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate things",
	Long:  `Generate things from the commandline. See subcommands.`,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
