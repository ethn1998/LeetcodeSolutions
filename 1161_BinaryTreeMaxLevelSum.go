/*
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxLevelSum(root *TreeNode) int {
    var fout int
    omega := -1 << 63
    k := 1
    scores := myfunc(k, root)
    
    var stop bool
    for !stop {
        n, ok := scores[k]
        if !ok {
            stop = true
        } else {
            if n > omega {
                fout = k
                omega = n
            }
        }
        k++
    }

    return fout
}

func myfunc(n int, root *TreeNode) map[int]int {
    fout := make(map[int]int)
    if root == nil {
        return fout
    }
    fout[n] = root.Val
    left := myfunc(n+1, root.Left)
    right := myfunc(n+1, root.Right)
    for key, val:= range left {
        fout[key] += val
    }
    for key, val:= range right {
        fout[key] += val
    }
    return fout
}
