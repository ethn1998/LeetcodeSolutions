func countNegatives(grid [][]int) int {
    var fout int
    for _,row := range(grid){
        for _,v := range(row){
            if v < 0 {
                fout++
            }
        }
    }
    return fout
}
