package main

func howSumWithMemo(sum int, values []int, memo map[int][]int) []int {
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
		list := howSumWithMemo(sum-value, values, memo)
		memo[sum] = list
		if list != nil {
			list = append(list, value)
			return list
		}
	}
	memo[sum] = nil
	return nil
}

func howSum(sum int, values []int) []int {
	if sum < 0 {
		return nil
	}
	if sum == 0 {
		return []int{}
	}
	for _, value := range values {
		list := howSum(sum-value, values)
		if list != nil {
			list = append(list, value)
			return list
		}
	}
	return nil
}

