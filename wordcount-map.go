package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	wordArray := strings.Fields(s)
	wordMap := make(map[string]int)
	
	for _, value := range wordArray {
		if count, exists := wordMap[value]; exists == true {
			wordMap[value] = count + 1
		} else {
			wordMap[value] = 1
		}
	}
	
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
