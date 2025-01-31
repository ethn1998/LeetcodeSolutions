func minMutation(startGene string, endGene string, bank []string) int {
    //Note: Formally, we should use DSU for this.
    seen := make(map[string]int) //Count number of mutations needed to get to gene
    seen[startGene] = 0
    geneQueue := []string{startGene}
    for i := 0; i < len(geneQueue); i++ {
        gene0 := geneQueue[i]
        for _, gene := range(bank){
            _, ok := seen[gene]
            if !ok {
                nmut := 0
                for k := range(gene){
                    if gene[k] != gene0[k] {
                        nmut++
                    }
                }
                if nmut == 1 {
                    if gene == endGene{
                        return seen[gene0] + 1
                    } else {
                        seen[gene] = seen[gene0] + 1
                        geneQueue = append(geneQueue,gene)
                    }
                }
            }

        }
    }
    d, ok := seen[endGene] //Probably redundant
    if ok {
        return d
    }
    return -1 //Not in subgraph
}
