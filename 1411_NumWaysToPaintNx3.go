/*
This is not a programming question, but rather combinatorics
Discussion:
2 col row can generate 2 possible 3 col, 3 possible 2 col
3 col row can generate 2 possible 3 col, 2 possible 2 col

n2 = 3 * n2 + 2 * n3
n3 = 2 * n2 + 2 * n3
*/

func numOfWays(n int) int {
    var x1, y1 int
    x := 6 // #2 col
    y := 6 // #3 col
    //matrix exponentiation
    for i := 1; i < n; i++{
        x1 = (3*x + 2*y) % 1000000007
        y1 = 2*(x + y) % 1000000007
        x = x1
        y = y1
    }
    return (x + y) % 1000000007
}
