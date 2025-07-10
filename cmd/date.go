package cmd

import (
	"fmt"
	"today/internal/logger"
	"today/internal/services"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Input a custom date.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			logger.LogErr(logger.MakeError("You must provide exactly 2 arguments (month and day)"))
			return
		}
		var month, day int
		n, err := fmt.Sscanf(args[0]+" "+args[1], "%d %d", &month, &day)
		if err != nil || n != 2 {
			logger.LogErr(logger.MakeError("Arguments must be integers (month day)"))
			return
		}
		if month > 12 || month < 1 {
			logger.LogErr(logger.MakeError("The month needs to be comprised between 1 and 12 (January and December)"))
		}

		if day > 31 || day < 1 {
			logger.LogErr(logger.MakeError("The day needs to be comprised between 1 and 31"))
		}

		services.Start(0, month, day)
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
