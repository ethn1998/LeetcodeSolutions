/*
Vectorize the matrix s.t. the ith element of the vector is the biggest index where the ith row has value 1.
Sort the vector MANUALLY.
Check vector.
Worst case allowed is (0,1,2,3,4,...,n). Earlier indices are most restrictive.
If vector[j] > j, return -1.
Else, return number of steps needed to sort vector
*/

func minSwaps(grid [][]int) int {
	n := len(grid)
    if n == 1 {
        return 0
    }
    idxs := make([]int,n)
    for i, row := range grid {
        for j, val := range row {
            if val == 1 {
                idxs[i] = j
            }
        }
    }
    //simulate sort
    var nSteps int
    for i := 0; i < n; i++ {
        //fmt.Println("In:",i,idxs) //DEBUG
        m := len(idxs)
        if idxs[0] > i { //need to swap
            //find earliest legal swap
            k := -1
            for j := 1; j < m; j++ {
                if idxs[j] <= i {
                    k = j
                    break
                }
            }
            if k == -1 { //if no legal swaps found
                return -1
            }
            nSteps += k
            idxs = append(idxs[0:k],idxs[k+1:m]...)
        } else {
            idxs = idxs[1:m]
        }
    }
    return nSteps
}
