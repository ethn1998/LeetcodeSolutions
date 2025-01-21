func majorityElement(nums []int) int {
    var fout, Omega int
    mem := make(map[int]int)

    if len(nums) == 1 {
        return nums[0]
    }

    T := len(nums)/2 //Integer division in Go?
    for i := 0; i < len(nums); i++{
        mem[nums[i]] += 1
        if mem[nums[i]] > T{
            return nums[i]
        }
    }

    for num, freq := range(mem){ //Redundancy for compilation
        if freq > Omega{
            fout = num
            Omega = freq
        }
    }
    return fout
}
