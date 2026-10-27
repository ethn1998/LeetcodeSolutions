func createGrid(m int, n int) []string {
    fout := make([]string,m)
    fout[0] = strings.Repeat(".",n)
    for i := 1; i < m; i++ {
        fout[i] = fmt.Sprintf("%s.",strings.Repeat("#",n-1))
    }
    return fout
}
