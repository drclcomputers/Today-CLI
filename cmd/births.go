package cmd

import (
	"today/internal/services"

	"github.com/spf13/cobra"
)

var birthsCmd = &cobra.Command{
	Use:     "births",
	Short:   "Show only births that have happened in the past.",
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		services.Start(2, 0, 0)
	},
}

func init() {
	rootCmd.AddCommand(birthsCmd)
}
