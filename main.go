package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typ0",
	Short: "A CLI typing practice tool",
	Long:  `A interactive CLI tool for typing practice.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🏁 Welcome to Typ0!")
		fmt.Println("Start typing: typ0 race")
		fmt.Println("Show help: typ0 --help")
	},
}

func init() {
	rootCmd.AddCommand(raceCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}