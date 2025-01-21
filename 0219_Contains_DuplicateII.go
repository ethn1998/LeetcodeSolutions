func containsNearbyDuplicate(nums []int, k int) bool {
    var d int
    idict := make(map[int]int)
    for i,v := range(nums){
        _, kp := idict[v]
        if !kp { //If value not seen yet
            idict[v] = i //Declare new value
        } else {
            d = i - idict[v] //Check distance between current index and last seen
            if d <= k { //If within distance
                return true
            } else {
                idict[v] = i //Update value
            }
        }
    }
    return false    
}
