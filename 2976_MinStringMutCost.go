/*
26 * 26 min cost matrix generated via Dijkstra's
Query from cost matrix elementwise.
*/

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var fout int64
	costmat := make(map[byte]map[byte]int64, 26) //preprocess original, changed and cost
	var src, trans, dst byte
	var srctotrans, transtodst int64

	for _, src = range original {
		_, ok := costmat[src]
		if !ok {
			costmat[src] = make(map[byte]int64, 26)
			costmat[src][src] = 0
			queue := []byte{src}
			for i := 0; i < len(queue); i++ {
				trans = queue[i]
				srctotrans = costmat[src][trans]
				for j := 0; j < len(original); j++ {
					if trans == original[j] {
						dst = changed[j]
						transtodst = int64(cost[j])
						val, ok := costmat[src][dst]
						if !ok {
							costmat[src][dst] = srctotrans + transtodst
							queue = append(queue, dst)
						} else {
							if val > srctotrans+transtodst {
								costmat[src][dst] = srctotrans + transtodst
								queue = append(queue, dst)
							}
						}
					}
				}
			}
		}
	}
	//query answer elementwise
	for i := 0; i < len(source); i++ {
		src = byte(source[i])
		dst = byte(target[i])
		if src != dst {
			dsts, ok := costmat[src]
			if !ok {
				return -1
			} else {
				val, ok := dsts[dst]
				if ok {
					fout += val
				} else {
					return -1
				}
			}
		}
	}
	return fout
}
