/*
convert knowledge into a trie for efficient memory retrieval
*/

type MyTrie struct {
	Children []*MyTrie
	Val      string
}

func newMyTrie() *MyTrie { //initializer
	kids := make([]*MyTrie, 26)
	return &MyTrie{kids, ""}
}

func (node *MyTrie) add(key, val string) {
	for _, char := range key {
		i := char - 'a'
		if node.Children[i] == nil {
			node.Children[i] = newMyTrie()
		}
		node = node.Children[i]
	}
	node.Val = val
}

func (node *MyTrie) find(key []byte) string {
	for _, char := range key {
		i := int(char - 'a')
		if node.Children[i] == nil {
			return "?"
		}
		node = node.Children[i]
	}
	return node.Val
}

func evaluate(s string, knowledge [][]string) string {
	mem := newMyTrie()
	for _, keyval := range knowledge {
		mem.add(keyval[0], keyval[1])
	}
	var readkey bool //default false
	key := make([]byte, 0)
    var builder strings.Builder
	n := len(s)
	for i := 0; i < n; i++ {
		char := s[i]
        //fmt.Println(i,char,readkey)
		if readkey {
			if char != ')' {
				key = append(key, char)
			} else {
				val := mem.find(key)
                builder.WriteString(val)
                //fmt.Println(string(key),val)
				if len(val) == 0 {
                    builder.WriteString("?")
				}
				key = make([]byte, 0)
				readkey = false
			}
		} else {
			if char != '(' {
                builder.WriteString(string(char))
			} else {
				readkey = true
				key = make([]byte, 0)
			}
		}
	}
	//fmt.Println(fout)
	return builder.String()
}
