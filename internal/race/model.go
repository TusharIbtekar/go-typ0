package race

import (
	"math/rand"
	"sort"
	"strings"
	"time"

	"go-typ0/internal/words"
)

type Model struct {
	input     string
	startTime time.Time
	finished  bool
	mistyped  map[rune]int
	sentence  string
	width     int
	height    int
	wordCount int
}

func NewModel(wordCount int) *Model {
	return &Model{
		wordCount: wordCount,
		mistyped:  make(map[rune]int),
	}
}

func (m *Model) Init() {
	m.startTime = time.Now()
	m.mistyped = make(map[rune]int)
	m.sentence = m.generateRandomSentence()
	m.finished = false
	m.input = ""
}

func (m *Model) GetStats() Stats {
	if !m.finished {
		return Stats{}
	}

	duration := time.Since(m.startTime)
	accuracy := m.calculateAccuracy()
	wpm := m.calculateWPM(duration)

	return Stats{
		Duration:  duration,
		Accuracy:  accuracy,
		WPM:       wpm,
		Mistyped:  m.getTopMistyped(5),
		Finished:  true,
	}
}

func (m *Model) HandleInput(input string) {
	if m.finished {
		return
	}

	if len(input) == 1 && len(m.input) < len(m.sentence) {
		expected := string(m.sentence[len(m.input)])
		
		if expected == "\n" {
			m.input += expected
		} else if input != expected {
			m.mistyped[rune(expected[0])]++
			m.input += input
		} else {
			m.input += input
		}
		
		if len(m.input) == len(m.sentence) {
			m.finished = true
		}
	}
}

func (m *Model) HandleBackspace() {
	if len(m.input) > 0 {
		m.input = m.input[:len(m.input)-1]
	}
}

func (m *Model) Restart() {
	m.Init()
}

type Stats struct {
	Duration  time.Duration
	Accuracy  float64
	WPM       float64
	Mistyped  []MistypedChar
	Finished  bool
}

type MistypedChar struct {
	Char  rune
	Count int
}

func (m *Model) generateRandomSentence() string {
	if m.wordCount <= 0 {
		m.wordCount = 20
	}
	
	var sentence []string
	for i := 0; i < m.wordCount; i++ {
		randomIndex := rand.Intn(len(words.Words))
		sentence = append(sentence, words.Words[randomIndex])
	}
	
	fullSentence := strings.Join(sentence, " ")
	return m.wrapText(fullSentence, 80)
}

func (m *Model) wrapText(text string, maxWidth int) string {
	words := strings.Fields(text)
	var lines []string
	currentLine := ""
	
	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxWidth {
			if currentLine != "" {
				currentLine += " " + word
			} else {
				currentLine = word
			}
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	
	return strings.Join(lines, "\n")
}

func (m *Model) calculateAccuracy() float64 {
	original := strings.TrimSpace(m.sentence)
	input := strings.TrimSpace(m.input)

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

func (m *Model) calculateWPM(duration time.Duration) float64 {
	charCount := len(m.input)
	words := float64(charCount) / 5.0
	minutes := duration.Minutes()

	if minutes == 0 {
		return 0
	}
	
	return words / minutes
}

func (m *Model) getTopMistyped(n int) []MistypedChar {
	if len(m.mistyped) == 0 {
		return nil
	}

	type kv struct {
		k rune
		v int
	}
	var sorted []kv
	for k, v := range m.mistyped {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].v > sorted[j].v })
	
	var result []MistypedChar
	for i, pair := range sorted {
		if i >= n {
			break
		}
		result = append(result, MistypedChar{
			Char:  pair.k,
			Count: pair.v,
		})
	}
	
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} 