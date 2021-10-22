//time: O(n)
//space: O(n); maintain two hashmaps

import "strings"

func wordPattern(pattern string, s string) bool {
    strArr := strings.Fields(s)
    bytesArr := []byte(pattern)
    strMap := map[string]int{}
    bytesMap := map[byte]int{}
    nStr, nBytes := len(strArr), len(bytesArr)
    if nStr != nBytes { return false }
    for i := 0; i < nStr; i++ {
        letter, word := bytesArr[i], strArr[i]
        if _, ok := strMap[word]; !ok {
            strMap[word] = i
        }
        if _, ok := bytesMap[letter]; !ok {
            bytesMap[letter] = i
        }
        if strMap[word] != bytesMap[letter] { return false }
    }
    return true 
}