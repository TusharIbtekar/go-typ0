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