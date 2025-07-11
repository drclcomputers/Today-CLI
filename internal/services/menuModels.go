// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"fmt"
	"io"
	"math/rand/v2"
	"slices"
	"time"
	"today/internal/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type customDelegate struct{}

func (d customDelegate) Height() int                               { return 2 }
func (d customDelegate) Spacing() int                              { return 0 }
func (d customDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d customDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	e, ok := item.(eventItem)
	if !ok {
		return
	}
	style := lipgloss.NewStyle()
	if index == m.Index() {
		style = style.Bold(true).
			Italic(true).
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#5c5cd4"))
	}
	fmt.Fprint(w, style.Width(m.Width()).Render(e.title))
}

type eventItem struct {
	title string
}

func (e eventItem) Title() string       { return e.title }
func (e eventItem) Description() string { return "" }
func (e eventItem) FilterValue() string { return e.title }

type eventListModel struct {
	list list.Model
}

func (m eventListModel) Init() tea.Cmd {
	return nil
}

func (m eventListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "c":
			if item, ok := m.list.SelectedItem().(eventItem); ok {
				utils.CopyToClipboard(item.title)
				m.list.StatusMessageLifetime = 4 * time.Second
				return m, m.list.NewStatusMessage("Copied to clipboard!")
			}
		case "n":
			if m.list.ShowStatusBar() == false {
				m.list.SetShowStatusBar(!false)
			} else {
				m.list.SetShowStatusBar(false)
			}
		case "r":
			randIndex := rand.IntN(len(m.list.Items()))
			m.list.Select(randIndex)
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m eventListModel) View() string {
	return m.list.View() + "\nPress c to copy the selected text or q to quit."
}

func showEventList(formattedEvents []string) {
	var items []list.Item
	for _, e := range formattedEvents {
		items = append(items, eventItem{title: e})
	}

	slices.Reverse(items)

	width, height := utils.GetTerminalSize()

	l := list.New(items, customDelegate{}, int(float64(width)), height)
	l.Title = fmt.Sprintf("Chosen date: %s %02d", utils.IntToMonthName(utils.MONTH), utils.DAY)
	l.SetShowTitle(true)
	l.SetShowStatusBar(false)
	//l.SetShowHelp(false)
	//l.SetShowPagination(false)
	l.SetStatusBarItemName("event", "events")

	m := eventListModel{list: l}

	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Println("Error running event list:", err)
	}
}
