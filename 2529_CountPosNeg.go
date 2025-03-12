func maximumCount(nums []int) int {
    var pos, neg int
    for _,v := range(nums){
        if v < 0 {
            neg++
        }
        if v > 0 {
            pos++
        }
    }
    if neg > pos{
        return neg
    }
    return pos
}
