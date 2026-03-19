func numberOfSubmatrices(grid [][]byte) int {
	var fout int
	n := len(grid)
	m := len(grid[0])
	xcount := make([][]int, n)
	ycount := make([][]int, n)
	for i := 0; i < n; i++ {
		xcount[i] = make([]int, m)
		ycount[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i > 0 {
				if j > 0 {
					xcount[i][j] = xcount[i-1][j] + xcount[i][j-1] - xcount[i-1][j-1]
					ycount[i][j] = ycount[i-1][j] + ycount[i][j-1] - ycount[i-1][j-1]
				} else {
					xcount[i][0] = xcount[i-1][0]
					ycount[i][0] = ycount[i-1][0]
				}
			} else { //top row
				if j > 0 {
					xcount[0][j] = xcount[0][j-1]
					ycount[0][j] = ycount[0][j-1]
				}
			}
			if grid[i][j] == 'X' {
				xcount[i][j]++
			} else if grid[i][j] == 'Y' {
				ycount[i][j]++
			}
			if xcount[i][j] > 0 {
				if xcount[i][j] == ycount[i][j] {
					fout++
				}
			}
		}
	}
	return fout
}
