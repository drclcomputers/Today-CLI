// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"today/internal/services"
)

var outputFile string

var rootCmd = &cobra.Command{
	Use:   "today",
	Short: "Get a detailed overview of past events that have happened today.",
	Long: `Today is a tool that enables its users to enlarge their historical
knowledge by presenting them with a rich list of births, deaths, or events
that have taken place on today's date decades or even centuries ago.

For instance, did you know that Leonardo da Vinci, the famous
genius of the Renaissance era, was born on 15 April 1452? Today can help
you discover such facts with a simple command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("help") {
			return
		}

		services.StartWithOutput(0, 0, 0, outputFile)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Save output to file (txt, json, md supported by extension)")
}
