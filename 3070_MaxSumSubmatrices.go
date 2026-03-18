func countSubmatrices(grid [][]int, k int) int {
    if grid[0][0] > k {
        return 0
    }
    var fout int

    n := len(grid)
    m := len(grid[0])
    
    cumsums := make([][]int,n)
    for i := 0; i < n; i++ {
        cumsums[i] = make([]int,m)
        for j := 0; j < m; j++ {
            if i > 0 {
                if j == 0 {
                    cumsums[i][0] = cumsums[i-1][0] + grid[i][0]
                } else {
                    cumsums[i][j] = grid[i][j] + cumsums[i-1][j] + cumsums[i][j-1] - cumsums[i-1][j-1]
                }
            } else { //top row
                if j == 0 {
                    cumsums[0][0] = grid[0][0]
                } else {
                    cumsums[0][j] = grid[0][j] + cumsums[0][j-1]
                }
            }
            if cumsums[i][j] > k {
                m = j
                break
            } else {
                fout++
            }
        }
    }
    return fout
}
