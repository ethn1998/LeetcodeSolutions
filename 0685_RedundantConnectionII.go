func findRedundantDirectedConnection(edges [][]int) []int {
    /*
    NOTES:
    (1) The root node will have no incoming edges.
    (2) Every node except for the root node must have exactly one parent by defn.
    (3) There is only one redundant connection added by problem specification.
    Idea: Count number of incoming edges.
    Cases:
    (Example 1) 
    Existence of node with 2 incoming edges: 
    Remove one of the incoming edges. (by (3) this should work)
    (Example 2) 
    All nodes have exactly one incoming edge: Use Kahn's Algorithm to trim leaves.
    Choose final edge of cycle.
    """
    */

    //STEP 1: Count number of incoming edges using a map
    /*
    NOTE: 
    Number of nodes in a (singularly) rooted tree is number of edges + 1
    But because we have one extra edge added, the number of nodes is len(edges)!
    */
    mark := -1
    outgoingEdges := make([]int,len(edges)) //There are len(edges) nodes.
    incomingEdges := make([]int,len(edges))
    for _,edge := range(edges) {
        src := edge[0] - 1 //For array indexing
        dst := edge[1] - 1
        outgoingEdges[src] += 1
        incomingEdges[dst] += 1
        if incomingEdges[dst] == 2 {
            mark = dst //Every node must have exactly one parent
        }
    }
    if mark != -1 {
    // Step 2: Use (reversed) Kahn's Algorithm to detect cycles
        rootQueue := make([]int,0)
        notCycleEdges := make(map[int]bool) //Store indices of removed edges
        for i, v := range(incomingEdges) {
            if v == 0 {
                rootQueue = append(rootQueue,i) //
            }
        }
        for j := 0; j < len(rootQueue); j++ {
            root := rootQueue[j]
            for k,edge := range(edges){
                src := edge[0] -1
                dst := edge[1] -1
                if src == root {
                    notCycleEdges[k] = true
                    outgoingEdges[src] -= 1
                    incomingEdges[dst] -= 1
                    if incomingEdges[dst] == 0 {
                        rootQueue = append(rootQueue,dst) //dst becomes new root
                    }
                }
            }
        }
        cycle := (len(notCycleEdges) != len(edges)) //boolean to see if cycles are present
        for i := len(edges) - 1 ; i >= 0 ; i-- {
            edge := edges[i]
            //src := edge[0]
            dst := edge[1] -1
            if dst == mark {
                if cycle {
                    if !notCycleEdges[i] { //Must choose something in cycle
                        return edge
                    }
                } else {
                    return edge
                }
            }
        }
    } else {
        // Step 2: Use Kahn's Algorithm to remove leaves
        leafQueue := make([]int,0)
        notCycleEdges := make(map[int]bool) //Store indices of removed edges
        for i, v := range(outgoingEdges) {
            if v == 0 {
                leafQueue = append(leafQueue,i) //Node i is a leaf node
            }
        }
        for j := 0; j < len(leafQueue); j++ {
            leaf := leafQueue[j]
            for k,edge := range(edges){
                src := edge[0] -1
                dst := edge[1] -1
                if dst == leaf { //trim leaf
                    notCycleEdges[k] = true
                    outgoingEdges[src] -= 1
                    incomingEdges[dst] -= 1
                    if outgoingEdges[src] == 0 {
                        leafQueue = append(leafQueue,src) //src becomes new leaf
                    }
                }
            }
        }
        for i := len(edges) - 1; i >= 0 ; i -- {
            if !notCycleEdges[i]{ //Choose anything from the cycle
                return edges[i]
            }
        }
    }
    return make([]int,0) //Redundancy measure. If this is returned, it means there's an error somewhere.
}
