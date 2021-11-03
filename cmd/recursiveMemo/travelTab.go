package main

func travelTab(m, n int) int {
	// Create a table (slice) that is two larger than m and m so that we can :
	// * fit the zero column
	// * fit the m+1 and n+1 column, so that we can avoid having to do index-out-of-bounds checks
	table := make([][]int, m+2)
	for i := range table {
		table[i] = make([]int, n+2)
	}
	table[1][1] = 1

	for x := 0; x <= m; x++ {
		for y := 0; y <= n; y++ {
			table[x+1][y] += table[x][y]
			table[x][y+1] += table[x][y]
		}
	}

	return table[m][n]
}
