/*
IDEA:
Assume nums is sorted (can sort in place at O(nlogn))
Assume nums[i] is the minimum, that is, we discard the left i elements. 
The wanted maximum is <= k*nums[i].
Binary search to find smallest j s.t. nums[j] > k*nums[i]?
*/

func minRemoval(nums []int, k int) int {
	n := len(nums)
	if n == 1 { //common sense case
		return 0
	}
	slices.Sort(nums)
	if n == 2 { //common sense case
		if k*nums[0] < nums[1] {
			return 1
		} else {
			return 0
		}
	}
	fmt.Println("sorted:", nums) //DEBUG
	fout := 1<<63 - 1
	removed := n - 1
	var left, mid, right int
	for i, m := range nums {
		//fmt.Println("in:", i, m) //DEBUG
		left = i
		right = n
		for left < right {
			mid = left + (right-left)/2
			//fmt.Printf("bin: left:%d mid:%d right:%d cmax:%d\n",left,mid,right,nums[mid])
			if nums[mid] > k*m {
				right = mid
			} else {
				left = mid + 1
			}
		}
		//To remove i from LHS of array and n-right from RHS of array
		removed = i + n - right
		if fout > removed {
			fout = removed
		}
        //fmt.Println("out:",fout) //DEBUG
		if fout == 0 {
			break
		}
	}
	return fout
}
