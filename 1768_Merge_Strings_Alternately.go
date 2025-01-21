package main

import(
    "fmt"
    "strings"
)

func mergeAlternately(word1 string, word2 string) string {
    var i,j int //Track position in each word
    var char,fout string

    if word1 == ""{
        return word2
    } else if word2 == "" {
        return word1
    } else {
        sword1 := strings.Split(word1,"") //Convert string into array
        sword2 := strings.Split(word2,"")

        for t := 0; t < len(word1)+len(word2); t++ {
            if t%2 == 0{
                char = sword1[i]
                fout = fmt.Sprintf("%s%s",fout,char)
                i++
                if i >= len(word1){
                    fout = fmt.Sprintf("%s%s",fout,word2[j:]) //Rest of second word
                    break
                }
            } else {
                char = sword2[j]
                fout = fmt.Sprintf("%s%s",fout,char)
                j++
                if j >= len(word2){
                    fout = fmt.Sprintf("%s%s",fout,word1[i:])//Rest of first word
                    break
                }
            }
        }
    return fout
    }
}
