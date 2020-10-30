package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(uninstallCommand)
}

var uninstallCommand = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a node",
	Run:   func(cmd *cobra.Command, args []string) {},
}
