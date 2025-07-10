package services

import (
	"fmt"
	"os"
	"today/internal/utils"

	spinnerpkg "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type spinnerModel struct {
	spinner spinnerpkg.Model
	done    bool
	message string
}

func (m spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "esc" {
			m.done = true
			os.Exit(0)
			return m, tea.Quit
		}
	case spinnerpkg.TickMsg:
		if m.done {
			return m, tea.Quit
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m spinnerModel) View() string {
	if m.done {
		return ""
	}
	return fmt.Sprintf("%s %s\n", m.spinner.View(), m.message)
}

func showLoadingSpinner(message string, fetchFunc func() utils.ZenquotesTodayAPIResp) utils.ZenquotesTodayAPIResp {
	resultChan := make(chan utils.ZenquotesTodayAPIResp)

	go func() {
		result := fetchFunc()
		resultChan <- result
	}()

	spinner := spinnerpkg.New()
	spinner.Spinner = spinnerpkg.Dot
	spinner.Style = spinner.Style.Bold(true).Foreground(lipgloss.Color("#5c5cd4"))
	m := spinnerModel{spinner: spinner, message: message}

	var result utils.ZenquotesTodayAPIResp
	done := make(chan struct{})

	go func() {
		result = <-resultChan
		m.done = true
		close(done)
	}()

	p := tea.NewProgram(m)

	go func() {
		_ = p.Start()
	}()

	<-done
	p.Quit()
	return result
}
