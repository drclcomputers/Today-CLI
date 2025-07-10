/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the app to the latest version.",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			fmt.Println("The included bash script only works on Linux and UNIX-like OSes. For Windows, update manually.")
		} else {
			err := os.Chmod("update.sh", 0755)
			if err != nil {
				fmt.Println("Failed to make script executable:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
