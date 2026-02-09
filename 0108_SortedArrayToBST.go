func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    n := len(nums)
    mid := n/2
    val := nums[mid]
    lnode := sortedArrayToBST(nums[0:mid])
    rnode := sortedArrayToBST(nums[mid+1:len(nums)])
    return &TreeNode{val,lnode,rnode}
}
