import "strings"

func wordPattern(pattern string, s string) bool {
    var char byte
    var token string
    tokens := strings.Split(s," ")
    if len(pattern) != len(tokens){
        return false
    }

    pat_to_token := make(map[byte]string) //Forward
    token_to_pat := make(map[string]byte) //Backward

    for i:=0; i<len(pattern); i++{
        char = pattern[i]
        token = tokens[i]
        val, ok := pat_to_token[char]
        if ok{
            if val != token{
                return false
            }
        } else{
            pat_to_token[char] = token
        }
        val1, ok := token_to_pat[token]
        if ok{
            if val1 != char{
                return false
            }
        } else{
            token_to_pat[token] = char
        }        
    }
    return true
}
