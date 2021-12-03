//time: O(n)
//space: O(n)

import "strings"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    //map src Index to replacement string 
    replaceIdxMap := make(map[int]int)
    //maps indices to search with idx of src and target string
    //find and replacement values
    for i := range indices {
        idx := indices[i]
        replaceIdxMap[idx] = i
    }
    fmt.Println(replaceIdxMap)
    //s consists of only lowercase English letters.
    var sb strings.Builder
    for i :=0; i < len(s); {
        switch _, ok := replaceIdxMap[i]; ok {
        case true:
            fmt.Println("here")
            srcSubStr := sources[replaceIdxMap[i]]
            targetSubStr:= targets[replaceIdxMap[i]]
            startIdx := i
            endIdx := startIdx
            for srcIdx := 0; srcIdx < len(srcSubStr) && srcIdx != len(s); srcIdx++ {
                if srcSubStr[srcIdx] != s[endIdx] {
                    break
                }
                endIdx++
            }
            if endIdx-startIdx == len(srcSubStr) {
                sb.WriteString(targetSubStr)
                i = endIdx
                continue
            }
            fallthrough
        default:
            sb.WriteByte(s[i])
            i++
        }
        
    } 
    
    return sb.String()
}