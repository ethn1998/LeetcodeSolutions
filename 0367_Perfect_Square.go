func isPerfectSquare(num int) bool {
    var m int
    for{
        if (m*m < num){
            m++
        }else if (m*m > num){
            return false
        } else {
            return true
        }
    }    
}
