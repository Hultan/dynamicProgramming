package main

import (
	"strings"
)

func canConstructWithMemo(target string, words []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return true
	}
	for _, word := range words {
		if strings.Index(target,word)  == 0 {
			if canConstructWithMemo(target[len(word):], words, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

func canConstruct(target string, words []string) bool {
	if target == "" {
		return true
	}
	for _, word := range words {
		if strings.Index(target,word)  == 0 {
			if canConstruct(target[len(word):], words) {
				return true
			}
		}
	}
	return false
}
