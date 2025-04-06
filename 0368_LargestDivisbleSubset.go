import "sort"

func largestDivisibleSubset(nums []int) []int {
    //sort.Slice(nums, func(i,j int) bool {return nums[i]<nums[j]}) //Sort in ascending order
    sort.Ints(nums)
    
    best := make([][]int,len(nums))
    var ifout, alpha int //variable to track index and size of largest subset
    
    for i,v := range(nums) {
        parent := -1
        depth := 0
        for j := 0; j < i; j++ { //Look for j that 
            if nums[i] % nums[j] == 0 {
                if len(best[j]) + 1 > depth {
                    parent = j
                    depth = len(best[j])+1
                }
            }
        }
        if parent != -1 {
            child := make([]int,len(best[parent]))
            copy(child,best[parent])
            best[i] = append(child,v)
        } else {
            best[i] = []int{v}
        }
        if len(best[i]) + 1 > alpha {
            ifout = i
            alpha = len(best[i]) + 1
        }
    }
    return best[ifout]
}
