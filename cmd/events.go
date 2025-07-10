package cmd

import (
	"today/internal/services"

	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:     "events",
	Short:   "Show only events that have happened in the past.",
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		services.Start(1, 0, 0)
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}
