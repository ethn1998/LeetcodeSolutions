func maximumLength(nums []int) int {
    var ones int
	head := make(map[int]int)
	legs := make(map[int]int)
	slices.Sort(nums) //descending order to get heads first
	//fmt.Println(nums) //DEBUG
	for _, val := range slices.Backward(nums) {
		if val > 1 {
			_, hasHead := head[val*val]
			if hasHead {
				legs[val]++
			}
			if legs[val] == 2 {
				head[val] = head[val*val] + 2
			}
			if head[val] < 1 {
				head[val] = 1
			}
		} else {
            ones++
        }
	}
    var fout int
    if ones > 0 {
        if ones % 2 != 0 {
            fout = ones
        } else {
            fout = ones - 1
        }
    }

	for _, val := range head {
		if fout < val {
			fout = val
		}
	}
	return fout
}
