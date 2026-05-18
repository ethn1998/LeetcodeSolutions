//bfs?

func minJumps(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}
	//todo: check neighbors when initializing dsts
	dsts := make(map[int][]int)
	for i, val := range arr {
		yes := false
		//check left
        if i == 0 {
            yes = true
        } else {
            if arr[i-1] != val {
                yes = true
            }
        }
		//check right
        if i == n-1 {
            yes = true
        } else {
            if arr[i+1] != val {
                yes = true
            }
        }
        if yes {
			_, ok := dsts[val]
			if ok {
				dsts[val] = append(dsts[val], i)
			} else {
				dsts[val] = []int{i}
			}
		}
	}
	memo := make([]bool, n)
	memo[0] = true
	var fout int
	curr := []int{0}
	for true {
		next := make([]int, 0)
		for _, i := range curr {
			val := arr[i]
			if i == n-1 {
				return fout
			} else {
				if !memo[i+1] && arr[i+1] != val {
					next = append(next, i+1)
					memo[i+1] = true
				}
			}
			if i > 0 {
				if !memo[i-1] && arr[i-1] != val {
					next = append(next, i-1)
					memo[i-1] = true
				}
			}
			for _, dst := range dsts[val] {
				if !memo[dst] {
					next = append(next, dst)
					memo[dst] = true
				}
			}
		}
		curr = next
		fout++
	}
	return fout
}
