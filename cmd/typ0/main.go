package main

import (
	"fmt"
	"os"

	"go-typ0/internal/race"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typ0",
	Short: "A CLI typing practice tool",
	Long:  `An interactive CLI tool for typing practice with real-time feedback and statistics.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🏁 Welcome to Typ0!")
		fmt.Println("Start typing: typ0 race")
		fmt.Println("Show help: typ0 --help")
	},
}

func init() {
	rootCmd.AddCommand(race.NewCommand())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 