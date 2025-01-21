func containsDuplicate(nums []int) bool {
    mem := make(map[int]bool)
    for _,num := range(nums){
        if mem[num]{
            return true
        }else{
            mem[num] = true
        }
    }
    return false
}
