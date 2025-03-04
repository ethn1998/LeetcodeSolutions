func checkPowersOfThree(n int) bool {
    //I think I learnt a trick to solve this one before in Number Theory class but forgot.
    //IDEA: Try to convert n into base 3. If a 2 exists, then return false
    for n > 0 {
        if n % 3 == 2{
            return false
        }
        n /= 3
    }
    return true
}
