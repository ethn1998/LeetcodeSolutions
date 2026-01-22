//This is an easy question so we should be able to just simulate with brute force

func minimumPairRemoval(nums []int) int {
    //common sense cases
    if len(nums) <= 1 {
        return 0
    } else if len(nums) == 2 {
        if nums[0] > nums[1] {
            return 1
        } else {
            return 0
        }
    }
    nums1 := make([]int,len(nums)) //make a copy for robustness
    l := copy(nums1,nums)
    var sum int
    for i := 0; i < l; i++ {
        //fmt.Println("In:",i,nums1)
        sorted := true
        //minSum := 2000 //follow problem constraints. THIS IS WRONG BECAUSE SUBSEQUENT SUMS CAN BECOME GREATER THAN 2000 AFTER MULTIPLE OPERATIONS
        minSum := 1 << 63 -1 //USE BIGGEST POSSIBLE INT64.
        k := 0
        for j := 0; j < len(nums1)-1; j++ {
            n0 := nums1[j]
            n1 := nums1[j+1]
            if n1 < n0 {
                sorted = false
            }
            sum = n0 + n1
            if sum < minSum {
                k = j
                minSum = sum
            }
        }
        //fmt.Println("Int:",k,minSum,sorted)
        if sorted {
            return i
        }
        nums1[k] = minSum
        //fmt.Println("Out:",nums1)
        nums1 = append(nums1[:k+1],nums1[k+2:]...)
    }
    return l-1 //redundancy
}
