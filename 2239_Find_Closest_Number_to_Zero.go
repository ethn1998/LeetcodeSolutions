func abs(n int) int{ //Defining my own abs function because Go doesn't have this built in.
    if n < 0 {
        return -n
    }
    return n
}

func findClosestNumber(nums []int) int {
    var m int
    //Initialize output using constraints
    s, closest := 100001, 100001
    
    for i := 0; i < len(nums); i++ {
        m = nums[i]
        if abs(m)<s {
            closest = m
            s = abs(m)
        } else if abs(m) == s {
            if m > closest {
                closest = m
            }
        }
    }
    return closest
}
