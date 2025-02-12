func maximumSum(nums []int) int {
    //Group values in nums based on digit sums.
    var n,digitsum int
    groups := make(map[int][]int)
    for _,num := range(nums){
        digitsum = 0
        n = num
        for n > 0 {
            digitsum += n % 10
            n /= 10
        }
        _,ok := groups[digitsum] //Check existence of key for stability
        if ok {
            groups[digitsum] = append(groups[digitsum],num)
        } else {
            groups[digitsum] = []int{num}
        }
    }
    fout := -1
    for _,group := range(groups){
        if len(group) >=2 {
            //Record two biggest numbers in group. Sum of the two biggest is the maximum value attainable by group.
            alpha := 0 //First biggest number
            beta := 0 //Second biggest number
            for _, val := range(group){
                if val > alpha {
                    beta = alpha
                    alpha = val
                } else if val > beta {
                    beta = val
                }
            }
            if alpha + beta > fout{
                fout = alpha + beta
            }
        }
    }
    return fout
}
