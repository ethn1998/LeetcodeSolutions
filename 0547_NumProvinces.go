func findCircleNum(isConnected [][]int) int {
    /*
    isConnected[i][j] means that there is an edge between node i and node j
    */
    var provinces int
    seen := make([]bool,len(isConnected)) //Number of nodes is dimension of matrix
    for i,t := range(seen){
        if !t {
            seen[i] = true
            provinces ++ //We are in a new province
            queue := []int{i} //Explore all cities in this province.
            for j := 0; j < len(queue); j++ {
                row := isConnected[queue[j]] //row of isConnected matrix
                for k,v := range(row){ 
                    if v == 1 && !seen[k]{ //Visit city k next
                        seen[k] = true
                        queue = append(queue,k)
                    }
                }
            }
        }
    }
    return provinces
}
