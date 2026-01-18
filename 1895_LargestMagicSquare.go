//dp? not intuitive how recursion relation works.

func largestMagicSquare(grid [][]int) int {
    var magic bool
    rows := len(grid)
    cols := len(grid[0])
    maxl := min(rows,cols)
    //fmt.Println("dim:",rows,cols,maxl)

    for l := maxl; l > 1; l-- { //check l*l squares
        for i0 := 0; i0 <= rows - l ; i0++ {
            for j0 := 0; j0 <= cols - l ; j0++ {
                //fmt.Println("In:",i0,j0,l)
                magic = isMagicSquare(grid,i0,j0,l)
                //fmt.Println("Out:",magic)
                if magic {
                    return l
                }
            }
        }
    }
    return 1
}

func isMagicSquare(grid [][]int, i0, j0, l int) bool {
    target := 0
    for k := 0; k < l; k++ {
        target += grid[i0+k][j0+k]
    }
    sum := 0
    for k := 0; k < l; k++ {
        sum += grid[i0 +l -k -1][j0 + k]
    }
    if sum != target {
        return false
    }
    //rowsum
    for i := i0; i < i0 + l; i++{
        sum = 0
        for k := 0; k < l; k++{
            sum += grid[i][j0+k]
        }
        if sum != target {
            return false
        }
    }
    //colsum
    for j := j0; j < j0 + l; j++{
        sum = 0
        for k := 0; k < l; k++{
            sum += grid[i0+k][j]
        }
        if sum != target {
            return false
        }
    }
    return true
}
