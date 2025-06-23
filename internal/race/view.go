package race

import (
	"fmt"
	"strings"
	"time"

	"go-typ0/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ViewModel struct {
	model  *Model
	styles *ui.Styles
}

func NewViewModel(model *Model) *ViewModel {
	return &ViewModel{
		model:  model,
		styles: ui.NewStyles(),
	}
}

func (vm *ViewModel) Init() tea.Cmd {
	vm.model.Init()
	return tea.EnterAltScreen
}

func (vm *ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		vm.model.width = msg.Width
		vm.model.height = msg.Height
		return vm, nil
	}

	if vm.model.finished {
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.Type {
			case tea.KeyCtrlC, tea.KeyEsc:
				return vm, tea.Quit
			case tea.KeyRunes:
				if len(key.Runes) == 1 && key.Runes[0] == 'q' {
					return vm, tea.Quit
				}
			case tea.KeyEnter:
				vm.model.Restart()
				return vm, nil
			}
		}
		return vm, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return vm, tea.Quit
		case tea.KeyEnter:
			vm.model.finished = true
			return vm, nil
		case tea.KeyBackspace:
			vm.model.HandleBackspace()
		default:
			vm.model.HandleInput(msg.String())
		}
	}

	return vm, nil
}

func (vm *ViewModel) View() string {
	sentenceView := vm.renderSentence(vm.model.sentence, vm.model.input)
	contentWidth := lipgloss.Width(vm.model.sentence) + 5
	sentenceBox := vm.styles.BoxStyle.Width(contentWidth).Render(sentenceView)

	cursor := " "
	if !vm.model.finished && time.Now().UnixNano()/500000000%2 == 0 {
		cursor = "_"
	}
	inputContent := vm.model.input + cursor
	inputBox := vm.styles.BoxStyle.Width(contentWidth).Render(inputContent)

	stats := vm.renderStats()

	content := sentenceBox + "\n\n" + inputBox + "\n" + stats

	if vm.model.width > 0 && vm.model.height > 0 {
		centered := lipgloss.Place(
			vm.model.width, vm.model.height,
			lipgloss.Center, lipgloss.Center,
			content,
		)
		return centered
	}

	return content
}

func (vm *ViewModel) renderSentence(sentence, input string) string {
	var sentenceView string
	for i := 0; i < len(sentence); i++ {
		if i < len(input) {
			if input[i] == sentence[i] {
				sentenceView += vm.styles.GreenStyle.Render(string(sentence[i]))
			} else {
				sentenceView += vm.styles.RedStyle.Render(string(sentence[i]))
			}
		} else if i == len(input) && !vm.model.finished {
			sentenceView += vm.styles.UnderlineStyle.Render(string(sentence[i]))
		} else {
			sentenceView += string(sentence[i])
		}
	}
	return sentenceView
}

func (vm *ViewModel) renderStats() string {
	if vm.model.finished {
		stats := vm.model.GetStats()
		return vm.renderFinishedStats(stats)
	}
	return "\nPress Enter when done. ESC/CTRL+C to quit"
}

func (vm *ViewModel) renderFinishedStats(stats Stats) string {
	statsLines := []string{
		fmt.Sprintf("%s %s", vm.styles.LabelStyle.Render("Time:"), vm.styles.ValueStyle.Render(fmt.Sprintf("%.2f seconds", stats.Duration.Seconds()))),
		fmt.Sprintf("%s %s", vm.styles.LabelStyle.Render("WPM:"), vm.styles.ValueStyle.Render(fmt.Sprintf("%.2f", stats.WPM))),
		fmt.Sprintf("%s %s", vm.styles.LabelStyle.Render("Accuracy:"), vm.styles.ValueStyle.Render(fmt.Sprintf("%.2f%%", stats.Accuracy))),
	}

	if len(stats.Mistyped) > 0 {
		mistypedStr := ""
		for _, mistyped := range stats.Mistyped {
			mistypedStr += fmt.Sprintf("- %s %s\n", 
				vm.styles.MistypedKeyStyle.Render(fmt.Sprintf("%q", mistyped.Char)), 
				vm.styles.ValueStyle.Render(fmt.Sprintf("%d", mistyped.Count)))
		}
		statsLines = append(statsLines, vm.styles.LabelStyle.Render("Mistypes: "))
		statsLines = append(statsLines, mistypedStr)
	}

	statsLines = append(statsLines, vm.styles.LabelStyle.Render("Press Enter to restart. ESC/CTRL+C/Q to quit"))
	return vm.styles.StatsBoxStyle.Render(strings.Join(statsLines, "\n"))
} 