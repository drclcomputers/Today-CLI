// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"fmt"
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

func Start(autoStart, month, day int) {
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

		utils.ClearScreen()

		showEventList(formattedEvents)
	default:
		logger.MakeError("Unknown selection chosen! Aborting!")
	}
}
