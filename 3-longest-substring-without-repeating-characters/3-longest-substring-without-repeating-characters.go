
//time: O(n)
//space: O(min(m, n)); m = len(charSet)

// testcase: "abba"
// we move left to lastIdx+1
// this indicates first non-repeating start seq, but only
// if lastIdx >= left; otherwiser we'd be moving left to a back
// position; e.g. @ idx=3; map['a'] =0; if we do left = 0+1
// this is a position left of our last known left pos

func max(a, b int) int { if a >= b { return a}; return b}

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n == 0 || n == 1 {
        return n
    }
    longestLen := 0
    left := 0
    charMap := make(map[byte]int, 0)
    for right := 0; right < n; right++ {
        if lastIdx, ok := charMap[s[right]]; ok {
            if lastIdx >= left {
            left = lastIdx + 1
            }
        }
        longestLen = max(longestLen, right - left + 1)
        charMap[s[right]] = right 
    }
    return longestLen
}


















//time: O(n)
//space: O(min(m, n)); m = len(charSet)


//s consists of English letters, digits, symbols and spaces.
// can fit in byte 128bit charSet


/*
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
*/