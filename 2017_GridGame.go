func gridGame(grid [][]int) int64 {
    // NOTE: Grid is 2 x n. Robots can only move down or right
    // Only decision is to choose when to move down.
    // IDEA1: A greedy algorithm maximizing robot 1 should solve this. NOTE: THIS DOES NOT WORK
    // IDEA2: Make robot 1 evil and only think about the score of robot 2.

    fout := 5555555555 //Upper bound on maximum possible score
    var n,s0,s1 int
    for _,val := range(grid[0]){
        s0 += val
        n++
    }

    for i := 0; i < n; i++{ //i is index when robot 1 moves down
        if i > 0 {
            s1 += grid[1][i-1]
        }
        s0 -= grid[0][i]

        if s0 > s1 {//Robot 2 chooses row 0
            if s0 < fout{
                fout = s0
            }
        }else{
            if s1 < fout{
                fout = s1
            }
        }
    }
    return int64(fout)
}
