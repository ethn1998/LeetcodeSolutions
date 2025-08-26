func areaOfMaxDiagonal(dimensions [][]int) int {
    var fout int
    var h2max int //square is monotone increasing function
    for _, dims := range dimensions {
        h2 := dims[0]*dims[0] + dims[1]*dims[1]
        if h2 > h2max {
            h2max = h2 //update
            fout = dims[0]*dims[1]
        } else if h2 == h2max {
            if dims[0]*dims[1] > fout {
                fout = dims[0]*dims[1]
            }
        }
    }
    return fout
}
