/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    IDEA: Balance.

    wlog, assume left depth > right depth. 
    Then we know that the answer is subtree of left tree.
    Same if right depth > left depth

    If balanced, then we found an answer
 */


func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    depths := make(map[int]int) //NOTE: node vals are unique -> key, val = depth
    _ = getMaxDepth(root,0,depths)

    queue := []*TreeNode{root}
    for i := 0; i < len(queue); i++ {
        node := queue[i]
        if node.Left != nil && node.Right != nil {
            if depths[node.Left.Val] < depths[node.Right.Val] {
                queue = append(queue,node.Right)
            } else if depths[node.Left.Val] > depths[node.Right.Val] {
                queue = append(queue,node.Left)
            } else {
                return node
            }
        } else {
            if node.Left != nil {
                queue = append(queue,node.Left)
            } else if node.Right != nil {
                queue = append(queue,node.Right)
            } else {
                return node
            }
        }
    }
    return queue[len(queue)-1] //redundancy
}

func getMaxDepth(node *TreeNode, n int, mem map[int]int) int {
    if node == nil {
        return n
    }
    left := getMaxDepth(node.Left, n+1, mem)   
    right := getMaxDepth(node.Right, n+1, mem)
    fout := max(left,right)
    mem[node.Val] = fout
    return fout
}
