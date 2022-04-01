// 1910. Remove All Occurrences of a Substring
// time: O(lenStr*lenPart)
// space: O(lenStr)

func removeOccurrences(s string, part string) string {
    lenStr := len(s)
    lenPart := len(part)
    res := make([]byte, lenStr)
    end := 0
    for i := 0; i < lenStr; i++ { 
        res[end] = s[i]
        end++
        if end >= lenPart {
            idxStart := end - lenPart
            isEqual := true 
            partIdx := 0
            idxEnd := idxStart + lenPart
            for idxStart < idxEnd {
                if res[idxStart] != part[partIdx] {
                    isEqual = false
                    break
                }
                idxStart++
                partIdx++            
            }
            if isEqual {
                end -= lenPart
            }
        } 
    }
    return string(res[0:end])
}