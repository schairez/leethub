
//time: O(n)
//space: O(min(m, n)); m = len(charSet)


//s consists of English letters, digits, symbols and spaces.
// can fit in byte 128bit charSet



func lengthOfLongestSubstring(s string) int {
    //precondition
    //0 <= s.length <= 5 * 104
    n := len(s)
    if n == 0 { return 0 }
    
    charIdxMap := make(map[byte]int)
    res := 0
    leftPtr := 0
    for i := range s {
        if chIdx, ok := charIdxMap[s[i]]; ok {
            if chIdx >= leftPtr {
                leftPtr = chIdx + 1
            }
        }
        res = max(res, i - leftPtr + 1)
        charIdxMap[s[i]] = i
    }
    return res
}
    
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}