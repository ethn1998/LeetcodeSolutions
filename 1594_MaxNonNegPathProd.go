//dp

type duo struct {
    pos int
    neg int
}

func newDuo() *duo {
    return &duo{-1,1}
}

func update(prev *duo, val int) *duo {
    if val > 0 {
        return &duo{val * prev.pos, val * prev.neg}
    } else if val < 0 {
        return &duo{val * prev.neg, val * prev.pos}
    }
    return &duo{0,0}
}

func combine(duo0, duo1 *duo) *duo {
    pos := max(duo0.pos, duo1.pos)
    neg := min(duo0.neg, duo1.neg)
    return &duo{pos,neg}
}

func maxProductPath(grid [][]int) int {
    if grid[0][0] == 0 {
        return 0
    }
    m := len(grid)
    n := len(grid[0])
    dp := make([][]*duo, m)

    for i := 0; i < m; i++ {
        dp[i] = make([]*duo, n)
        for j := 0; j < n; j++ {
            if i > 0 || j > 0 {
                fromtop := newDuo()
                fromleft := newDuo()
                if i > 0 { //from top
                    fromtop = update(dp[i-1][j],grid[i][j])
                }
                if j > 0 { //from left
                    fromleft = update(dp[i][j-1],grid[i][j])
                }
                dp[i][j] = combine(fromtop,fromleft)
            } else {
                dp[0][0] = newDuo()
                if grid[0][0] > 0 {
                    dp[0][0].pos = grid[0][0]
                } else {
                    dp[0][0].neg = grid[0][0]
                }
            }
        }
    }
    
    fout := dp[m-1][n-1].pos
    if fout < 0 {
        return -1
    }
    return fout % 1000000007
}
