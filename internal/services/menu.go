// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"today/internal/logger"
	"today/internal/utils"

	"github.com/charmbracelet/lipgloss"
)

func ShowDate(month, day int) {
	style := lipgloss.NewStyle()
	style = style.Bold(true).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#5c5cd4"))
	fmt.Println(style.Render(fmt.Sprintf("Chosen date: %s %02d", utils.IntToMonthName(month), day)))
	fmt.Println()
}

func StartWithOutput(autoStart, month, day int, outputFile string) {
	var selection int

	switch autoStart {
	case 0:
		if month != 0 && day != 0 {
			utils.MONTH = month
			utils.DAY = day
		}

		ShowDate(utils.MONTH, utils.DAY)

		_, selection = RunChooseMenu()
		selection++ // Easier for me to use 1, 2, 3, etc. instead of 0, 1, 2, etc.
	case 1, 2, 3:
		selection = autoStart
	default:
		logger.LogErr(logger.MakeError("Unknown selection chosen! Aborting!"))
	}

	switch selection {
	case -1:
		return
	case 1, 2, 3:
		var formattedEvents []string
		eventsFromAPI := GetEventsType(selection)
		for _, event := range eventsFromAPI {
			event.Text = strings.ReplaceAll(event.Text, "&nbsp;", " ")
			event.Text = strings.ReplaceAll(event.Text, "   ", " ")
			event.Text = strings.ReplaceAll(event.Text, "  ", " ")
			formattedEvents = append(formattedEvents, event.Text)
		}

		if outputFile != "" {
			err := saveEventsToFile(formattedEvents, outputFile)
			if err != nil {
				logger.LogErr(logger.MakeError("Failed to save output: " + err.Error()))
			} else {
				fmt.Printf("Saved output to %s\n", outputFile)
			}
			return
		}

		utils.ClearScreen()
		showEventList(formattedEvents)
	default:
		logger.MakeError("Unknown selection chosen! Aborting!")
	}
}

func saveEventsToFile(events []string, filename string) error {
	ext := "txt"
	if idx := strings.LastIndex(filename, "."); idx != -1 {
		ext = filename[idx+1:]
	}
	var content string
	switch ext {
	case "json":
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			return err
		}
		content = string(b)
	case "md":
		for _, e := range events {
			content += "- " + e + "\n"
		}
	default:
		for _, e := range events {
			content += e + "\n"
		}
	}
	return os.WriteFile(filename, []byte(content), 0644)
}
