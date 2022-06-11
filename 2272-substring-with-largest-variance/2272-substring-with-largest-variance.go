func max(a, b int) int {if a >= b {return a}; return b}
func largestVariance(s string) int {
    n := len(s)
    var freq [26]int 
    alphaSet := make(map[byte]struct{}, 26)
    // is char freq greater than 1
    isCharFreq := false
    for i := range s {
        key := int(s[i] - 'a')
        freq[key]++
        alphaSet[s[i]] = struct{}{}
        if !isCharFreq && freq[key] > 1 {
            isCharFreq = true
        }
    }
    res := 0
    if !(isCharFreq && len(alphaSet) > 1) {
        return res
    }
    for char1 := range alphaSet {
        for char2 := range alphaSet {
            if char1 == char2 {
                continue
            } 
            cnt1, cnt2, rem2 := 0, 0, freq[char2-'a']
            for i := 0; i < n; i++ {
                curr := s[i]
                if curr == char1 {
                    cnt1++
                } else if curr == char2 {
                    cnt2++
                    rem2--
                }
                if cnt1 < cnt2 && rem2 > 0 {
                    cnt1, cnt2 = 0, 0
                }
                if cnt1 != 0 && cnt2 != 0 {
                    res = max(res, cnt1 - cnt2)
                }
            }
        }
    }
    return res
}