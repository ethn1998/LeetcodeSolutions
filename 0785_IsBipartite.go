func isBipartite(graph [][]int) bool {
    color := make([]int,len(graph)) //Color nodes with \pm 1. 
    for src := range(graph){
        if color[src] == 0 { //0 means node is colorless
            //We are in a new connected subgraph
            color[src] = 1 //Fix
            queue := []int{src}
            for i := 0; i < len(queue); i++{
                node := queue[i]
                edges := graph[node]
                for _, dst := range(edges) {
                    if color[dst] == 0 { //If node has not been colored
                        color[dst] = -color[node]
                        queue = append(queue,dst)
                    } else {
                        if color[dst] == color[node]{
                            return false
                        }
                    }
                }
            }
        }
    }
    return true
}
