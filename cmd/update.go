/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
			fmt.Println("Starting update...")
			err := os.Chmod("update.sh", 0755)
			if err != nil {
				fmt.Println("Failed to make script executable:", err)
			} else {
				fmt.Println("Indexing the update script...")
			}
			exePath, err := os.Executable()
			if err != nil {
				log.Fatalf("Failed to get executable path: %v", err)
			}
			exeDir := filepath.Dir(exePath)
			scriptPath := filepath.Join(exeDir, "update.sh")
			cmd := exec.Command(scriptPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
