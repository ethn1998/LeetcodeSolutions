func firstCompleteIndex(arr []int, mat [][]int) int {
    n := len(mat)
    m := len(mat[0])
    arrtomat := make(map[int][]int)
    for i,row := range(mat){
        for j,v := range(row){
            arrtomat[v] = []int{i,j} //Store cell of value
        }
    }
    rowcount := make(map[int]int)
    colcount := make(map[int]int)
    for i,v := range(arr){
        rowcount[arrtomat[v][0]]++
        colcount[arrtomat[v][1]]++
        if rowcount[arrtomat[v][0]] == m || colcount[arrtomat[v][1]] == n{
            return i
        }
    }
    return len(arr) //Redundancy because Go always wants a return value outside of loops
}
