package main

import (
	"fmt"
	"strings"
)

type TextStats struct {
	Char   int
	Word   int
	Sentence int
}

func textStats(text string) TextStats {
	char := len([]rune(text))

	words := strings.Fields(text)
	word := len(words)

	sentence := 0
	for _, char := range text {
		if char == '.' || char == '!' || char == '?' {
			sentence++
		}
	}

	return TextStats{
		Char:     char,
		Word:     word,
		Sentence: sentence,
	}
}

func main() {
	Text := "Привет! Как твои дела? Что нового?"

	stats := textStats(Text)

	fmt.Printf("Количество символов: %d\n", stats.Char)
	fmt.Printf("Количество слов: %d\n", stats.Word)
	fmt.Printf("Количество предложений: %d\n", stats.Sentence)
}