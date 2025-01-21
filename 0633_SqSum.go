//import math

func judgeSquareSum(c int) bool {
    var i,j,s int
    for j*j < c{
        j++
    }
    for i <= j {
        s = i*i + j*j
        if s == c{
            return true
        } else if s < c{
            i++
        } else{
            j--
        }
    }
    return false
}
