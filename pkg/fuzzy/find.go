package fuzzycustom

import (
	"math"
)

// improve with algos from https://tilores.io/fuzzy-wuzzy-online-tool

var (
	proximityMultiplier = 1000
	matchMultiplyer     = 100
	sequenceMultiplyer  = 30
)

// Find searches for pattern matches in a slice of strings and returns sorted results
func Find(pattern string, data []string) *Results {
	results := &Results{
		Ranked: make([]Result, 0, len(data)),
	}

	for _, word := range data {
		finder := NewFinder(pattern, word)
		score := finder.Match()

		if score > 1 {
			result := Result{
				Score:     score,
				Word:      word,
				Positions: finder.bestPositions,
			}
			results.Ranked = append(results.Ranked, result)
		}
	}

	results.Sort()

	return results
}

// Finder contains all the state and methods for the fuzzy matching algorithm
type Finder struct {
	matches       map[rune][]int
	wordStr       string
	wordRune      []rune
	patternStr    string
	patternRune   []rune
	bestScore     int
	bestPositions []int
}

// NewFinder creates a new Finder instance with initialized fields
func NewFinder(pattern string, word string) *Finder {
	return &Finder{
		patternStr:    pattern,
		patternRune:   []rune(pattern),
		wordStr:       word,
		wordRune:      []rune(word),
		matches:       make(map[rune][]int),
		bestScore:     -1,
		bestPositions: []int{},
	}
}

// Match compares the pattern with the word and returns a score
func (f *Finder) Match() int {
	// Find all possible matches between pattern and word
	f.findMatches()
	f.calculateScore()
	return f.bestScore
}

// findMatches identifies all occurrences of pattern characters in the word
// TODO: is this helping at all?
func (f *Finder) findMatches() {
	for i, char := range f.wordRune {
		for _, patternChar := range f.patternRune {
			if char == patternChar {
				f.matches[char] = append(f.matches[char], i)
			}
		}
	}
}

func (f *Finder) calculateScore() int {
	if len(f.matches) == 0 {
		return -1
	}

	score := 1
	longestSeq, startingPos := f.findLongestOrderedSequence()
	if longestSeq == len(f.patternStr) {
		f.bestPositions = PosBuilder(startingPos, longestSeq)
		score *= matchMultiplyer
	} else {
		score *= sequenceMultiplyer
	}

	f.bestPositions = f.calculateProximity()
	score *= PositionsToScore(f.bestPositions)

	// NOTE: weighting the matches at the back is for gotex. Needs to be an optional thing to reuse pkg. Score functions as options?
	if len(f.bestPositions) > 0 {
		score += f.bestPositions[0]
	}

	// TODO: algos could be used to get different types of scores
	// - minimum edit distance might be good for typos? but could sqew numbers in larger strings | parse the shortest in order occurrences?

	f.bestScore = score
	return score
}

// findLongestOrderedSequence is just for scoring
// TODO: could this just be LCS algo?
func (f *Finder) findLongestOrderedSequence() (int, int) {
	maxLen := -1
	startingInd := -1
	for i := range len(f.wordStr) {
		patternIndex := 0
		j := i
		length := 0

		for j < len(f.wordStr) && patternIndex < len(f.patternStr) {
			if f.wordStr[j] == f.patternStr[patternIndex] {
				length++
				patternIndex++
				j++
			} else {
				break
			}
		}

		if length > 0 && length > maxLen {
			maxLen = length
			startingInd = i
		}
	}

	return maxLen, startingInd
}

func (f *Finder) calculateProximity() []int {
	charPositions := make(map[rune][]int)
	for _, char := range f.patternStr {
		positions := []int{}
		for i, c := range f.wordStr {
			if c == char {
				positions = append(positions, i)
			}
		}
		if len(positions) == 0 {
			return []int{}
		}
		charPositions[char] = positions
	}

	substringRunes := []rune(f.patternStr)
	minLength := math.MaxInt32
	var bestIndices []int

	// Recursive function to find the shortest valid sequence
	var backtrack func(charIdx, startPos int, currentIndices []int)
	backtrack = func(charIdx, startPos int, currentIndices []int) {
		if charIdx == len(substringRunes) {
			if len(currentIndices) > 0 {
				spanLength := currentIndices[len(currentIndices)-1] - currentIndices[0] + 1
				if spanLength < minLength {
					minLength = spanLength
					bestIndices = make([]int, len(currentIndices))
					copy(bestIndices, currentIndices)
				}
			}
			return
		}

		char := substringRunes[charIdx]
		for _, pos := range charPositions[char] {
			if pos >= startPos { // Ensure we maintain the order
				currentIndices = append(currentIndices, pos)
				backtrack(charIdx+1, pos+1, currentIndices)
				currentIndices = currentIndices[:len(currentIndices)-1] // Remove the last element
			}
		}
	}

	backtrack(0, 0, []int{})
	return bestIndices
}
