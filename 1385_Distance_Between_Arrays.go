import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    var fout int
    for _, n1 := range(arr1){
        add := true //innocent until proven guilty
        for _, n2 := range(arr2) {
            if math.Abs(float64(n1-n2)) <= float64(d){ //check for existence
                add = false
                break
            }
        }
        if add {
            fout++
        }
    }
    return fout
}
