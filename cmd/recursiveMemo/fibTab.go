package main

func fibTab(n int) uint {
	// Create table for n+3 integers
	// n + 1 values for 0..n, + 2 additional values
	// for	table[i+1] and table[i+2] in the loop
	table := make([]uint, n+3)
	table[1] = 1

	for i := 0; i <= n; i++ {
		table[i+1] += table[i]
		table[i+2] += table[i]
	}

	return table[n]
}
