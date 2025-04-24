package fuzzy

import (
	"fmt"
	"sort"
	"strings"
)

// Result represents a single match result with score and position data
type Result struct {
	Score     int
	Word      string
	Positions []int
	Threshold int // TODO: this will be a Threshold to limit out low scoring matches
	// TODO: Some variance of threshold where if it doesn't have 90% of characters don't store it (allow for typos)
}

// Print outputs the word with matched positions highlighted in red
func (r *Result) Print() {
	ansiWord := HighlightANSI(r.Word, r.Positions)

	fmt.Printf("Word: %s, Score: %d\n", ansiWord, r.Score)
}

func HighlightANSI(word string, positions []int) string {
	wordRunes := []rune(word)

	posMap := make(map[int]bool)
	for _, pos := range positions {
		posMap[pos] = true
	}

	var sb strings.Builder

	for i, char := range wordRunes {
		if posMap[i] {
			// ANSI escape code for red text
			sb.WriteString("\033[31m")
			sb.WriteRune(char)
			sb.WriteString("\033[0m") // Reset color
		} else {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

// Results holds multiple ranked search results
type Results struct {
	Ranked []Result
}

func (r *Results) ToStringSlice() []string {
	str := make([]string, 0)
	for _, result := range r.Ranked {
		str = append(str, result.Word)
	}
	return str
}

func (r *Results) ToStringSliceANSI() []string {
	str := make([]string, 0)
	for _, result := range r.Ranked {
		str = append(str, HighlightANSI(result.Word, result.Positions))
	}
	return str
}

func (r *Results) Sort() {
	sort.Slice(r.Ranked, func(i, j int) bool {
		return r.Ranked[i].Score > r.Ranked[j].Score
	})
}

// Print displays all results with highlighting
func (r *Results) Print() {
	fmt.Printf("Found %d results:\n", len(r.Ranked))
	for i, result := range r.Ranked {
		fmt.Printf("%d. ", i+1)
		result.Print()
	}
}

func PositionsToScore(pos []int) int {
	if len(pos) == 0 {
		return 0
	}

	diff := pos[len(pos)-1] - pos[0]

	if diff == 0 {
		return 0
	}

	return proximityMultiplier / diff
}

func PosBuilder(start, size int) []int {
	pos := make([]int, size)
	for i := range size {
		pos[i] = start + i - 1
	}
	return pos
}
