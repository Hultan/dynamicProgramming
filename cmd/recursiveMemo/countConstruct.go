package main

import (
	"strings"
)

func countConstructWithMemo(target string, words []string, memo map[string]int) int {
	if value, ok := memo[target];ok {
		return value
	}
	if target == "" {
		return 1
	}
	count := 0
	for _, word := range words {
		if strings.Index(target,word)  == 0 {
			count += countConstructWithMemo(target[len(word):], words, memo)
		}
	}
	memo[target] = count
	return count
}

func countConstruct(target string, words []string) int {
	if target == "" {
		return 1
	}
	count := 0
	for _, word := range words {
		if strings.Index(target,word)  == 0 {
			count += countConstruct(target[len(word):], words)
		}
	}
	return count
}
