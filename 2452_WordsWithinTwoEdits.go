func twoEditWords(queries []string, dictionary []string) []string {
    fout := make([]string,0)
    /* maybe needed for binary search if more efficient soln needed
    slices.SortFunc(dictionary, func(i,j int) int {
        return strings.Compare(dictionary[i],dictionary[j])
    })
    */
    for _,query := range queries {
        for _,word := range dictionary {
            var diff int
            for i := range query {
                if query[i] != word[i] {
                    diff++
                }
                if diff > 2 {
                    break
                }
            }
            if diff <= 2 {
                fout = append(fout,query)
                break
            } 
        }
    }
    return fout
}
