// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

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
		services.StartWithOutput(2, 0, 0, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(birthsCmd)
}
