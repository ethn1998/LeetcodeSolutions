func leftRightDifference(nums []int) []int {
    n := len(nums)
    answer := make([]int,n)
    cumSum := make([]int,n)
    var tot int
    for i, val := range nums {
        if i > 0 {
            cumSum[i] = val + cumSum[i-1]
        } else { //i == 0
            cumSum[0] = val
        }
        tot += val
    }
    //lSum[i] = cumSum[i-1] if i > 0, 0 otherwise
    //rSum[i] = tot - cumSum[i]
    var lSum, rSum int
    for i, sum := range cumSum {
        if i > 0 {
            lSum = cumSum[i-1]
        } //no need else zero case because Go defaults to zero value upon defn
        rSum = tot - sum
        if lSum > rSum {
            answer[i] = lSum - rSum
        } else {
            answer[i] = rSum - lSum
        }
    }
    return answer
}
