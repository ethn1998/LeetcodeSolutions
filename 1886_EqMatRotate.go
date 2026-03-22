func findRotation(mat [][]int, target [][]int) bool {
    n := len(mat) //NOTE: Square matrix
    if n == 1 {
        return mat[0][0] == target[0][0]
    }
    for r := 0; r < 4; r++ {
        eq := true
        for i, row := range target {
            for j, val := range row {
                if mat[i][j] != val {
                    eq = false
                    break
                }
            }
        }
        if eq {
            return true
        } else { //rotate
            new := make([][]int,n)
            for i := 0; i < n; i++ {
                new[i] = make([]int,n)
            }
            //transpose
            for i := 0; i < n; i++ {
                for j := 0; j < n; j++ {
                    new[i][j] = mat[j][i]
                }
            }
            //reverse
            for i := 0; i < n; i++ {
                for j := 0; j < n; j++{
                    mat[i][j] = new[i][n-1-j]
                }
            }
        }
    }
    return false
}
