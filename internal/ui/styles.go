package ui

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BoxStyle           lipgloss.Style
	GreenStyle         lipgloss.Style
	RedStyle           lipgloss.Style
	UnderlineStyle     lipgloss.Style
	StatsBoxStyle      lipgloss.Style
	LabelStyle         lipgloss.Style
	ValueStyle         lipgloss.Style
	MistypedKeyStyle   lipgloss.Style
}

func NewStyles() *Styles {
	return &Styles{
		BoxStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(1, 2),
		
		GreenStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")),
		
		RedStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")),
		
		UnderlineStyle: lipgloss.NewStyle().
			Underline(true),
		
		StatsBoxStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			MarginTop(1),
		
		LabelStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Bold(true),
		
		ValueStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("6")).
			Bold(true),
		
		MistypedKeyStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")).
			Bold(true),
	}
} 