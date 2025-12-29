/*
IDEA: 
We are basically constructing a recombining binary tree backwards.
Solve recursively!
*/

//import fmt

func pyramidTransition(bottom string, allowed []string) bool {
	fmt.Println("Inputs:", bottom, allowed)
	n := len(bottom) - 1 //length of next layer
    if n <= 0 {
        fmt.Println("CASE INPUT LETTER 0")
        return true
    }
	fmt.Println(n)

	//generate all possible permutations of next layer
	kids := make([]string, 0) //prelim
	adults := make([]string, 0)

	legs := bottom[0:2] // case where bottom has length <= 0 is already handled above.
	//birth kids
	for _, word := range allowed {
		if word[0:2] == legs {
			kids = append(kids, string(word[2]))
		}
	}
	fmt.Println("kids:", kids)

	if len(kids) == 0 {
        fmt.Println("CASE CANNOT BUILD NEXT GEN")
		return false
	}
	if n <= 1 {
		fmt.Println("CASE INPUT LETTER1")
		return true
	} else {
		//grow kids
		var age int
		for i := 0; i < len(kids); i++ {
			kid := kids[i]
			age = len(kid)
			legs = bottom[age : age+2]
			for _, word := range allowed {
				//fmt.Println(word[0:2])
				if word[0:2] == legs {
					//fmt.Println("DEBUG INDICATOR")
					if age == n-1 {
						adults = append(adults, kid+string(word[2]))
					} else {
						kids = append(kids, kid+string(word[2]))
					}
				}
			}
			fmt.Println("age:", age, "kids:", kids)
			fmt.Println("i:", i, "kids:", kids)
		}
	}
	fmt.Println("adults:", adults)
	var fout bool
	for _, word := range adults {
		if pyramidTransition(word, allowed) {
			fout = true
			break
		}
	}
	//fmt.Println(fout)
	return fout
}
