package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var practiceCmd = &cobra.Command{
	Use: "practice",


	Run: func(cmd *cobra.Command, args []string) {
		sentence := "The quick brown fox jumps over the lazy dog"
		fmt.Println("Type the following")
		fmt.Println(sentence)
		fmt.Println("Press Enter to start")

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		fmt.Println("Start typing! Press Enter when done")

		startTime := time.Now()
		input, _ := reader.ReadString('\n')
		if input == "" {
			fmt.Println("You didn't type anything!")
			return
		}
		
		duration := time.Since(startTime)


		input = strings.TrimSpace(input)
		accuracy := calculateAccuracy(sentence, input)
		wpm := calculateWPM(len(sentence), duration)

		fmt.Printf("Results:\n")
		fmt.Printf("Time: %.2f seconds\n", duration.Seconds())
		fmt.Printf("Accuracy: %.2f%%\n", accuracy)
		fmt.Printf("WPM: %.2f\n", wpm)

	},
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

	fmt.Printf("DEBUG: charCount=%d, words=%.2f, minutes=%.4f\n", charCount, words, minutes)

	if minutes == 0 {
		return 0
	}
	
	return words / minutes
}