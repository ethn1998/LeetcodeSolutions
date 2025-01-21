func thirdMax(nums []int) int {
    set := make(map[int]bool)
    for _,n := range(nums){
        set[n] = true
    }

    //Initialize outputs
    b0 := -1
    b1 := -1
    b2 := -1
    for i := 0 ; i < 31; i++{
        b0 *= 2
        b1 *= 2
        b2 *= 2
    } 

    N := len(set)
    if N == 1 {
        return nums[0] //Only possible value
    } else if N == 2 {
        fout := b0
        for k := range set{
            if k > fout{
                fout = k
            }
        }
        return fout
    }
    for k := range(set){
        if k > b0{
            b2 = b1
            b1 = b0
            b0 = k
        }else if k > b1{
            b2 = b1
            b1 = k
        }else if k > b2{
            b2 = k
        }
    }
    return b2
}
