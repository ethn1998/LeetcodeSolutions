/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumRootToLeaf(root *TreeNode) int {
	var fout int
    vals := getVals(root,0,[]int{})
    for _,val := range vals {
        fout += val
    }
    return fout
}

func getVals(root *TreeNode, x int, vals []int) []int {
	x = x << 1
	x += root.Val
	if root.Left == nil && root.Right == nil {
		vals = append(vals, x)
	} else {
		if root.Left != nil {
			vals = getVals(root.Left, x, vals)
		}
		if root.Right != nil {
			vals = getVals(root.Right, x, vals)
		}
	}
	return vals
}
