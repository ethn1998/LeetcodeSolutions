import "fmt"

func getHint(secret string, guess string) string {
    var bulls, cows int
    counts := make(map[rune]int)
    for _, schar := range(secret){
        counts[schar]++
    }
    for i, gchar := range(guess){ //Find all bulls first
        if rune(secret[i]) == gchar{
            bulls++
            counts[gchar]--
        }
    }
    for i, gchar := range(guess){ //Then find all cows
            if rune(secret[i]) != gchar && counts[gchar] > 0 {
            cows++
            counts[gchar]--
        }
    }

    return fmt.Sprintf("%dA%dB",bulls,cows)   
}
