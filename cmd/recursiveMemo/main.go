package main

// https://www.youtube.com/watch?v=oBt53YbR9Kk&list=WL&index=13

import (
	"fmt"
)

func main() {
	// fmt.Println(fib(50, map[int]int{}))
	// fmt.Println(travel(18,18, map[string]int{}))
	// fmt.Println(canSum(7, []int{5, 3, 4, 7}, map[int]bool{}))
	// fmt.Println(canSum(7, []int{2,4}, map[int]bool{}))
	// fmt.Println(canSum(300, []int{7,21}, map[int]bool{}))
	// fmt.Println(howSum(300, []int{7,14}, map[int][]int{}))
	fmt.Println(bestSum(100, []int{1,2,5,25}, map[int][]int{}))
}

func bestSum(sum int, values []int, memo map[int][]int) []int {
	if list, ok := memo[sum];ok {
		return list
	}
	if sum < 0 {
		return nil
	}
	if sum == 0 {
		return []int{}
	}
	var best []int
	for _, value := range values {
		list := bestSum(sum-value, values, memo)
		if list != nil {
			list = append(list, value)
			memo[sum] = list
			if best == nil || len(list) < len(best) {
				best = list
			}
		}
	}
	if best!=nil {
		return best
	}
	memo[sum] = nil
	return nil
}

func howSum(sum int, values []int, memo map[int][]int) []int {
	if list, ok := memo[sum];ok {
		return list
	}
	if sum < 0 {
		return nil
	}
	if sum == 0 {
		return []int{}
	}
	for _, value := range values {
		list := howSum(sum-value, values, memo)
		memo[sum] = list
		if list != nil {
			list = append(list, value)
			return list
		}
	}
	memo[sum] = nil
	return nil
}

func canSum(sum int, values []int, memo map[int]bool) bool {
	if value, ok := memo[sum]; ok {
		return value
	}
	if sum < 0 {
		return false
	}
	if sum == 0 {
		return true
	}
	for _, value := range values {
		if canSum(sum-value, values, memo) {
			memo[sum] = true
			return true
		}
	}
	memo[sum] = false
	return false
}

func travel(m, n int, memo map[string]int) int {
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
	memo[key] = travel(m-1, n, memo) + travel(m, n-1, memo)
	return memo[key]
}

func fib(n int, memo map[int]int) int {
	if value, ok := memo[n]; ok {
		return value
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
