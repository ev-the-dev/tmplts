package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ev-the-dev/tmplts/cmd/models"
)

type WizardAnswers struct {
	appName                  string
	appNameTextInput         textinput.Model
	appNameTextInputRendered bool

	cursor int

	selected map[int]bool
	styles   *Styles

	height int
	width  int
}

func New() *models.UserAnswers {
	logFile, logFileErr := tea.LogToFile("debug.log", "debug")
	if logFileErr != nil {
		log.Fatalf("log file err: %v", logFileErr)
	}
	defer logFile.Close()

	w := WizardAnswers{
		appName: "",
		selected: map[int]bool{
			models.ES_BUILD:   false,
			models.ES_LINT:    false,
			models.JEST:       false,
			models.PRETTIER:   false,
			models.SWC:        false,
			models.TYPESCRIPT: false,
		},
		styles: DefaultStyles(),
	}

	p := tea.NewProgram(w)
	m, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error starting bubbletea: %v", err)
	}

	updatedW := m.(WizardAnswers)

	userAnswers := &models.UserAnswers{
		AppName:    updatedW.appName,
		EsBuild:    updatedW.selected[models.ES_BUILD],
		EsLint:     updatedW.selected[models.ES_LINT],
		Jest:       updatedW.selected[models.JEST],
		Prettier:   updatedW.selected[models.PRETTIER],
		Swc:        updatedW.selected[models.SWC],
		Typescript: updatedW.selected[models.TYPESCRIPT],
	}

	return userAnswers
}

func (w WizardAnswers) Init() tea.Cmd {
	return nil
}

func (w WizardAnswers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// not working
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.height = msg.Height
		w.width = msg.Width
	}

	if w.appName == "" {
		if w.appNameTextInputRendered == false {
			w.appNameTextInput = textinput.New()
			w.appNameTextInput.Placeholder = "Type your App Name"
			w.appNameTextInput.Focus()
			w.appNameTextInputRendered = true
		}
		return w.updateTextInput(msg)
	} else {
		return w.updateSelector(msg)
	}
}

func (w WizardAnswers) View() string {
	if w.width == 0 {
		return "Loading..."
	}

	if w.appName == "" {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			w.appNameTextInput.Value(),
			w.styles.InputField.Render(w.appNameTextInput.View()),
		)
	} else {
		// header
		s := "\nSelect all the configurations that you would like to generate:\n\n"

		// have to do it this way because maps don't guarantee insertion order
		for i := 0; i < len(w.selected); i++ {
			cursor := " "
			if w.cursor == i {
				cursor = ">"
			}

			selected := " "
			if v, ok := w.selected[i]; ok && v {
				selected = "X"
			}

			keyName := mapSelectedKeyToName(i)
			s += renderRow(cursor, selected, keyName)
		}

		return s
	}
}

func mapSelectedKeyToName(key int) string {
	switch key {
	case models.ES_BUILD:
		return "ESBuild"
	case models.ES_LINT:
		return "ESLint"
	case models.JEST:
		return "Jest"
	case models.PRETTIER:
		return "Prettier"
	case models.SWC:
		return "SWC"
	case models.TYPESCRIPT:
		return "TypeScript"
	}
	return ""
}

func renderRow(cursor string, checked string, choice string) string {
	return fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
}

func (w WizardAnswers) updateSelector(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			if int(w.cursor) < len(w.selected)-1 {
				w.cursor++
			}
		case "enter", " ":
			b, ok := w.selected[w.cursor]
			if ok {
				w.selected[w.cursor] = !b
			} else {
				log.Fatalf("Couldn't access selected map. Accessed with key: (%d)", w.cursor)
			}
		}
	}

	return w, nil
}

func (w WizardAnswers) updateTextInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	w.appNameTextInput, cmd = w.appNameTextInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "enter":
			w.appName = w.appNameTextInput.Value()
			w.appNameTextInput.Blur()
			return w, cmd
		}
	}
	return w, cmd
}
