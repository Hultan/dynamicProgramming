package main

import (
	"strings"
)

func allConstructWithMemo(target string, words []string, memo map[string][][]string) [][]string {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range words {
		if strings.Index(target,word) == 0 {
			suffixWays := allConstructWithMemo(target[len(word):], words, memo)
			targetWays := Map(suffixWays, func(way []string) []string { return append([]string{word}, way...) })
			result = append(result,targetWays...)
		}
	}
	memo[target] = result
	return result
}

func allConstruct(target string, words []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range words {
		if strings.Index(target,word) == 0 {
			suffixWays := allConstruct(target[len(word):], words)
			targetWays := Map(suffixWays, func(way []string) []string { return append([]string{word}, way...) })
			result = append(result,targetWays...)
		}
	}
	return result
}

// Map : Since Go does not have a builtin Map function, we need to
// write one ourselves. See https://gobyexample.com/collection-functions
func Map(vs [][]string, f func([]string) []string) [][]string {
	vsm := make([][]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// func allConstructWithMemo(target string, words []string, memo map[string]int) int {
// 	if value, ok := memo[target];ok {
// 		return value
// 	}
// 	if target == "" {
// 		return 1
// 	}
// 	count := 0
// 	for _, word := range words {
// 		if strings.Index(target,word)  == 0 {
// 			count += allConstructWithMemo(target[len(word):], words, memo)
// 		}
// 	}
// 	memo[target] = count
// 	return count
// }
