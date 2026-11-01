func minMoves(nums []int) int {
    n0 := math.MaxInt32 //based on problem statement
    for _,n := range nums {
        if n0 > n {
            n0 = n
        }
    }
    var moves int
    for _,n := range nums {
        moves += n-n0
    }
    return moves
}
