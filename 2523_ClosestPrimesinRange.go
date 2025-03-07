func closestPrimes(left int, right int) []int {
    if right <= 2 {
        return []int{-1,-1}
    }
    //Find primes in range (left, right)
    primes := make([]int,0) //array to store primes in left,right
    primality := make([]bool,right+1) //Boolean array if number is NOT prime.
    for n := 2; n <= right; n++ {
        if !primality[n] {
            if n >= left {
                primes = append(primes,n)
            }
            p := n
            for k := 2*p; k <= right; k += p {
                primality[k] = true                
            }
        }
    }
    if len(primes) < 2 {
        return []int{-1,-1}
    }
    var num1,num2 int
    diff := 1000000
    for j := 1; j <= len(primes)-1; j++ {
        p1 := primes[len(primes)-j-1]
        p2 := primes[len(primes)-j]
        if p1 < left {
            break
        } else {
            if p2 - p1 <= diff{
                num1 = p1
                num2 = p2
                diff = p2 - p1
            }
        }
    }
    return []int{num1,num2}
}
