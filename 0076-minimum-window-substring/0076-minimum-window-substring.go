
// time: O(lenS + lenT)
// space: O(lenS + lenT)
func minWindow(s string, t string) string {
    lenS, lenT := len(s), len(t)
    if lenT > lenS {
        return ""
    }
    var (
        freqT [128]int
        freqWin [128]int
    )
    for i := 0; i < lenT; i++ {
        freqT[t[i]]++
    }
    var res string
    maxInt32 := 1 << 31 -1 
    lenRes := maxInt32
    remaining := lenT
    left, right := 0, 0
    for right < lenS {
        freqWin[s[right]]++
        if freqT[s[right]] != 0 && freqWin[s[right]] <= freqT[s[right]] {
            remaining--
        }
        //contract the win
        if remaining == 0 {
            for left < lenS && freqWin[s[left]] > freqT[s[left]] {
                freqWin[s[left]]--
                left++
            }
            if (right - left + 1) < lenRes {
                res = string(s[left:right+1])
                lenRes = len(res)
            }
        }
        right++
    }
    if lenRes == maxInt32 {
        return ""
    }
    return res
}

