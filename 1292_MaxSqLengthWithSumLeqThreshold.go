func maxSideLength(mat [][]int, threshold int) int {
    nRow := len(mat)
    nCol := len(mat[0])
    //preprocess cumsums
    cumsums := make([][]int,nRow)
    for i, row := range mat {
        cumsums[i] = make([]int,nCol+1)
        for j, v := range row {
            cumsums[i][j+1] = cumsums[i][j] + v
        }
    }
    //fmt.Println(nRow,nCol)
    //fmt.Println(cumsums)
    left := 0
    right := min(nRow,nCol)
    for left < right {
        mid := left + (right - left)/2
        if right == left + 1 {
            mid = right
        }
        yes := exist(cumsums,threshold,mid)
        //fmt.Printf("l: %d m: %d r: %d out: %s\n",left,mid,right,yes)
        if yes {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return left
}

func exist(cumsums [][]int, threshold, l int) bool {
    nRow := len(cumsums)
    nCol := len(cumsums[0])-1
    
    for i := 0; i <= nRow - l; i++ {
        for j := 0; j <= nCol - l; j++ {
            var sqsum int
            for k := 0; k < l; k++ {
                sqsum += cumsums[i+k][j + l] - cumsums[i+k][j]
            }
            //fmt.Println(i,j,sqsum)
            if sqsum <= threshold {
                return true
            }
        }
    }
    return false
}
