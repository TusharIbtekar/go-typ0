package race

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var wordCount int

	cmd := &cobra.Command{
		Use:     "race",
		Aliases: []string{"r", "type", "practice"},
		Short:   "Start a typing race",
		Long:    `Start a typing race with random sentences. Race against time to improve your typing speed!`,
		Run: func(cmd *cobra.Command, args []string) {
			model := NewModel(wordCount)
			viewModel := NewViewModel(model)

			p := tea.NewProgram(viewModel)
			if _, err := p.Run(); err != nil {
				fmt.Println("Error running program: ", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().IntVarP(&wordCount, "words", "w", 20, "Number of words in the sentence")

	return cmd
} 