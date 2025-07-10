// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func RunChooseMenu() (string, int) {
	p := tea.NewProgram(initialMenuModel())

	finalModel, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Error running program:", err)
		return "", -1
	}
	m, ok := finalModel.(menuModel)
	if !ok || !m.confirmed || m.selected >= len(m.choices) {
		return "", -1
	}

	if m.choices[m.selected] == "Quit" {
		return "", -1
	}
	return m.choices[m.selected], m.selected
}
