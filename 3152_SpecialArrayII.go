func isArraySpecial(nums []int, queries [][]int) []bool {
    /* 
    THIS BUGGED MESS IS AN ATTEMPT TO IMPLEMENT MEMORY OF PREVIOUSLY SEEN QUERIES TO AVOID TLE.
    BUT ENDS UP MESSING UP OTHER TEST CASES. THIS IS VERY HARD TO FIX. CODE COMMENTED OUT. DO NOT USE.
    fout := make([]bool,0)
    fromList := make([]int,0)
    toList := make([]int,0)
    for _,q := range(queries){
        if q[0] == q[1] { //Trivial case
            fout = append(fout,true)
        } else {
            unseen := true
            updateFrom := -1 //Update existing range start
            updateTo := -1 //Update existing range end
            updateIndex := -1 //Index in from/toList to be updated
            //Get indices of start and end of query
            from := q[0]
            to := q[1]
            //Check memory of previous queries.
            for i, start := range(fromList) {
                if from >= toList[i] || toList[i] < start { //Completely out of range
                    //Do nothing
                } else if from >= start {
                    if to <= toList[i] { // Query is a subset of previously seen query
                        unseen = false //No need to re-calculate
                        fout = append(fout,true)
                    } else {
                        updateIndex = i
                        updateTo = to
                        from = toList[i] //Start from endpoint
                        break
                    }
                } else {
                    updateIndex = i
                    updateFrom = from
                    if to < toList[i] {
                        to = fromList[i]
                    } else {
                        updateTo = to
                    }
                    break
                }
            }
            if unseen {
                for i := from; i < to ; i++ {
                p := nums[i] + nums[i+1]
                    if p % 2 != 1 {
                        fout = append(fout,false)
                        break
                    }
                    if i == to - 1 {
                        fout = append(fout,true)
                        if updateIndex >= 0 { //If similar memory, extend memory
                            if updateFrom >= 0 {
                                fromList[updateIndex] = updateFrom
                            }
                            if updateTo >= 0 {
                                toList[updateIndex] = updateTo
                            }
                        } else { //Make new memory
                            fromList = append(fromList,from)
                            toList = append(toList,to)
                        }
                    }
                }
            }
        }
    }
    */
    //IDEA: Try analyzing array itself, splitting into subarrays of acceptable solutions.
    var s,group int
    groups := []int{group}
    for i := 1 ; i < len(nums) ; i++ {
        s = nums[i] + nums[i-1]
        if s % 2 != 1 {
            group++
        }
        groups = append(groups,group)
    }
    fout := make([]bool,0)
    for _,q := range(queries) {
        fout = append(fout, groups[q[0]] == groups[q[1]])
   }
    return fout
}
