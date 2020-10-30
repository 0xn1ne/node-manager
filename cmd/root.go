package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "nman",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		root(cmd, args)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("example command.")
}
