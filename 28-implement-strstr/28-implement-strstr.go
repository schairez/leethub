// 28. Implement strStr()
// time: O(n)
// space: O(1)
func strStr(haystack string, needle string) int {
    k := len(needle)
    if k == 0 {
        return 0
    }
    n := len(haystack)
    start := 0
    for start < n {
        pIdx := 0
        sIdx := start
        for pIdx < k && sIdx < n && haystack[sIdx] == needle[pIdx] {
            sIdx++
            pIdx++
        } 
        if pIdx == k {
            return start
        }
        start++
    }
    
    return -1
}