func canConstruct(ransomNote string, magazine string) bool {
    cuts := make(map[rune]int)
    for _, char := range magazine{
        cuts[char] ++
    }

    for _, char := range ransomNote{
        if cuts[char] == 0{
            return false
        }
        cuts[char] --
    }
    return true
}
