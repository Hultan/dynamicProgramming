package main

import (
	"fmt"
)

func travelWithMemo(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	if value, ok := memo[key]; ok {
		return value
	}
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	memo[key] = travelWithMemo(m-1, n, memo) + travelWithMemo(m, n-1, memo)
	return memo[key]
}

func travel(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	return travel(m-1, n) + travel(m, n-1)
}

