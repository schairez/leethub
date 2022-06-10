func max(a, b int) int {if a >= b {return a}; return b}

func largestVariance(s string) int {
    n := len(s)
    var freq [26]int
    for i := range s {
        freq[s[i]-'a']++
    }
    f := func(ch1, ch2 byte, isRev bool) int {
        cnt1, cnt2, i, variance := 0, 0, 0, 0
        if isRev { 
            i = n-1 
        }
        for i >= 0 && i < n {
            if s[i] == ch1 {
                cnt1++
            } else if s[i] == ch2 {
                cnt2++
            }
            if cnt1 < cnt2 {
                cnt1, cnt2 = 0, 0
            }
            if cnt1 != 0 && cnt2 != 0 {
                variance = max(variance, cnt1 - cnt2)
            }
            if isRev {
                i--
            } else {
                i++
            }
        }
        return variance
    }
    res := 0
    for ch1 := byte('a'); ch1 <= byte('z'); ch1++ {
        for ch2 := byte('a'); ch2 <= byte('z'); ch2++ {
            if ch1 == ch2 || freq[ch1-'a'] == 0 || freq[ch2-'a'] == 0 {
                continue
            }
            res = max(res, f(ch1, ch2, true))
            res = max(res, f(ch1, ch2, false))
        }
    }
    return res
}
