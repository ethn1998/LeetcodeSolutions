func minScore(n int, roads [][]int) int {
    //bidirectional bfs
    fout := 10000
    mat := make([]map[int]int,n+1)
    for i := range mat {
        mat[i] = make(map[int]int)
    }
    for _,road := range roads {
        i := road[0]
        j := road[1]
        mat[i][j] = road[2]
        mat[j][i] = road[2]
    }
    seen := make([]bool,n+1) //track if city has been visited
    q := []int{1,n}
    seen[1] = true
    seen[n] = true
    for i := 0; i < len(q); i++ {
        src := q[i]
        for dst,score := range mat[src] {
            if !seen[dst] {
                q = append(q,dst)
                seen[dst] = true
            }
            if fout > score {
                fout = score
            }
        }
    }
    return fout
}
