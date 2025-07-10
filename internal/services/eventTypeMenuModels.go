package services

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	choices   []string
	selected  int
	confirmed bool
}

func initialMenuModel() menuModel {
	return menuModel{
		choices:  []string{"Events", "Births", "Deaths", "Quit"},
		selected: 0,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.choices)-1 {
				m.selected++
			}
		case "enter", " ":
			m.confirmed = true
			return m, tea.Quit
			//case "q", "esc":
			//	os.Exit(0)
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "Select which type of historical events you'd like to view:\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.selected == i {
			cursor = ">"
			style := lipgloss.NewStyle()
			style = style.Bold(true).
				Italic(true).
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#5c5cd4"))
			s += style.Render(fmt.Sprintf("%s %s", cursor, choice))
		} else {
			s += fmt.Sprintf("%s %s", cursor, choice)
		}
		s += "\n"
	}
	s += "\nUse ↑/↓ or j/k to move, Enter/Space to select."
	return s
}
