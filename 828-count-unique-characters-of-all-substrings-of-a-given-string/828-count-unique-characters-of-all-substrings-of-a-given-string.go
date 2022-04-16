
// condition: 1 <= s.length <= 105
func uniqueLetterString(s string) int {
    n := len(s)
    if n == 1 {
        return 1
    }
    var (
        charIdxLoc [26]int  //
        prevCharCnt [26]int //
        cntUnique int       // cnt of unique chars among substrings ending at idx
        total int           // total unique chars among all substrings
    )
    for i := 0; i < 26; i++ {
        charIdxLoc[i] = -1
    }
    for i := 0; i < n; i++ {
        idx := int(s[i] - 'A')
        maxUniqueAtIdx := i - charIdxLoc[idx]
        cntUnique = cntUnique - prevCharCnt[idx] + maxUniqueAtIdx
        charIdxLoc[idx] = i
        prevCharCnt[idx] = maxUniqueAtIdx
        total += cntUnique
    }
    
    return total
}