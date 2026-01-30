func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
    //fmt.Println(len(source),len(cost))//DEBUG
	var src, trans, dst string
	var srctotrans, transtodst int64

	costmat := make(map[string]map[string]int64, len(original))
	for _, src = range original {
		_, ok := costmat[src]
		if !ok {
			costmat[src] = make(map[string]int64)
			//costmat[src][src] = 0 //redundant?
			queue := []string{src}
			for i := 0; i < len(queue); i++ {
				trans = queue[i]
				srctotrans = costmat[src][trans]
				for j := 0; j < len(original); j++ {
					if original[j] == trans {
						dst = changed[j]
						transtodst = int64(cost[j])
						_, ok := costmat[src][dst]
						if ok {
							if costmat[src][dst] > srctotrans+transtodst {
								costmat[src][dst] = srctotrans + transtodst
								queue = append(queue, dst)
							}
						} else {
							costmat[src][dst] = srctotrans + transtodst
							queue = append(queue, dst)
						}
					}
				}
			}
		}
	}
	/*DEBUG
	for key, val := range costmat {
		fmt.Println(key, ":", val)
	}
	check := make([]bool, len(source)+1) //redundancy for verification
	*/
	dp := make([]int64, len(source)+1) //cost to mutate first 'right' elements to target
	var prev int64
	for right := 1; right < len(dp); right++ {
		//fmt.Println("In:", right, source[0:right], target[0:right]) //DEBUG
		dp[right] = -1 //default output for nonexistent solution
		if source[right-1] == target[right-1] {
            dp[right] = dp[right-1]
        }        
        for left := 0; left < right; left++ { //NOTE: THIS IMPLEMENTATION IS INEFFICIENT
			//sanity check if partial solution exists
            prev = dp[left]
            //fmt.Println("idx:",left,prev) //DEBUG
			if prev > -1 { //only proceed if operation is valid
				src = source[left:right]
				dst = target[left:right]
				dsts, ok := costmat[src]
                //fmt.Println(src,dst,dsts,ok) //DEBUG
				if ok {
					val, ok := dsts[dst]
					if ok {
						if dp[right] > -1 {
							dp[right] = min(dp[right], prev+val)
						} else {
							dp[right] = prev + val
                            //check[right] = true //DEBUG
						}
					}
				}
			}
			//fmt.Println("Out:", dp[right], check[right]) //DEBUG
		}
		//fmt.Println("Final Out:", right, source[0:right], target[0:right], dp[right], check[right]) //DEBUG
	}
    //DEBUG
	//fmt.Println(dp) 
	//fmt.Println(check)
	//
    return dp[len(source)]
}
