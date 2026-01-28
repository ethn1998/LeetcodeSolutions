/*
This is actually a 3D dp problem. i,j,k where k is number of teleports forms the 3rd dimension 
NOTE: This 3D dp idea works but it checks all cases in an inefficient manner. Editorial + hint is more efficient because it precalculates BEST teleport position and goes from there.

type sq struct {
    x int
    y int
}

type state struct {
    xy sq
    z int //number of teleports remaining
}

func minCost(grid [][]int, k int) int {
    m := len(grid)
    n := len(grid[0])
        
    sqs := make([]sq, 0, m*n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            sqs = append(sqs, sq{i,j})
        }
    }
    sort.Slice(sqs, func(i,j int) bool {
        return grid[sqs[i].x][sqs[i].y] < grid[sqs[j].x][sqs[j].y]
    })
    n0s := make(map[int]int)
    for i, sq := range sqs {
        val := grid[sq.x][sq.y]
        //fmt.Println(i,sq0,val) //DEBUG
        _, ok := n0s[val]
        if !ok {
            n0s[val] = i
        }
    }

    mem := make(map[state]int) //visited states
    for t := 0; t <= k; t++ {
        mem[state{sq{0,0},t}] = 0
    }

    return dp(state{sq{m-1,n-1},k}, grid, mem, sqs, n0s)
    //return dp(sq{m-1,n-1}, k, grid, sqs, n0s)
}

func dp(key state, grid [][]int, mem map[state]int, sqs []sq, n0s map[int]int) int {
    v, ok := mem[key]
    if ok { //previously visited
        return v
    }
    i := key.xy.x
    j := key.xy.y
    k := key.z
    v = grid[i][j]
    fout := math.MaxInt
    if k > 0 {
        for n := n0s[v]; n < len(sqs); n++ { //to optimize this
            sq1 := sqs[n]
            if sq1.x != i || sq1.y != j {
                fout = min(fout,dp(state{sq1,k-1},grid,mem,sqs,n0s))
            }
            //fmt.Printf("(%d,%d,%d) -> (%d,%d,%d): %d\n",sq1.x,sq1.y,k-1,i,j,k,fout) //DEBUG
        }
    }
    if i > 0 {
        fout = min(fout,v+dp(state{sq{i-1,j},k},grid,mem,sqs,n0s))
    }
    if j > 0 {
        fout = min(fout,v+dp(state{sq{i,j-1},k},grid,mem,sqs,n0s))
    }
    mem[key] = fout
    return fout
}
*/

//Below follows hints/editorial

type sq struct {
    x int
    y int
}

func minCost(grid [][]int, k int) int {
    m := len(grid)
    n := len(grid[0])
    sqs := make([]sq,0,m*n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            sqs = append(sqs,sq{i,j})
        }
    }
    sort.Slice(sqs, func(i,j int) bool { //sort teleport groups
        return grid[sqs[i].x][sqs[i].y] < grid[sqs[j].x][sqs[j].y]
    })

    /*DEBUG
    for i, sq := range sqs {
        fmt.Println(i,sq,grid[sq.x][sq.y])
    }
    */

    dp := make([][]int,m) //dp array
    for i := range dp {
        dp[i] = make([]int,n)
        for j := range dp[i]{
            dp[i][j] = 1 << 63 - 1
        }
    }
    for t := 0; t <= k; t++ { //t is number of teleports
        minCost := 1 << 63 - 1
        //teleport
        for l,r := len(sqs)-1 , len(sqs) - 1 ; l >= 0; l-- {
            minCost = min(minCost, dp[sqs[l].x][sqs[l].y])
            if l > 0 && grid[sqs[l-1].x][sqs[l-1].y] == grid[sqs[l].x][sqs[l].y] {
                continue //l--, go back to start of loop
            }
            //found terminal l
            for n := l; n <= r; n++ {
                dp[sqs[n].x][sqs[n].y] = minCost
            }
            r = l - 1
        }
        //regular traversal
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if i == 0 && j == 0 {
                    dp[0][0] = 0
                }
                if i > 0 {
                    dp[i][j] = min(dp[i][j],dp[i-1][j] + grid[i][j])
                }
                if j > 0 {
                    dp[i][j] = min(dp[i][j],dp[i][j-1] + grid[i][j])
                }
            }
        }
    }
    return dp[m-1][n-1]
}
