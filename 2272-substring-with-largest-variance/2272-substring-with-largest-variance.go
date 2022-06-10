func max(a, b int) int { if a >= b { return a}; return b}

func largestVariance(s string) int {
    res := 0
    n := len(s)
    var freq [26]int 
    for i := range s {
        freq[s[i]-'a']++
    }
    //fmt.Println(freq)
    keys := make([]byte, 0, 26)
    for i := range freq {
        if freq[i] != 0 {
            keys = append(keys, byte(i + 'a'))
        }
    }
    //fmt.Println(keys)
    for _, char1 := range keys {
        for _, char2 := range keys {
            if char1 == char2 {
                continue
            }
            diff, diffWchar2 := 0, -n
            for i := 0; i < n; i++ {
                char := s[i]
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