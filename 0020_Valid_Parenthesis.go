func isValid(s string) bool {
    
    opens := make(map[string]bool)
    opens["("] = true
    opens["{"] = true
    opens["["] = true
    opens[")"] = false
    opens["}"] = false
    opens["]"] = false

    closes := make(map[string]string)
    closes[")"] = "("
    closes["}"] = "{"
    closes["]"] = "["

    chars := strings.Split(s,"")
    test := make([]string,0,len(s))

    for _, char := range(chars) {
        if (opens[char]){
            test = append(test, char)
        }else{
            if (len(test) == 0){
                return false
            } else if (test[len(test)-1] != closes[char]){
                return false
            } else {
                test = test[:len(test)-1]
            }
        }
    }
    return (len(test) == 0)
}
