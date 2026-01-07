/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Given a fixed total sum a + b, a * b is maximized when a = b
To find a, b s.t. a, b = (a + b) / 2 (approx)
*/

func maxProduct(root *TreeNode) int {
	totSum := getSum(root)
	var fout, sum int
	queue := []*TreeNode{root}
	//fmt.Println(totSum, totSum/2)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
        //fmt.Println(node)
        if node.Left != nil {
    		sum = node.Left.Val
        	if sum*(totSum-sum) > fout {
		    	fout = sum * (totSum - sum)
    		}
            if sum > totSum / 2 {
                queue = append(queue, node.Left)
            }
        }
        if node.Right != nil {
            sum = node.Right.Val
    		if sum*(totSum-sum) > fout {
	    		fout = sum * (totSum - sum)
    		}
            if sum > totSum / 2 {
                queue = append(queue, node.Right)
            }
        }
	}
	return fout % 1000000007
}

func getSum(root *TreeNode) int {
	//fmt.Println("In:",root)
	/*if root == nil {
		return 0
	}
	lsum := getSum(root.Left)
	rsum := getSum(root.Right)

	//return root.Val + lsum + rsum
    */
    //try to update to store sum instead of value
    if root.Left != nil {
        root.Val += getSum(root.Left)
    }
    if root.Right != nil {
        root.Val += getSum(root.Right)
    }
	return root.Val
}
