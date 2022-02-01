package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type Word struct {
	word           string
	usageFrequency int
}

func Top10(text string) []string {
	const countOfTopWords = 10

	if len(text) == 0 {
		return nil
	}

	words := regexp.MustCompile(`\s`).Split(text, -1)
	var (
		wordsCounter = make(map[string]int)
		result       = make([]string, 0)
	)

	for _, word := range words {
		if word == "" {
			continue
		}
		wordsCounter[word]++
	}

	for i, word := range rankWords(wordsCounter) {
		if i > countOfTopWords-1 {
			break
		}
		result = append(result, word.word)
	}
	return result
}

func rankWords(usedWords map[string]int) []Word {
	rankedWords := make([]Word, 0)
	for usedWord, usageFrequency := range usedWords {
		rankedWords = append(rankedWords, Word{usedWord, usageFrequency})
	}

	sort.Slice(rankedWords, func(i, j int) bool {
		if rankedWords[i].usageFrequency == rankedWords[j].usageFrequency {
			return rankedWords[i].word < rankedWords[j].word
		}
		return rankedWords[i].usageFrequency > rankedWords[j].usageFrequency
	})

	return rankedWords
}
