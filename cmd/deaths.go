// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

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
