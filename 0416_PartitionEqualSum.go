import "sort"

/*
IDEA: 
Try some kind of binary search? Might not be appropriate
Try some kind of recursive DFS?
*/

func canPartition(nums []int) bool {
    if len(nums) < 2 {
        return false
    }
    sort.Ints(nums) //From example 1, we can choose to not preserve original order.
    var s0 int //Track sums
    //cumsum := make([]int,len(nums))
    for _,v := range(nums){
        s0 += v
        //cumsum[i] = s0
    }
    if s0 % 2 == 1 { //Common sense cases
        return false
    }
    target := s0 / 2 //This breaks if s0 is odd.
    //Brute force looks like a memory hog
    mem := make(map[int]bool)
    for _,v := range nums {
        new := []int{}
        for u := range(mem) {
            if u + v == target{
                return true
            }
            if u + v < target {
                new = append(new,u+v)
            }
        }
        for _,w := range(new){
            mem[w] = true
        }
        if v == target {
            return true
        }
        mem[v] = true
    }   
    return false
}
