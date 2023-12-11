package tui

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderTopForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

	return s
}
