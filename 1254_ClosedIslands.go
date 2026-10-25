func closedIsland(grid [][]int) int {
    var islands int
    n := len(grid)
    m := len(grid[0])
    //mutate grid in place to save memory.

    for k0 := 0; k0 < n*m; k0++ {
        r0 := k0 / m
        c0 := k0 % m
        if grid[r0][c0] < 1 {
            count := true
            q := []int{k0}
            grid[r0][c0] = 1
            for i := 0; i < len(q); i++ {
                k := q[i]
                r := k / m
                c := k % m
                if count {
                    if r == 0 || r == n - 1 || c == 0 || c == m - 1 {
                        count = false
                    }
                }
                if r > 0 {
                    if grid[r-1][c] == 0 {
                        q = append(q,m*(r-1)+c)
                    }
                    grid[r-1][c] = 1
                }
                if r < n-1 {
                    if grid[r+1][c] == 0 {
                        q = append(q,m*(r+1)+c)
                    }
                    grid[r+1][c] = 1
                }
                if c > 0 {
                    if grid[r][c-1] == 0 {
                        q = append(q,m*r+c-1)
                    }
                    grid[r][c-1] = 1
                }
                if c < m-1 {
                    if grid[r][c+1] == 0 {
                        q = append(q,m*r+c+1)
                    }
                    grid[r][c+1] = 1
                }
            }
            if count {
                islands++
            }
        }
    }
    return islands
}
