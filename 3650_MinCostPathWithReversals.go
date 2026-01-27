func minCost(n int, edges [][]int) int {
    costs := make([]int,n)
    //costs[0] = 0 //redundancy
    for i := 1; i < n; i++ {
        costs[i] = -1        
    }
    //process edges into (sparse) matrix for querying
    costmat := make([]map[int]int,n)
    for i := 0; i < n; i++ {
        costmat[i] = make(map[int]int)
    }
    var src, dst, cost int
    for _,edge := range edges {
        src = edge[0]
        dst = edge[1]
        cost = edge[2]
        val, ok := costmat[src][dst]
        if ok {
            costmat[src][dst] = min(val,cost)
        } else {
            costmat[src][dst] = cost
        }
        //include reversed edges
        val, ok = costmat[dst][src]
        if ok {
            costmat[dst][src] = min(val,2*cost)
        } else {
            costmat[dst][src] = 2*cost
        }
    }

    queue := []int{0}
    //var sol bool
    for i := 0; i < len(queue); i++ {
        src = queue[i]
        if src < n - 1 {
            cost = costs[src]
            for dst, val := range costmat[src] {
                if cost + val < costs[dst] || costs[dst] == -1 {
                    costs[dst] = cost + val
                    queue = append(queue,dst)
                }
            }
        }
    }
    return costs[n-1]
}
