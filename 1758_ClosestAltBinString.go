func minOperations(s string) int {
    var x0, x1 int //start from 0, start from 1
    for i,char := range s {
        if i % 2 == 0 {
            if char == '0' {
                x1++
            } else {
                x0++
            }
        } else {
            if char == '0' {
                x0++
            } else {
                x1++
            }
        }
    }
    return min(x0,x1)
}
