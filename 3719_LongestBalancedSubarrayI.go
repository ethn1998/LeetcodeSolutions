/*
IDEA:
Two pointer??
Init using i = 0: gradually adding elements until end of nums and tracking fout.
For i = k, where k > 0:
left: delete nums[k-1]
right: delete nums[i+fout+1],...,nums[n-1]
*/

func longestBalanced(nums []int) int {
	var fout int
	n := len(nums)
	odds := make(map[int]int)
	evens := make(map[int]int)

	//i = 0
	for j, val := range nums {
		if val%2 == 0 {
			evens[val]++
		} else {
			odds[val]++
		}
		if len(odds) == len(evens) {
			fout = j + 1
		}
	}

	//i > 0
	var dval, val int //value to be deleted from memory
	for i := 1; i < n; i++ {
		if i+fout > n {
			break
		}
		dval = nums[i-1]
		if dval%2 == 0 {
			evens[dval]--
			if evens[dval] == 0 {
				delete(evens, dval)
			}
		} else {
			odds[dval]--
			if odds[dval] == 0 {
				delete(odds, dval)
			}
		}
		for j := n - 1; j >= i+fout; j-- {
			dval = nums[j]
			if dval%2 == 0 {
				evens[dval]--
				if evens[dval] == 0 {
					delete(evens, dval)
				}
			} else {
				odds[dval]--
				if odds[dval] == 0 {
					delete(odds, dval)
				}
			}
		}
		for j := i + fout; j < n; j++ {
			val = nums[j]
			if val%2 == 0 {
				evens[val]++
			} else {
				odds[val]++
			}
			if len(odds) == len(evens) && fout < j - i + 1 {
				fout = j - i + 1
			}
		}
	}
	return fout
}
