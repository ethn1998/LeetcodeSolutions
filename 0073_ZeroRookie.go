func setZeroes(matrix [][]int)  {
    zerorows := make(map[int]bool)
    zerocols := make(map[int]bool)
    for i,row := range matrix{
        for j,v := range row{
            if v == 0 {
                zerorows[i] = true
                zerocols[j] = true
            }
        }
    }
    for i := 0; i < len(matrix); i++{
        if zerorows[i] {
            for j := 0; j < len(matrix[0]); j++{
                matrix[i][j] = 0
            } 
        }
    }
    for j := 0; j < len(matrix[0]); j++{
        if zerocols[j] {
            for i := 0; i < len(matrix); i++{
                matrix[i][j] = 0
            }
        }
    }
}
