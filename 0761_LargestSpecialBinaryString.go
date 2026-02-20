/*
IDEA:
BFS to simulate all possible swaps.
op: l0l1r1r0 -> l0r1l1r0: where l1, r1 are special strings.

Identify all consecutive special substring pairs, simulate swaps as candidates
Trie to cache previously found strings. [For this question, this is isomorphic to a binary tree.]
Use the trie to get lexicographically greatest possible string:
If Right != nil, go right and append "1"
else if Left != nil, go left and append "0"
else terminate (when depth of tree = length of string)

type Node struct {
    Val bool //true : visited
    Zero *Node //0
    One *Node //1
}

func newNode() *Node {
    return &Node{false,nil,nil}
}

func write(s string, root *Node) *Node {
    n := len(s)
    node := root //point to root
    for i, char := range s {
        if i < n - 1 {
            if char == '0'{
                if node.Zero == nil {
                    node.Zero = newNode()
                }
                node = node.Zero
            } else {
                if node.One == nil {
                    node.One = newNode()
                }
                node = node.One
            }
        } else {
            node.Val = true
        }
    }
    return root
}

func read(s string, root *Node) bool {
    //fmt.Println("read in:",s,len(s),root) //DEBUG
    node := root
    for _, char := range s {
        //fmt.Println(i, char, node) //DEBUG
        if char == '0' {
            if node.Zero != nil {
                node = node.Zero
            } else {
                break
            }
        } else {
            if node.One != nil {
                node = node.One
            } else {
                break
            }
        }
    }
    return node.Val
}

func makeLargestSpecial(s string) string {
    n := len(s)
    fmt.Println(n)
    queue := []string{s}
    mem := newNode()
    mem = write(s, mem)

    var l0, l1, r1, r0 string
    for i := 0; i < len(queue); i++ {
        parent := queue[i]
        for l := 0; l < n; l++ {
            if parent[l] == '1' {
                l0 = parent[0:l]
                for j := l+2; j < n; j += 2 {
                    l1 = parent[l:j]
                    if isSpecial(l1) && parent[j] == '1' {
                        for r := j; r <= n; r += 2 {
                            r1 = parent[j:r]
                            if isSpecial(r1) {
                                r0 = parent[r:n]
                                child := fmt.Sprintf("%s%s%s%s",l0,r1,l1,r0)
                                seen := read(child,mem)
                                // DEBUG
                                fmt.Println("indices:",l,j,r)
                                fmt.Println("l0:",l0,"l1:",l1,"r1:",r1,"r0:",r0)
                                //
                                if !seen {
                                    //fmt.Println(seen,child) //DEBUG
                                    mem = write(child, mem)
                                    queue = append(queue,child)
                                }
                            }
                        }
                    }
                }
            }
        }
        //fmt.Println("queue:",queue) //DEBUG
        //fmt.Println(i,len(queue)) //DEBUG
    }
    fmt.Println("queue end",len(queue))
    var fout string
    node := mem
    for true {
        if node.One != nil { //prioritize
            fout = fmt.Sprintf("%s%d",fout,1)
            node = node.One
        } else {
            fout = fmt.Sprintf("%s%d",fout,0)
            if node.Zero != nil {
                node = node.Zero
            } else {
                break
            }
        }
    }
    return fout
}

func isSpecial(s string) bool {
    var sum int
    n := len(s)
    if n == 0 || n % 2 != 0 { //Property 1 implies that len(s) must be even
        return false
    }
    for i, char := range s {
        if char == '1' {
            sum++
        }
        if 2 * sum <= i { //less error prone than int divide
            return false
        }
    }
    return sum == n / 2 
}

func isLesser(s0, s1 string) { //s0 < s1
    n := len(s0)
    for i := 0; i < n; i++ {
        char0 := s0[i]
        char1 := s1[i]
        if char0 == "0" && char1 == "1" {
            return true
        } else if char0 == "1" && char1 == "0" {
            return false
        }
    }
    return false
}

//PROBLEM: TOO SLOW

//HINT: 
Case where S cannot be decomposed as a mountains, ignore first and last element, return 
"1%s0",where S is recursion of middle.

*/

func makeLargestSpecial(s string) string {
    //fmt.Println("In:",s) //DEBUG
    if len(s) == 0 {
        return s
    }
    var fout string
    //decomposition
    var left, sum int
    n := len(s)
    decomp := make([]string,0)
    for i, char := range s {
        sum += int(char - '0')
        if sum == i + 1 - sum {
            decomp = append(decomp,s[left:i+1])
            left = i + 1
        }
    }
    //fmt.Println("unprocessed:",decomp,len(decomp)) //DEBUG
    if len(decomp) > 1 {
        processed := make([]string,len(decomp))
        for i, word := range decomp {
            processed[i] = makeLargestSpecial(word)
        }
        //fmt.Println("processed:",processed) //DEBUG
        slices.Sort(processed)
        //fmt.Println("sorted",processed) //DEBUG
        for _,u := range processed {
            fout = fmt.Sprintf("%s%s",u,fout)
        }
    } else {
        fout = fmt.Sprintf("1%s0",makeLargestSpecial(s[1:n-1]))//recursion
    }
    return fout
}
