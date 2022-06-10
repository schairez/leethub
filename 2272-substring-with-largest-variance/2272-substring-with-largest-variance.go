func max(a, b int) int { if a >= b { return a}; return b}

func largestVariance(s string) int {
    res := 0
    n := len(s)
    var freq [26]int
    for i := range s {
        freq[s[i]-'a']++
    }
    for char1 := 'a'; char1 <= 'z'; char1++ {
        for char2 := 'a'; char2 <= 'z'; char2++ {
            if char1 == char2 || freq[char1-'a'] == 0 || freq[char2-'a'] == 0 {
                continue
            }
            diff, diffWchar2 := 0, -n
            for _, char := range s {
                if char == char1 {
                    diff++
                    diffWchar2++
                } else if char == char2 {
                    diff--
                    diffWchar2 = diff 
                    diff = max(diff, 0)
                }
                res = max(res, diffWchar2)
            }
        }
    }
    return res
    
}