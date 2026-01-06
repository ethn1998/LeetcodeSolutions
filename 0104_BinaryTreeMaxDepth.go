/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return depth(0, root)
}

func depth(n int, root *TreeNode) int { //n to track depth of tree
    //fmt.Println(n, root)
    if root == nil {
        return n
    }
    l := depth(n+1, root.Left)
    r := depth(n+1, root.Right)

    return max(l,r)
}
