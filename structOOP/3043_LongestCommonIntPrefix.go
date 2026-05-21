type TrieNode struct {
    Children [10]*TrieNode
    Val int
}

func newTrieNode() *TrieNode {
    return &TrieNode{[10]*TrieNode{nil,nil,nil,nil,nil,nil,nil,nil,nil,nil},0}
}

func (node *TrieNode) Insert(n string) { //insert new value into trie
    for i,d := range n {
        j := d - '0'
        if node.Children[j] == nil {
            node.Children[j] = newTrieNode()
        }
        node = node.Children[j]
        node.Val = i + 1
    }
}

func (node *TrieNode) Fetch(n string) int { //get longest common prefix
    var fout int
    for _,d := range n {
        j := d - '0'
        if node.Children[j] != nil {
            node = node.Children[j]
            fout++
        } else {
            break
        }
    }
    return fout
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
    var fout int
    root := newTrieNode()
    for _,n := range arr1 {
        root.Insert(strconv.Itoa(n))
    }
    for _,n := range arr2 {
        x := root.Fetch(strconv.Itoa(n))
        if fout < x {
            fout = x
        }
    }
    return fout
}
