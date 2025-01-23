func countServers(grid [][]int) int {
    var fout int
    toCheck := make(map[int]int) //Store indices to be checked columnwise
    for i,row := range(grid) {
        rowSum := 0 //Number of servers in row
        j0 := 0 //Check index
        for j,v := range(row) {
            if v != 0 {
                rowSum++
                j0 = j
            }
        }
        if rowSum > 1 {
            fout += rowSum
        } else if rowSum == 1 {
            toCheck[i] = j0 //Store index to be checked columnwise. Row i column j0
        }
    }
    for key, val := range(toCheck){
        for k := 0; k < len(grid); k++{
            if (k != key) && (grid[k][val] == 1){ //Check existence of another server in column
                fout += 1
                break
            }
        }
    }
    return fout
}
