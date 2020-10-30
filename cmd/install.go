package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(installCommand)
}

var installCommand = &cobra.Command{
	Use:   "install",
	Short: "Install a node",
	Run:   func(cmd *cobra.Command, args []string) {},
}
