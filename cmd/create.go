package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create entities",
	Long:  `Create entities from the commandline. See subcommands.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
