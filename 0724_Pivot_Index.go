func pivotIndex(nums []int) int {
    var runsum0 int
    runsum := make([]int,len(nums))
    for i,v := range(nums){
        runsum0 += v
        runsum[i] = runsum0 
    }
    //runsum0 will have the sum of all elements in nums at the end.
    for j,s := range(runsum){
        if j == 0 {
            if runsum0 == s { // If sum of everything else is zero 
                return 0
            }
        } else if j < len(nums)-1 {
            if runsum[j-1] == runsum0 - s {
                return j
            }
        } else {
            if runsum[j-1] == 0{
                return len(nums)-1
            }
        }
    }
    return -1 //If pivot does not exist
}
