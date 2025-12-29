func pyramidTransition(bottom string, allowed []string) bool {
    //preprocess allowed
    mem := make(map[string][]byte)
    for _,word := range allowed{
        key := word[0:2]
        _, ok := mem[key]
        if !ok {
            mem[key] = []byte{word[2]}
        } else {
            mem[key] = append(mem[key],word[2])
        }
    }
    
    //DFS: Build from left to right
    qmem := make(map[string]bool)
    qmem[string(bottom[0])] = true
    queue := []string{string(bottom[0])}
    //mem := make(map[string]bool)
    for i := 0; i < len(queue); i++ {
        leaf := queue[i]
        n := len(queue[i])
        target := n + 1
        kids := []string{string(bottom[n])}
        for j := 0; j < len(kids); j++ {
            kid := kids[j]
            k := len(kid) - 1
            //left := leaf[k]
            //right := kid[k]
            base := string(leaf[k]) + string(kid[k])
            for _,char := range mem[base] {
                next := kid + string(char)
                    if len(next) < target {
                        kids = append(kids,next)
                    } else {
                        if target == len(bottom){
                            return true
                        } else if !qmem[next]{
                            queue = append(queue,next)
                            qmem[next] = true
                        }
                    }
                }
            }
        }
    return false
}
