/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func balanceBST(root *TreeNode) *TreeNode {
    //Step 0: Degenerate case
    if root == nil {
        return root
    }
    //Step 1: Extract values
    vals := make([]int,0)
    queue := []*TreeNode{root}
    for i := 0; i < len(queue); i++ {
        node := queue[i]
        vals = append(vals,node.Val)
        if node.Left != nil {
            queue = append(queue,node.Left)
        }
        if node.Right != nil {
            queue = append(queue,node.Right)
        }
    }

    //Step 2: Sort vals
    slices.Sort(vals)

    //Step 3: Convert sorted vals array into BST. This is Leetcode Problem 108.
    return sortedArrayToBST(vals)
}

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
