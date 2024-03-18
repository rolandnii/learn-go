package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordMap := make(map[string]int)

	for _, v := range words {
		sum := 0

		for _, word := range words {
			if v == word {
				sum++
			}
		}
		wordMap[v] = sum
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
