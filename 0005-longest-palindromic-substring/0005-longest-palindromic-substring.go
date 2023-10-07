// time: O(n^2)
// space: O(1)

func longestPalindrome(s string) string {
    n := len(s)
    start, end, maxLen := 0, 0, 0
    for i := 0; i < n; i++ {
        oddLen := lenLongestPal(i, i, s)
        if oddLen > maxLen {
            dist := oddLen / 2
            start, end = i - dist, i + dist
            maxLen = oddLen
        }
        evenLen := lenLongestPal(i, i+1, s)
        if evenLen > maxLen {
            dist := (evenLen / 2 ) - 1
            start, end = i - dist, i + 1 + dist
            maxLen = evenLen 
        }
    }
    return s[start:end+1]
}

func lenLongestPal(i int, j int, s string) int {
    n := len(s)
    for i >= 0 && j < n && s[i] == s[j] {
        i--
        j++
    }
    return j - i - 1
}