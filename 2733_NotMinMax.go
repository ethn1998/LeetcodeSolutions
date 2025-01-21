func findNonMinOrMax(nums []int) int {
    if len(nums) <= 2 {//force max or min
        return -1
    }
    var v, vmax int
    vmin := 100 //Check min
    for _,n := range(nums){
        if n < vmin{
            vmin = n
        }
        if n > vmax{
            v = vmax
            vmax = n
        }else if n > v{
            v = n
        }
    }
    return v
}
