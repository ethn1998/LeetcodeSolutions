func countIslands(grid [][]int, k int) int {
    var fout int
    m := len(grid)
    n := len(grid[0])

    for k0 := 0; k0 < m*n; k0++ {
        r0 := k0 / n
        c0 := k0 % n
        if grid[r0][c0] > 0 {
            val := grid[r0][c0]
            grid[r0][c0] = 0
            q := []int{k0}
            for i := 0; i < len(q); i++ {
                j := q[i]
                r := j / n
                c := j % n
                if r > 0 {
                    if grid[r-1][c] > 0 {
                        val += grid[r-1][c]
                        grid[r-1][c] = 0
                        q = append(q,j-n)
                    }
                }
                if r < m-1 {
                    if grid[r+1][c] > 0 {
                        val += grid[r+1][c]
                        grid[r+1][c] = 0
                        q = append(q,j+n)
                    }
                }
                if c > 0 {
                    if grid[r][c-1] > 0 {
                        val += grid[r][c-1]
                        grid[r][c-1] = 0
                        q = append(q,j-1)
                    }
                }
                if c < n-1 {
                    if grid[r][c+1] > 0 {
                        val += grid[r][c+1]
                        grid[r][c+1] = 0
                        q = append(q,j+1)
                    }
                }
                //fmt.Println(q)
            }
            if val % k == 0 {
                fout++
            }

        }

    }

    return fout
}
