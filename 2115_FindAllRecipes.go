func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    /*
    IDEA: Directed graph of a -> b means that a is an ingredient of b.
    Make supplies into roots of directed graph.
    Generate iteratively from roots.
    Set up a queue of things that can be made.

    REMARK: supplies array is the queue!
    REMARK: This is actually a really weird Kahn's Algorithm Topsort problem.
    */
    fout := make([]string,0)
    need := make(map[string]int) //TODO: optimize this
    for i := 0 ; i < len(recipes); i++ {
        recipe := recipes[i]
        for j := 0; j < len(ingredients[i]); j++ {
            need[recipe]++
        }
    }
    for i := 0; i < len(supplies); i++{
        //item := supplies[i]
        for j, items := range(ingredients){
            for _, item := range(items){
                if item == supplies[i]{
                    need[recipes[j]]--
                    if need[recipes[j]] == 0 {
                        supplies = append(supplies,recipes[j])
                    }
                }
            }
        }
        _,isRecipe := need[supplies[i]] //Check if this is a recipe
        if isRecipe {
            fout = append(fout,supplies[i])
        }
    }
    return fout
}
