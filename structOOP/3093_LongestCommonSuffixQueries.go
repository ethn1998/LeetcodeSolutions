/*
suffix : prefix read backwards
store in trie. each node is a valid suffix
also, to store wordsContainer index in trie value for lookup.
*/

type MyTrie struct {
    Children []*MyTrie
    Val int
}

func newMyTrie() *MyTrie { //initializer
    kids := make([]*MyTrie,26)
    return &MyTrie{kids,-1}
}

func (this *MyTrie) Insert(word string, val int) { //NOTE: Insert characters backwards
    curr := this
    n := len(word)
    for i := n-1; i >= 0; i-- {
        j := int(word[i] - 'a') //coerce int
        if curr.Children[j] == nil {
            curr.Children[j] = newMyTrie()
        }
        curr = curr.Children[j]
    }
    curr.Val = val
}

func (this *MyTrie) SlowFetch(query string) int { //bfs search of dictionary trie
    fout := -1
    //phase 1: follow query (dfs)
    n := len(query)
    for i := n-1; i >= 0; i-- {
        j := int(query[i] - 'a')
        if this.Children[j] != nil {
            this = this.Children[j]
        } else {
            break
        }
    }

    //phase 2: bfs
    var found bool
    curr := []*MyTrie{this}
    for true {
        next := []*MyTrie{}
        for _,node := range curr {
            for i := range node.Children {
                if node.Children[i] != nil {
                    next = append(next,node.Children[i])
                }
            }
            if node.Val != -1 {
                found = true
            }
        }
        if found || len(next) == 0 {
            break
        }
        //else
        curr = next
    }
    if !found {
        return -1
    }

    //phase 3: find min index
    for _,node := range curr {
        if node.Val != -1 {
            if fout > -1 {
                if fout > node.Val {
                    fout = node.Val
                }
            } else {
                fout = node.Val
            }
        }
    }
    return fout
}

func (this *MyTrie) FastFetch(query string) int { //naive search from query memory
    n := len(query)
    i := n-1
    for i >= 0 {
        j := int(query[i] - 'a')
        if this.Children[j] != nil {
            this = this.Children[j]
        } else {
            break
        }
        i--
    }
    if i < 0 { //query exists in memory
        return this.Val
    } else {
        return -1
    }
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
    //fmt.Println("InSize:",len(wordsContainer),len(wordsQuery)) //DEBUG
    //process wordsContainer
    root := newMyTrie()
    m := len(wordsContainer)
    i0 := 0 //default output
    for i := m-1; i >= 0; i-- {
        word := wordsContainer[i]
        //fmt.Println("Insert:",i,len(word)) //DEBUG
        root.Insert(word,i)
        if len(word) <= len(wordsContainer[i0]) {
            i0 = i
        }
    }

    //process wordsQuery
    n := len(wordsQuery)
    memo := newMyTrie() //to lookup previously seen queries
    fout := make([]int,n)
    for i,word := range wordsQuery {
        val := memo.FastFetch(word)
        if val < 0 { //if not in memo
            val = root.SlowFetch(word)
            if val == -1 {
                val = i0
            }
            memo.Insert(word,val) //for quick lookup
        }
        fout[i] = val
    }

    return fout
}
