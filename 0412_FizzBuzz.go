func fizzBuzz(n int) []string {
    var f,b,s string
    answer := make([]string,0)
    for i := 1; i<=n; i++ {
        if i % 3 == 0 {
            f = "Fizz"
        }else{
            f = ""
        }
        if i % 5 == 0 {
            b = "Buzz"
        } else {
            b = ""
        }
        s = fmt.Sprintf("%s%s",f,b)
        if len(s) == 0 {
            s = fmt.Sprintf("%d",i)
        }
        answer = append(answer,s)
    }
    return answer    
}
