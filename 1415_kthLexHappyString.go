import "fmt"

func getHappyString(n int, k int) string {
    //IDEA: Greedily generate happy strings. Should not need lexicographic sorting.
    mystrings := []string{"a","b","c"}
    nhappystrings := 3
    for l := 1; l < n; l++ { //Generate happy strings of length l+1
        mystrings1 := make([]string,0)
        for _,s := range(mystrings){
            char := string(s[len(s)-1])
            if char != "a" {
                mystrings1 = append(mystrings1,fmt.Sprintf("%sa",s))
            } 
            if char != "b" {
                mystrings1 = append(mystrings1,fmt.Sprintf("%sb",s))
            }
            if char != "c"{
                mystrings1 = append(mystrings1,fmt.Sprintf("%sc",s))
            }
        }
        mystrings = make([]string,len(mystrings1))
        nhappystrings = copy(mystrings,mystrings1)
    }
    if k <= nhappystrings {
        return mystrings[k-1]
    }
    return ""
}
