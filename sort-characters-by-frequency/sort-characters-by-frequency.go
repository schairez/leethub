import "strings"

func max(a, b int) int { if a >= b { return a}; return b}

func frequencySort(s string) string {
    charFreqMap := make(map[byte]int)
    maxFreq := 0
    for i :=0; i < len(s); i++ {
        if _, ok := charFreqMap[s[i]]; !ok {
            charFreqMap[s[i]] = 0
        }
        charFreqMap[s[i]]++
        maxFreq = max(maxFreq, charFreqMap[s[i]])
    }
    invFreqChars := make(map[int][]byte)
    for charKey, charFreqVal := range charFreqMap {
        if _, ok := invFreqChars[charFreqVal]; !ok {
            invFreqChars[charFreqVal] = make([]byte, 0)
        }
        invFreqChars[charFreqVal] = append(invFreqChars[charFreqVal], charKey)
    }
    var sb strings.Builder
    sb.Grow(len(s))
    fmt.Println(invFreqChars)
    for currFreq := maxFreq; currFreq > 0; currFreq-- {
        fmt.Println(currFreq)
        if chars, ok := invFreqChars[currFreq]; ok {
            iter := 0
            for iter < len(chars) {
                for i := 0; i < currFreq; i++ {
                    sb.WriteByte(chars[iter])
                }
                iter++
            }
        }
    }
    
    return sb.String()
    
    
}