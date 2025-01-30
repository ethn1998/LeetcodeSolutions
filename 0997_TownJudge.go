func findJudge(n int, trust [][]int) int {
    outEdges := make([]int,n)
    inEdges := make([]int,n)
    for _,edge := range(trust){
        src := edge[0] -1
        dst := edge[1] -1
        outEdges[src] += 1
        inEdges[dst] += 1
    }
    townjudge := -1 //Default to non-existent
    for i,v := range(outEdges){
        if v == 0 { //(1) The Town Judge trusts no one.
            if inEdges[i] == n-1 { //(2) Everyone but the Town Judge trusts the Town Judge
                if townjudge == -1 {
                    townjudge = i+1
                } else { //(3) The Town Judge is unique
                    return -1
                }
            }
        }
    }
    return townjudge
}
