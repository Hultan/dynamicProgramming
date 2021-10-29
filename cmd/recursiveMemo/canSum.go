package main

func canSumWithMemo(sum int, values []int, memo map[int]bool) bool {
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
		if canSumWithMemo(sum-value, values, memo) {
			memo[sum] = true
			return true
		}
	}
	memo[sum] = false
	return false
}

func canSum(sum int, values []int) bool {
	if sum < 0 {
		return false
	}
	if sum == 0 {
		return true
	}
	for _, value := range values {
		if canSum(sum-value, values) {
			return true
		}
	}
	return false
}
