/*
An integer n can only have four divisors if it satisfies either of the two criteria:
(1) n is the product of two distinct prime numbers.
(2) n is a perfect cube
*/

func sumFourDivisors(nums []int) int {
	var fout int
	//erato
	omega := 100000
	isPrime := make([]bool, omega+1)
	primes := make([]int, 0)
	for i := 2; i < omega; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= omega; p++ {
		if isPrime[p] {
			primes = append(primes, p)
			for k := p * p; k < omega; k += p {
				isPrime[k] = false
			}
		}
	}
	//fmt.Println(primes)
	mem := make(map[int]int) //cache previously seen numbers.
    var incr int
	for _, n := range nums {
		v, ok := mem[n]
		if ok {
			fout += v
		} else {
            incr = 0
			if !isPrime[n] {
				for _, p := range primes {
					if n%p == 0 {
						q := n / p
						if p != q {
							if isPrime[q] || q == p*p {
								//fmt.Printf("%d * %d = %d\n",p,q,n)
								incr = 1 + p + q + n
								break
							}
						}
					}
				}
			}
            fout += incr
            mem[n] = incr
		}
	}
	return fout
}
