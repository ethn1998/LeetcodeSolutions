func findRestaurant(list1 []string, list2 []string) []string {
    var s int
    fout := make([]string,0)
    restaurants1 := make(map[string]int)
    restaurants2 := make(map[string]int)
    for i,restaurant := range(list1){
        restaurants1[restaurant] = i
    }    
    for i,restaurant := range(list2){
        restaurants2[restaurant] = i
    }
    minIndexSum := len(list1) + len(list2)
    for restaurant,i := range(restaurants1){
        j,ok := restaurants2[restaurant]
        if ok {
            s = i+j
            if s < minIndexSum {
                minIndexSum = s
                fout = []string{restaurant}
            } else if s == minIndexSum{
                fout = append(fout,restaurant)
            }
        }
    }
    return fout
}
