type Trie struct {
	child [26]*Trie //array of capacity 26, size of alphabet
	id    int
}

func newTrie() *Trie {
	return &Trie{id: -1}
}

func add(node *Trie, word string, idx *int) int { //inserts new word into trie
	for _, char := range word {
		i := char - 'a'
		if node.child[i] == nil {
			node.child[i] = newTrie()
		}
		node = node.child[i]
	}
	if node.id == -1 { //set unique id (key) for each word in trie
		*idx++
		node.id = *idx
	}
	return node.id //return unique id of new word
}

/*
func update(x *int64, y int64) { //DO WE NEED THIS?!
    if x == -1 || y < *x {
        *x = y
    }
}
*/

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	//fmt.Println(len(source),len(cost))//DEBUG
	//var src, trans, dst string
	var isrc, curr, itrans, idst int
	var srctotrans, transtodst int64

	n := len(source)   //length of string
	m := len(original) //2*m = max number of nodes in graph

	root := newTrie() //Trie for encrypting substrings
	id := -1          //init counter for id

	//costmat := make(map[string]map[string]int64, len(original))
	costmat := make(map[int]map[int]int64, len(original))
	for _, src := range original {
		isrc = add(root, src, &id)
		_, ok := costmat[isrc]
		if !ok {
			costmat[isrc] = make(map[int]int64)
			costmat[isrc][isrc] = 0 //redundant?
			queue := []int{isrc}
			for i := 0; i < len(queue); i++ {
				curr = queue[i]
				srctotrans = costmat[isrc][curr]
				for j := 0; j < m; j++ {
					itrans = add(root, original[j], &id)
					if itrans == curr {
						idst = add(root, changed[j], &id)
						transtodst = int64(cost[j])
						_, ok := costmat[isrc][idst]
						if ok {
							if costmat[isrc][idst] > srctotrans + transtodst {
								costmat[isrc][idst] = srctotrans + transtodst
								queue = append(queue, idst)
							}
						} else {
							costmat[isrc][idst] = srctotrans + transtodst
							queue = append(queue, idst)
						}
					}
				}
			}
		}
	}

    dp := make([]int64,n)
    for i := range dp {
        dp[i] = -1
    }

    for left := 0; left < n; left++ {
        if left > 0 { //to avoid indexing errors, and only trigger after we init left = 0 case
            if dp[left - 1] == -1 {
                continue
            }
        }
        var prev int64
        if left == 0 {
            prev = 0
        } else {
            prev = dp[left-1]
        }
        if source[left] == target[left] {
            if dp[left] > -1 {
                if dp[left] > prev {
                    dp[left] = prev
                }
            } else {
                dp[left] = prev
            }
        }
        u := root
        v := root
        for right := left; right < n; right++ { //try to convert string[left:right]
            u = u.child[source[right]-'a']
            v = v.child[target[right]-'a']
            if u == nil || v == nil {
                break
            }
            if u.id != -1 && v.id != -1 {
                val, ok := costmat[u.id][v.id]
                if ok {
                    if dp[right] > -1 {
                        if dp[right] > prev + val {
                            dp[right] = prev + val
                        }
                    } else {
                        dp[right] = prev + val
                    }
                }
            }
        }
    }
	//fmt.Println(dp) //DEBUG
	return dp[len(source)-1]
}
