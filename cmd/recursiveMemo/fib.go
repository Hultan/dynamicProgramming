package main

func fibWithMemo(n int, memo map[int]int) int {
	if value, ok := memo[n]; ok {
		return value
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fibWithMemo(n-1, memo) + fibWithMemo(n-2, memo)
	return memo[n]
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
