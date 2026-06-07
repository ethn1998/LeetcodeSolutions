/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//IDEA: Kahn's algorithm?

func createBinaryTree(descriptions [][]int) *TreeNode {
    iskid := make(map[int]bool) //0: parent, 1: kid
    kids := make(map[int][][]int)
    var parent, kid, isLeft int
    for _,data := range descriptions {
        parent = data[0]
        kid = data[1]
        isLeft = data[2]
        kids[parent] = append(kids[parent],[]int{kid,isLeft})
        iskid[data[1]] = true
    }
    var rootval int
    for key := range kids {
        _,ok := iskid[key]
        if !ok {
            rootval = key
            break
        }
    }
    root := &TreeNode{rootval,nil,nil}
    q := []*TreeNode{root}

    for i := 0; i < len(q); i++ {
        node := q[i]
        parent = node.Val
        for _,child := range kids[parent] {
            kid = child[0]
            newNode := &TreeNode{kid,nil,nil}
            isLeft = child[1]
            if isLeft == 1 { //left
                node.Left = newNode
            } else {
                node.Right = newNode
            }
            q = append(q,newNode)
        }
    }
    return root
}
