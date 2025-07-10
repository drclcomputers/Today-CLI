package cmd

import (
	"today/internal/services"

	"github.com/spf13/cobra"
)

var deathsCmd = &cobra.Command{
	Use:     "deaths",
	Short:   "Show only deaths that have happened in the past.",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		services.Start(3, 0, 0)
	},
}

func init() {
	rootCmd.AddCommand(deathsCmd)
}
