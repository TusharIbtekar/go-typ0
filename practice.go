package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/spf13/cobra"
)


var practiceCmd = &cobra.Command{
	Use: "practice",

	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(&model{})
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program: ", err)
			os.Exit(1)
		}
	},
}

var sentence = "The quick brown fox jumps over the lazy dog" 
type model struct {
	input string
	startTime time.Time
	finished bool
	mistyped map[rune]int
}

func (m *model) Init() tea.Cmd {
	m.startTime = time.Now()
	m.mistyped = make(map[rune]int)
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.finished {
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.Type {
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
				case tea.KeyEnter: 
					m.input = ""
					m.startTime = time.Now()
					m.finished = false
					m.mistyped = make(map[rune]int)
					return m, nil
			}

		}
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg: 
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc: 
			return m, tea.Quit
		case tea.KeyEnter:
			m.finished = true
			return m, nil
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(msg.String()) == 1 && len(m.input) < len(sentence) {
				typed := msg.String()
				expected := string(sentence[len(m.input)])
				if typed != expected {
					m.mistyped[rune(expected[0])]++
				}
				m.input += typed
				
				if len(m.input) == len(sentence) {
					m.finished = true
				}
			}
	}
	}
	return m, nil
}

func (m *model) View() string {
	boxStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1, 2)
	green := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	red := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	underline := lipgloss.NewStyle().Underline(true)

	statsBoxStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).MarginTop(1)
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Bold(true)
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true)
	mistypedKeyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)

	var sentenceView string
	for i := 0; i < len(sentence); i++ {
		if i < len(m.input) {
			if m.input[i] == sentence[i] {
				sentenceView += green.Render(string(sentence[i]))
			} else {
				sentenceView += red.Render(string(sentence[i]))
			}
		} else if i == len(m.input) && !m.finished {
			sentenceView += underline.Render(string(sentence[i]))
		} else {
			sentenceView += string(sentence[i])
		}
	}
	contentWidth := lipgloss.Width(sentence) + 5
	sentenceBox := boxStyle.Width(contentWidth).Render(sentenceView)

	cursor := " "
	if !m.finished && time.Now().UnixNano()/500000000%2 == 0 {
		cursor = "_"
	}
	inputContent := m.input + cursor
	inputBox := boxStyle.Width(contentWidth).Render(inputContent)

	stats := ""
	if m.finished {
		duration := time.Since(m.startTime)
		accuracy := calculateAccuracy(sentence, m.input)
		wpm := calculateWPM(len(m.input), duration)
		// stats = fmt.Sprintf(
		// 	"\nTime: %.2f seconds | WPM: %.2f | Accuracy: %.2f%%\n", 
		// 	duration.Seconds(), wpm, accuracy,
		// )

		statsLines := []string {
			fmt.Sprintf("%s %s", labelStyle.Render("Time:"), valueStyle.Render(fmt.Sprintf("%.2f seconds", duration.Seconds()))),
			fmt.Sprintf("%s %s", labelStyle.Render("WPM:"), valueStyle.Render(fmt.Sprintf("%.2f", wpm))),
			fmt.Sprintf("%s %s", labelStyle.Render("Accuracy:"), valueStyle.Render(fmt.Sprintf("%.2f%%", accuracy))),
		}

		// mistypes section
		if len(m.mistyped) > 0 {
			type kv struct {k rune; v int}
			var sorted []kv
			for k, v := range m.mistyped {
				sorted = append(sorted, kv{k, v})
			}
			sort.Slice(sorted, func(i, j int) bool { return sorted[i].v > sorted[j].v})
			mistypedStr := ""
			for i, pair := range sorted {
				if i >= 5 {
					break
				}
				mistypedStr += fmt.Sprintf("- %s %s\n", mistypedKeyStyle.Render(fmt.Sprintf("%q", pair.k)), valueStyle.Render(fmt.Sprintf("%d", pair.v)))
			}
			statsLines = append(statsLines, labelStyle.Render("Mistypes: "))
			statsLines = append(statsLines, mistypedStr)
		}
		statsLines = append(statsLines, labelStyle.Render("Press Enter to restart. ESC/CTRL+C to quit"))
		stats = statsBoxStyle.Render(strings.Join(statsLines, "\n"))
	} else {
		stats = "\nPress Enter when done. ESC/CTRL+C to quit"
	}


	return sentenceBox + "\n\n" + inputBox + "\n" + stats
}


func calculateAccuracy(original, input string) float64 {
	original = strings.TrimSpace(original)
	input = strings.TrimSpace(input)

	if len(original) == 0 || len(input) == 0 {
		return 0
	}

	correct := 0
	minLen := min(len(original), len(input))
	for i := 0; i < minLen; i++ {
		if original[i] == input[i] {
			correct++
		}
	}
	return float64(correct) / float64(minLen) * 100
}

func calculateWPM(charCount int, duration time.Duration) float64 {
	words := float64(charCount) / 5.0
	minutes := duration.Minutes()


	if minutes == 0 {
		return 0
	}
	
	return words / minutes
}