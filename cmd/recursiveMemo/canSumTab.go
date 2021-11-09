package main

func canSumTab(sum int, values []int) bool {
	memo := make([]bool, sum+1, sum+1)
	memo[0] = true
	for i := 0; i <= sum; i++ {
		if !memo[i] {
			continue
		}

		for _, value := range values {
			if i+value == sum {
				return true
			} else if i+value < sum {
				memo[i+value] = true
			}
		}
	}
	return false
}
