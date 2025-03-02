func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    //This is exactly the same as 2570
    retmap := make(map[int]int)
    for _, item := range(items1){
        retmap[item[0]] += item[1]
    }
    for _, item := range(items2){
        retmap[item[0]] += item[1]
    }
    ret := make([][]int,0)
    for v,w := range(retmap){
        ret = append(ret,[]int{v,w})
    }
    sort.Slice(ret, func(i,j int) bool{
        return ret[i][0] < ret[j][0]
    })
    return ret

}
