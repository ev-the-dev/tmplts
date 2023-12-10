package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ehutchllew/template.ts/cmd/models"
)

type WizardAnswers struct {
	models.UserAnswers

	height int
	width  int
}

func New(questionnaire *models.UserAnswers) *WizardAnswers {
	logFile, logFileErr := tea.LogToFile("debug.log", "debug")
	if logFileErr != nil {
		log.Fatalf("log file err: %v", logFileErr)
	}
	defer logFile.Close()

	w := WizardAnswers{}
	w.AppName = questionnaire.AppName
	w.EsBuild = questionnaire.EsBuild
	w.EsLint = questionnaire.EsLint
	w.Jest = questionnaire.Jest
	w.Swc = questionnaire.Swc
	w.Typescript = questionnaire.Typescript

	p := tea.NewProgram(w)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error starting bubbletea: %v", err)
	}

	return &w
}

func (w WizardAnswers) Init() tea.Cmd {
	return nil
}

func (w WizardAnswers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.height = msg.Height
		w.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		}
	}

	return w, nil
}

func (w WizardAnswers) View() string {
	if w.width == 0 {
		return "Loading..."
	}
	return lipgloss.JoinVertical(lipgloss.Center)
}
