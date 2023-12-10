package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehutchllew/template.ts/cmd/models"
)

type WizardAnswers struct {
	*models.UserAnswers
}

func (w *WizardAnswers) Init() tea.Cmd {
	return nil
}

func (w *WizardAnswers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		}
	}

	return w, nil
}

func (w *WizardAnswers) View() string {
	return "hello wizard"
}
