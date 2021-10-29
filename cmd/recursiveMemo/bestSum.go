package main

func bestSumWithMemo(sum int, values []int, memo map[int][]int) []int {
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
		list := bestSumWithMemo(sum-value, values, memo)
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

func bestSum(sum int, values []int) []int {
	if sum < 0 {
		return nil
	}
	if sum == 0 {
		return []int{}
	}
	var best []int
	for _, value := range values {
		list := bestSum(sum-value, values)
		if list != nil {
			list = append(list, value)
			if best == nil || len(list) < len(best) {
				best = list
			}
		}
	}
	if best!=nil {
		return best
	}
	return nil
}

